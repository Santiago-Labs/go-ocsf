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

	currentFindingsPath   string
	currentActivitiesPath string

	BaseDatastore
}

// NewS3JsonDatastore creates a new S3 JSON datastore.
func NewS3JsonDatastore(ctx context.Context, bucketName string, s3Client *s3.Client) (Datastore, error) {
	s := &s3JsonDatastore{
		s3Bucket: bucketName,
		s3Client: s3Client,
	}

	s.BaseDatastore = BaseDatastore{
		store: s,
	}

	return s, nil
}

// GetFindingsFromFile retrieves all vulnerability findings from a specific file path.
// It reads the gzipped JSON file and parses it into a slice of vulnerability findings.
func (s *s3JsonDatastore) GetFindingsFromFile(ctx context.Context, key string) ([]ocsf.VulnerabilityFinding, error) {
	result, err := s.s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.s3Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, oops.Wrapf(err, "failed to get gzipped JSON file from S3")
	}
	defer result.Body.Close()

	data, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to read JSON file data")
	}

	var findings []ocsf.VulnerabilityFinding
	if err := json.Unmarshal(data, &findings); err != nil {
		return nil, oops.Wrapf(err, "failed to parse JSON file")
	}

	return findings, nil
}

// WriteBatch creates a new JSON file for storing vulnerability findings.
// It marshals the findings into a JSON object and writes it to the specified file path.
func (s *s3JsonDatastore) WriteBatch(ctx context.Context, findings []ocsf.VulnerabilityFinding) error {
	allFindings := findings

	if s.currentFindingsPath == "" {
		s.currentFindingsPath = filepath.Join(BasepathFindings, fmt.Sprintf("%s.json", time.Now().Format("20060102T150405Z")))
	} else {
		fileFindings, err := s.GetFindingsFromFile(ctx, s.currentFindingsPath)
		if err != nil {
			return oops.Wrapf(err, "failed to get existing activities from s3")
		}

		allFindings = append(allFindings, fileFindings...)
	}

	jsonData, err := json.Marshal(allFindings)
	if err != nil {
		return oops.Wrapf(err, "failed to marshal findings to JSON")
	}

	_, err = s.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      &s.s3Bucket,
		Key:         &s.currentFindingsPath,
		Body:        bytes.NewReader(jsonData),
		ContentType: aws.String("application/json"),
	})
	if err != nil {
		return oops.Wrapf(err, "failed to upload JSON to S3")
	}

	slog.Info("Wrote JSON file to S3",
		"bucket", s.s3Bucket,
		"key", s.currentFindingsPath,
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
		return nil, oops.Wrapf(err, "failed to get gzipped JSON file from S3")
	}
	defer result.Body.Close()

	data, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to read JSON file data")
	}

	var activities []ocsf.APIActivity

	if err := json.Unmarshal(data, &activities); err != nil {
		return nil, oops.Wrapf(err, "failed to parse JSON file")
	}

	return activities, nil
}

func (s *s3JsonDatastore) WriteAPIActivityBatch(ctx context.Context, activities []ocsf.APIActivity) error {
	allActivities := activities

	if s.currentActivitiesPath == "" {
		s.currentActivitiesPath = filepath.Join(BasepathActivities, fmt.Sprintf("%s.json", time.Now().Format("20060102T150405Z")))
	} else {
		fileActivities, err := s.GetAPIActivitiesFromFile(ctx, s.currentActivitiesPath)
		if err != nil {
			return oops.Wrapf(err, "failed to get existing activities from disk")
		}

		allActivities = append(allActivities, fileActivities...)
	}

	jsonData, err := json.Marshal(allActivities)
	if err != nil {
		return oops.Wrapf(err, "failed to marshal activities to JSON")
	}

	_, err = s.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      &s.s3Bucket,
		Key:         &s.currentActivitiesPath,
		Body:        bytes.NewReader(jsonData),
		ContentType: aws.String("application/json"),
	})
	if err != nil {
		return oops.Wrapf(err, "failed to upload JSON to S3")
	}

	slog.Info("Wrote JSON file to S3",
		"bucket", s.s3Bucket,
		"key", s.currentActivitiesPath,
		"activities", len(allActivities),
	)

	return nil
}
