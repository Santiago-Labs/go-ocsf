package datastore

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"path/filepath"
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
func NewS3ParquetDatastore(bucketName string, s3Client *s3.Client) Datastore {
	s := &s3ParquetDatastore{
		s3Bucket: bucketName,
		s3Client: s3Client,
	}

	s.BaseDatastore = BaseDatastore{
		store: s,
	}

	return s
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
func (s *s3ParquetDatastore) WriteBatch(ctx context.Context, findings []ocsf.VulnerabilityFinding, key *string) error {
	allFindings := findings
	if key == nil {
		newkey := filepath.Join(Basepath, fmt.Sprintf("%s.parquet", time.Now().Format("20060102T150405Z")))
		key = &newkey
	} else {
		var err error
		allFindings, err = s.GetFindingsFromFile(ctx, *key)
		if err != nil {
			return oops.Wrapf(err, "failed to get existing findings from disk")
		}

		allFindings = append(allFindings, findings...)
	}

	var buf bytes.Buffer
	writer := io.Writer(&buf)
	if err := goParquet.Write(writer, allFindings); err != nil {
		return oops.Wrapf(err, "failed to write findings to parquet buffer")
	}

	_, err := s.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &s.s3Bucket,
		Key:    key,
		Body:   bytes.NewReader(buf.Bytes()),
	})
	if err != nil {
		return oops.Wrapf(err, "failed to upload Parquet to S3")
	}

	slog.Info("Wrote Parquet file to S3",
		"bucket", s.s3Bucket,
		"key", *key,
		"findings", len(allFindings),
	)
	return nil
}

func (s *s3ParquetDatastore) WriteAPIActivityBatch(ctx context.Context, activities []ocsf.APIActivity, key *string) error {
	allActivities := activities
	if key == nil {
		newkey := filepath.Join(BasepathActivities, fmt.Sprintf("%s.parquet", time.Now().Format("20060102T150405Z")))
		key = &newkey
	} else {
		var err error
		allActivities, err = s.GetAPIActivitiesFromFile(ctx, *key)
		if err != nil {
			return oops.Wrapf(err, "failed to get existing activities from disk")
		}

		allActivities = append(allActivities, activities...)
	}

	var buf bytes.Buffer
	writer := io.Writer(&buf)
	if err := goParquet.Write(writer, allActivities); err != nil {
		return oops.Wrapf(err, "failed to write activities to parquet buffer")
	}

	_, err := s.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &s.s3Bucket,
		Key:    key,
		Body:   bytes.NewReader(buf.Bytes()),
	})
	if err != nil {
		return oops.Wrapf(err, "failed to upload Parquet to S3")
	}

	slog.Info("Wrote Parquet file to S3",
		"bucket", s.s3Bucket,
		"key", *key,
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
