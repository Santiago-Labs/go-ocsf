package datastore

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"path/filepath"
	"strings"
	"time"

	"github.com/Santiago-Labs/go-ocsf/ocsf"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/samsarahq/go/oops"
)

type s3JsonDatastore struct {
	s3Bucket string
	s3Client *s3.Client

	findingIndex map[string]string
}

func NewS3JsonDatastore(bucketName string, s3Client *s3.Client) Datastore {
	s := &s3JsonDatastore{
		s3Bucket: bucketName,
		s3Client: s3Client,

		findingIndex: make(map[string]string),
	}

	ctx := context.Background()
	if err := s.buildFindingIndex(ctx); err != nil {
		slog.Warn("failed to build complete finding index", "error", err)
	}

	return s
}

func (s *s3JsonDatastore) GetFinding(ctx context.Context, findingID string) (*ocsf.VulnerabilityFinding, error) {
	filePath, exists := s.findingIndex[findingID]
	if !exists {
		return nil, ErrNotFound
	}

	findings, err := s.GetAllFindings(ctx, filePath)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to get all findings")
	}

	for _, finding := range findings {
		if finding.FindingInfo.UID == findingID {
			return &finding, nil
		}
	}

	return nil, ErrNotFound
}

func (s *s3JsonDatastore) GetAllFindings(ctx context.Context, key string) ([]ocsf.VulnerabilityFinding, error) {
	result, err := s.s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.s3Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, oops.Wrapf(err, "failed to get JSON file from S3")
	}
	defer result.Body.Close()

	data, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to read JSON file data")
	}

	var findings struct {
		VulnerabilityFindings []ocsf.VulnerabilityFinding `json:"vulnerability_findings"`
	}

	if err := json.Unmarshal(data, &findings); err != nil {
		return nil, oops.Wrapf(err, "failed to parse JSON file")
	}

	return findings.VulnerabilityFindings, nil
}

func (s *s3JsonDatastore) SaveFindings(ctx context.Context, findings []ocsf.VulnerabilityFinding) error {
	newFindings := []ocsf.VulnerabilityFinding{}
	existingFindings := make(map[string][]ocsf.VulnerabilityFinding)

	for _, finding := range findings {
		filePath, exists := s.findingIndex[finding.FindingInfo.UID]
		if !exists {
			newFindings = append(newFindings, finding)
		} else {
			existingFindings[filePath] = append(existingFindings[filePath], finding)
		}
	}

	newFindingsFilename := fmt.Sprintf("%s.json", time.Now().Format("20060102T150405Z"))
	newFindingsKey := filepath.Join(basepath, newFindingsFilename)
	if len(newFindings) > 0 {
		if err := s.createFile(ctx, newFindingsKey, newFindings); err != nil {
			return oops.Wrapf(err, "failed to write new findings to parquet")
		}
	}

	updatedCount := 0
	for updatedFindingsPath, updatedFindings := range existingFindings {
		if err := s.updateFile(ctx, updatedFindingsPath, updatedFindings); err != nil {
			slog.Error("failed to update json file", "file", updatedFindingsPath, "error", err)
			// Fall back to creating a new file for these findings

			newFindingsFilename := fmt.Sprintf("%s.json", time.Now().Format("20060102T150405Z"))
			newFindingsKey := filepath.Join(basepath, newFindingsFilename)
			if err := s.createFile(ctx, newFindingsKey, updatedFindings); err != nil {
				return oops.Wrapf(err, "failed to write fallback parquet file")
			}
		} else {
			updatedCount += len(updatedFindings)
		}
	}

	slog.Info("saved findings",
		"new", len(newFindings),
		"updated", updatedCount)

	return nil
}

func (s *s3JsonDatastore) createFile(ctx context.Context, key string, findings []ocsf.VulnerabilityFinding) error {
	outerSchema := map[string]interface{}{
		"vulnerability_findings": findings,
	}
	jsonData, err := json.Marshal(outerSchema)
	if err != nil {
		return oops.Wrapf(err, "failed to marshal findings to JSON")
	}

	_, err = s.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &s.s3Bucket,
		Key:    &key,
		Body:   bytes.NewReader(jsonData),
	})
	if err != nil {
		return oops.Wrapf(err, "failed to upload JSON to S3")
	}

	slog.Info("Wrote JSON file to S3",
		"bucket", s.s3Bucket,
		"key", key,
	)

	return nil
}

func (s *s3JsonDatastore) loadFileIntoIndex(ctx context.Context, key string) error {
	result, err := s.s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.s3Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return oops.Wrapf(err, "failed to get JSON file from S3")
	}
	defer result.Body.Close()

	data, err := io.ReadAll(result.Body)
	if err != nil {
		return oops.Wrapf(err, "failed to read JSON file")
	}

	var findings struct {
		VulnerabilityFindings []ocsf.VulnerabilityFinding `json:"vulnerability_findings"`
	}

	if err := json.Unmarshal(data, &findings); err != nil {
		return oops.Wrapf(err, "failed to parse JSON file")
	}

	for _, finding := range findings.VulnerabilityFindings {
		s.findingIndex[finding.FindingInfo.UID] = key
	}

	return nil
}

func (s *s3JsonDatastore) buildFindingIndex(ctx context.Context) error {
	paginator := s3.NewListObjectsV2Paginator(s.s3Client, &s3.ListObjectsV2Input{
		Bucket: aws.String(s.s3Bucket),
		Prefix: aws.String(basepath),
	})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return oops.Wrapf(err, "failed to list objects in S3")
		}

		for _, object := range output.Contents {
			if strings.HasSuffix(*object.Key, "/") || !strings.HasSuffix(*object.Key, ".json") {
				continue
			}

			if err := s.loadFileIntoIndex(ctx, *object.Key); err != nil {
				slog.Warn("error indexing json file", "key", *object.Key, "error", err)
			}
		}
	}

	slog.Info("built finding index from S3", "count", len(s.findingIndex))
	return nil
}

func (s *s3JsonDatastore) updateFile(ctx context.Context, filePath string, updatedFindings []ocsf.VulnerabilityFinding) error {
	updateMap := make(map[string]ocsf.VulnerabilityFinding)
	for _, finding := range updatedFindings {
		updateMap[finding.FindingInfo.UID] = finding
	}

	allFindings, err := s.GetAllFindings(ctx, filePath)
	if err != nil {
		return oops.Wrapf(err, "failed to read all findings from JSON file")
	}

	for i, finding := range allFindings {
		if updated, exists := updateMap[finding.FindingInfo.UID]; exists {
			allFindings[i] = updated
			slog.Debug("updated finding in memory", "id", finding.FindingInfo.UID)
		}
	}

	outerSchema := map[string]interface{}{
		"vulnerability_findings": allFindings,
	}
	jsonData, err := json.Marshal(outerSchema)
	if err != nil {
		return oops.Wrapf(err, "failed to marshal findings to JSON")
	}

	_, err = s.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.s3Bucket),
		Key:    aws.String(filePath),
		Body:   bytes.NewReader(jsonData),
	})
	if err != nil {
		return oops.Wrapf(err, "failed to upload updated JSON to S3")
	}

	slog.Info("replaced JSON file in S3",
		"bucket", s.s3Bucket,
		"key", filePath,
		"size", len(jsonData),
		"findings", len(updatedFindings))

	return nil
}
