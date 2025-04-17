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

	BaseDatastore
}

// NewS3ParquetDatastore creates a new S3 Parquet datastore.
// It initializes an in-memory index of finding IDs to file paths.
func NewS3ParquetDatastore(ctx context.Context, bucketName string, s3Client *s3.Client) (Datastore, error) {
	s := &s3ParquetDatastore{
		s3Bucket: bucketName,
		s3Client: s3Client,
	}

	s.BaseDatastore = BaseDatastore{
		store:                  s,
		findingsTableName:      "vulnerability_finding",
		apiActivitiesTableName: "api_activities",
	}

	return s, nil
}

// GetFindingsFromFile retrieves all vulnerability findings from a specific file path.
// It reads the Parquet file and parses it into a slice of vulnerability findings.
func (s *s3ParquetDatastore) GetFindingsFromFile(ctx context.Context, key string) ([]ocsf.VulnerabilityFinding, error) {
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

// WriteBatch creates a new Parquet file for storing vulnerability findings.
// It writes the findings to the specified file path and updates the datastore's in-memory index.
func (s *s3ParquetDatastore) WriteBatch(ctx context.Context, findings []ocsf.VulnerabilityFinding, keyPrefix string) error {
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
			if strings.HasSuffix(*file.Key, ".parquet.gz") {
				fullKey = *file.Key

				fileFindings, err := s.GetFindingsFromFile(ctx, *file.Key)
				if err != nil {
					return oops.Wrapf(err, "failed to get existing activities from disk")
				}

				allFindings = append(allFindings, fileFindings...)
			}
		}
	}

	if fullKey == "" {
		fullKey = filepath.Join(keyPrefix, fmt.Sprintf("%s.parquet.gz", time.Now().Format("20060102T150405Z")))
	}

	var buf bytes.Buffer
	writer := io.Writer(&buf)
	if err := goParquet.Write(writer, allFindings, goParquet.Compression(&goParquet.Gzip)); err != nil {
		return oops.Wrapf(err, "failed to write findings to parquet buffer")
	}

	_, err = s.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:          &s.s3Bucket,
		Key:             &fullKey,
		Body:            bytes.NewReader(buf.Bytes()),
		ContentType:     aws.String("application/octet-stream"),
		ContentEncoding: aws.String("gzip"),
	})
	if err != nil {
		return oops.Wrapf(err, "failed to upload Parquet to S3")
	}

	slog.Info("Wrote Parquet file to S3",
		"bucket", s.s3Bucket,
		"key", fullKey,
		"findings", len(allFindings),
	)
	return nil
}

func (s *s3ParquetDatastore) WriteAPIActivityBatch(ctx context.Context, activities []ocsf.APIActivity, keyPrefix string) error {
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
			if strings.HasSuffix(*file.Key, ".parquet.gz") {
				fullKey = *file.Key

				fileActivities, err := s.GetAPIActivitiesFromFile(ctx, *file.Key)
				if err != nil {
					return oops.Wrapf(err, "failed to get existing activities from s3")
				}

				allActivities = append(allActivities, fileActivities...)
			}
		}
	}

	if fullKey == "" {
		fullKey = filepath.Join(keyPrefix, fmt.Sprintf("%s.parquet.gz", time.Now().Format("20060102T150405Z")))
	}

	var buf bytes.Buffer
	writer := io.Writer(&buf)
	if err := goParquet.Write(writer, allActivities, goParquet.Compression(&goParquet.Gzip)); err != nil {
		return oops.Wrapf(err, "failed to write activities to parquet buffer")
	}

	_, err = s.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:          &s.s3Bucket,
		Key:             &fullKey,
		Body:            bytes.NewReader(buf.Bytes()),
		ContentType:     aws.String("application/octet-stream"),
		ContentEncoding: aws.String("gzip"),
	})
	if err != nil {
		return oops.Wrapf(err, "failed to upload Parquet to S3")
	}

	slog.Info("Wrote Parquet file to S3",
		"bucket", s.s3Bucket,
		"key", fullKey,
		"activities", len(allActivities),
	)
	return nil
}

func (s *s3ParquetDatastore) GetAPIActivitiesFromFile(ctx context.Context, key string) ([]ocsf.APIActivity, error) {
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

	activities, err := goParquet.Read[ocsf.APIActivity](bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return nil, oops.Wrapf(err, "failed to read parquet file")
	}

	return activities, nil
}
