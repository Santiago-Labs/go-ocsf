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

	"github.com/Santiago-Labs/go-ocsf/clients/duckdb"
	"github.com/Santiago-Labs/go-ocsf/ocsf"
	"github.com/apache/arrow/go/v15/arrow"
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
func NewS3JsonDatastore(bucketName string, s3Client *s3.Client) (Datastore, error) {
	s := &s3JsonDatastore{
		s3Bucket: bucketName,
		s3Client: s3Client,
	}

	dbClient, err := duckdb.NewS3Client(context.Background(), bucketName)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to create S3 client")
	}

	basePatterns := map[string]string{
		"vulnerability_finding": fmt.Sprintf("s3://%s/%s", s.s3Bucket, BasepathFindings),
		"api_activities":        fmt.Sprintf("s3://%s/%s", s.s3Bucket, BasepathActivities),
	}

	fields := map[string][]arrow.Field{
		"vulnerability_finding": ocsf.VulnerabilityFindingFields,
		"api_activities":        ocsf.APIActivityFields,
	}

	var queries string
	queries += "INSTALL json; LOAD json; "
	for view, pattern := range basePatterns {
		selectFields := duckdb.GenerateDuckDBSelectFields(view, pattern, fields[view])
		columnsDict := duckdb.GenerateDuckDBColumnsDict(fields[view])
		if filesExist(pattern) {
			queries += fmt.Sprintf(`
				CREATE OR REPLACE VIEW %s AS
				%s FROM read_json_auto('%s',
				ignore_errors=true,
				union_by_name=true,
				hive_partitioning=true,
				%s
				);`,
				view, selectFields, pattern, columnsDict,
			)
		} else {
			queries += duckdb.GenerateDuckDBNullView(view, fields[view])
		}
	}

	_, err = dbClient.Exec(queries)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to create view")
	}

	s.BaseDatastore = BaseDatastore{
		store: s,
		db:    dbClient,
	}

	return s, nil
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

	var findings []ocsf.VulnerabilityFinding

	if err := json.Unmarshal(data, &findings); err != nil {
		return nil, oops.Wrapf(err, "failed to parse JSON file")
	}

	return findings, nil
}

// WriteBatch creates a new JSON file for storing vulnerability findings.
// It marshals the findings into a JSON object and writes it to the specified file path.
func (s *s3JsonDatastore) WriteBatch(ctx context.Context, findings []ocsf.VulnerabilityFinding, keyPrefix string) error {
	allFindings := findings

	var fullKey string
	resp, err := s.s3Client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: &s.s3Bucket,
		Prefix: &keyPrefix,
	})
	if err != nil {
		return oops.Wrapf(err, "failed to list objects in S3")
	}

	files := resp.Contents
	if len(files) > 0 {
		for _, file := range files {
			if strings.HasSuffix(*file.Key, ".json") {

				fileFindings, err := s.GetFindingsFromFile(ctx, *file.Key)
				if err != nil {
					return oops.Wrapf(err, "failed to get existing activities from disk")
				}

				allFindings = append(allFindings, fileFindings...)

				fullKey = *file.Key
			}
		}
	}

	if fullKey == "" {
		fullKey = filepath.Join(keyPrefix, fmt.Sprintf("%s.json", time.Now().Format("20060102T150405Z")))
	}

	jsonData, err := json.Marshal(allFindings)
	if err != nil {
		return oops.Wrapf(err, "failed to marshal findings to JSON")
	}

	_, err = s.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &s.s3Bucket,
		Key:    &fullKey,
		Body:   bytes.NewReader(jsonData),
	})
	if err != nil {
		return oops.Wrapf(err, "failed to upload JSON to S3")
	}

	slog.Info("Wrote JSON file to S3",
		"bucket", s.s3Bucket,
		"key", fullKey,
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

	var activities []ocsf.APIActivity

	if err := json.Unmarshal(data, &activities); err != nil {
		return nil, oops.Wrapf(err, "failed to parse JSON file")
	}

	return activities, nil
}

func (s *s3JsonDatastore) WriteAPIActivityBatch(ctx context.Context, activities []ocsf.APIActivity, keyPrefix string) error {
	allActivities := activities

	var fullKey string
	resp, err := s.s3Client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: &s.s3Bucket,
		Prefix: &keyPrefix,
	})
	if err != nil {
		return oops.Wrapf(err, "failed to list objects in S3")
	}

	files := resp.Contents
	if len(files) > 0 {
		for _, file := range files {
			if strings.HasSuffix(*file.Key, ".json") {
				fullKey = *file.Key

				fileActivities, err := s.GetAPIActivitiesFromFile(ctx, *file.Key)
				if err != nil {
					return oops.Wrapf(err, "failed to get existing activities from disk")
				}

				allActivities = append(allActivities, fileActivities...)
			}
		}
	}

	if fullKey == "" {
		fullKey = filepath.Join(keyPrefix, fmt.Sprintf("%s.json", time.Now().Format("20060102T150405Z")))
	}

	jsonData, err := json.Marshal(allActivities)
	if err != nil {
		return oops.Wrapf(err, "failed to marshal activities to JSON")
	}

	_, err = s.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &s.s3Bucket,
		Key:    &fullKey,
		Body:   bytes.NewReader(jsonData),
	})
	if err != nil {
		return oops.Wrapf(err, "failed to upload JSON to S3")
	}

	slog.Info("Wrote JSON file to S3",
		"bucket", s.s3Bucket,
		"key", fullKey,
		"activities", len(allActivities),
	)

	return nil
}
