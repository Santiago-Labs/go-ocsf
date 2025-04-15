package datastore

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"path/filepath"
	"time"

	"github.com/Santiago-Labs/go-ocsf/ocsf"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/samsarahq/go/oops"
)

type s3JsonDatastore struct {
	s3Bucket string
	s3Client *s3.Client

	BaseDatastore
}

// NewS3JsonDatastore creates a new S3 JSON datastore.
// It initializes an in-memory index of finding IDs to file paths.
func NewS3JsonDatastore(bucketName string, s3Client *s3.Client) Datastore {
	s := &s3JsonDatastore{
		s3Bucket: bucketName,
		s3Client: s3Client,
	}

	s.BaseDatastore = BaseDatastore{
		store: s,
	}

	return s
}

// GetFindingsFromFile retrieves all vulnerability findings from a specific file path.
// It reads the JSON file and parses it into a slice of vulnerability findings.
func (s *s3JsonDatastore) GetFindingsFromFile(ctx context.Context, key string) ([]ocsf.VulnerabilityFinding, error) {
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
		VulnerabilityFindings []ocsf.VulnerabilityFinding `json:"vulnerability_finding"`
	}

	if err := json.Unmarshal(data, &findings); err != nil {
		return nil, oops.Wrapf(err, "failed to parse JSON file")
	}

	return findings.VulnerabilityFindings, nil
}

// WriteBatch creates a new JSON file for storing vulnerability findings.
// It marshals the findings into a JSON object and writes it to the specified file path.
func (s *s3JsonDatastore) WriteBatch(ctx context.Context, findings []ocsf.VulnerabilityFinding, key *string) error {
	allFindings := findings
	if key == nil {
		newkey := filepath.Join(Basepath, fmt.Sprintf("%s.json", time.Now().Format("20060102T150405Z")))
		key = &newkey
	} else {
		var err error
		allFindings, err = s.GetFindingsFromFile(ctx, *key)
		if err != nil {
			return oops.Wrapf(err, "failed to get existing findings from disk")
		}

		allFindings = append(allFindings, findings...)
	}

	outerSchema := map[string]interface{}{
		"vulnerability_finding": allFindings,
	}
	jsonData, err := json.Marshal(outerSchema)
	if err != nil {
		return oops.Wrapf(err, "failed to marshal findings to JSON")
	}

	_, err = s.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &s.s3Bucket,
		Key:    key,
		Body:   bytes.NewReader(jsonData),
	})
	if err != nil {
		return oops.Wrapf(err, "failed to upload JSON to S3")
	}

	slog.Info("Wrote JSON file to S3",
		"bucket", s.s3Bucket,
		"key", *key,
		"findings", len(allFindings),
	)

	return nil
}

func (s *s3JsonDatastore) GetAPIActivitiesFromFile(ctx context.Context, key string) ([]ocsf.APIActivity, error) {
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

	var activities struct {
		APIActivities []ocsf.APIActivity `json:"api_activities"`
	}

	if err := json.Unmarshal(data, &activities); err != nil {
		return nil, oops.Wrapf(err, "failed to parse JSON file")
	}

	return activities.APIActivities, nil
}

func (s *s3JsonDatastore) WriteAPIActivityBatch(ctx context.Context, activities []ocsf.APIActivity, key *string) error {
	allActivities := activities
	if key == nil {
		newkey := filepath.Join(BasepathActivities, fmt.Sprintf("%s.json", time.Now().Format("20060102T150405Z")))
		key = &newkey
	} else {
		var err error
		allActivities, err = s.GetAPIActivitiesFromFile(ctx, *key)
		if err != nil {
			return oops.Wrapf(err, "failed to get existing activities from disk")
		}

		allActivities = append(allActivities, activities...)
	}

	outerSchema := map[string]interface{}{
		"api_activities": allActivities,
	}
	jsonData, err := json.Marshal(outerSchema)
	if err != nil {
		return oops.Wrapf(err, "failed to marshal activities to JSON")
	}

	_, err = s.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &s.s3Bucket,
		Key:    key,
		Body:   bytes.NewReader(jsonData),
	})
	if err != nil {
		return oops.Wrapf(err, "failed to upload JSON to S3")
	}

	slog.Info("Wrote JSON file to S3",
		"bucket", s.s3Bucket,
		"key", *key,
		"activities", len(allActivities),
	)

	return nil
}
