package datastore

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"path/filepath"
	"strings"
	"time"

	"github.com/Santiago-Labs/go-ocsf/ocsf"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	goParquet "github.com/parquet-go/parquet-go"
	"github.com/samsarahq/go/oops"
)

type s3ParquetDatastore struct {
	s3Bucket string
	s3Client *s3.Client

	findingIndex map[string]string
}

func NewS3ParquetDatastore(bucketName string, s3Client *s3.Client) Datastore {
	s := &s3ParquetDatastore{
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

func (s *s3ParquetDatastore) buildFindingIndex(ctx context.Context) error {
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
			if strings.HasSuffix(*object.Key, "/") || !strings.HasSuffix(*object.Key, ".parquet") {
				continue
			}

			if err := s.loadFileIntoIndex(ctx, *object.Key); err != nil {
				slog.Warn("error indexing parquet file", "key", *object.Key, "error", err)
			}
		}
	}

	slog.Info("built finding index from S3", "count", len(s.findingIndex))
	return nil
}

func (s *s3ParquetDatastore) GetFinding(ctx context.Context, findingID string) (*ocsf.VulnerabilityFinding, error) {
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

func (s *s3ParquetDatastore) GetAllFindings(ctx context.Context, key string) ([]ocsf.VulnerabilityFinding, error) {
	result, err := s.s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.s3Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, oops.Wrapf(err, "failed to get parquet file from S3")
	}
	defer result.Body.Close()

	data, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to read parquet file data")
	}

	findings, err := goParquet.Read[ocsf.VulnerabilityFinding](bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return nil, oops.Wrapf(err, "failed to read parquet file")
	}

	return findings, nil
}

func (s *s3ParquetDatastore) SaveFindings(ctx context.Context, findings []ocsf.VulnerabilityFinding) error {
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

	newFindingsFilename := fmt.Sprintf("%s.parquet", time.Now().Format("20060102T150405Z"))
	newFindingsKey := filepath.Join(basepath, newFindingsFilename)
	if len(newFindings) > 0 {
		if err := s.createFile(ctx, newFindingsKey, newFindings); err != nil {
			return oops.Wrapf(err, "failed to write new findings to parquet")
		}
	}

	updatedCount := 0
	for updatedFindingsPath, updatedFindings := range existingFindings {
		if err := s.updateFile(ctx, updatedFindingsPath, updatedFindings); err != nil {
			slog.Error("failed to update parquet file", "file", updatedFindingsPath, "error", err)
			// Fall back to creating a new file for these findings

			newFindingsFilename := fmt.Sprintf("%s.parquet", time.Now().Format("20060102T150405Z"))
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

func (s *s3ParquetDatastore) createFile(ctx context.Context, key string, findings []ocsf.VulnerabilityFinding) error {
	var buf bytes.Buffer
	writer := io.Writer(&buf)
	if err := goParquet.Write(writer, findings); err != nil {
		return oops.Wrapf(err, "failed to write findings to parquet buffer")
	}

	for _, f := range findings {
		s.findingIndex[f.FindingInfo.UID] = key
	}

	_, err := s.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &s.s3Bucket,
		Key:    &key,
		Body:   bytes.NewReader(buf.Bytes()),
	})
	if err != nil {
		return oops.Wrapf(err, "failed to upload Parquet to S3")
	}

	slog.Info("Wrote Parquet file to S3",
		"bucket", s.s3Bucket,
		"key", key,
	)
	return nil
}

func (s *s3ParquetDatastore) loadFileIntoIndex(ctx context.Context, key string) error {
	findings, err := s.GetAllFindings(ctx, key)
	if err != nil {
		return oops.Wrapf(err, "failed to read all findings from parquet file")
	}

	for _, finding := range findings {
		s.findingIndex[finding.FindingInfo.UID] = key
	}

	slog.Debug("indexed findings from s3", "key", key, "count", len(findings))
	return nil
}

func (s *s3ParquetDatastore) updateFile(ctx context.Context, filePath string, updatedFindings []ocsf.VulnerabilityFinding) error {
	updateMap := make(map[string]ocsf.VulnerabilityFinding)
	for _, finding := range updatedFindings {
		updateMap[finding.FindingInfo.UID] = finding
	}

	allFindings, err := s.GetAllFindings(ctx, filePath)
	if err != nil {
		return oops.Wrapf(err, "failed to read all findings from parquet file")
	}

	for i, finding := range allFindings {
		if updated, exists := updateMap[finding.FindingInfo.UID]; exists {
			allFindings[i] = updated
			slog.Debug("updated finding in memory", "id", finding.FindingInfo.UID)
		}
	}

	var buf bytes.Buffer
	writer := io.Writer(&buf)
	if err := goParquet.Write(writer, allFindings); err != nil {
		return oops.Wrapf(err, "failed to write findings to parquet buffer")
	}

	_, err = s.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.s3Bucket),
		Key:    aws.String(filePath),
		Body:   bytes.NewReader(buf.Bytes()),
	})
	if err != nil {
		return oops.Wrapf(err, "failed to upload updated parquet to S3")
	}

	slog.Info("replaced parquet file in S3",
		"bucket", s.s3Bucket,
		"key", filePath,
		"findings", len(allFindings))

	return nil
}
