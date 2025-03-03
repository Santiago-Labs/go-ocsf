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
func NewS3ParquetDatastore(bucketName string, s3Client *s3.Client) Datastore {
	s := &s3ParquetDatastore{
		s3Bucket: bucketName,
		s3Client: s3Client,
	}
	s.BaseDatastore = BaseDatastore{
		findingIndex: make(map[string]string),
		fileIndex:    make(map[string]int),
		store:        s,
	}

	ctx := context.Background()
	if err := s.buildFindingIndex(ctx); err != nil {
		slog.Warn("failed to build complete finding index", "error", err)
	}

	return s
}

// buildFindingIndex builds the datastore's in-memory index of finding IDs to file paths.
// It reads all Parquet files in the base directory and parses them into a slice of vulnerability findings.
// The datastore updates its in-memory index with the finding IDs and their corresponding file paths.
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

	slog.Info("built finding index from S3", "count", len(s.BaseDatastore.findingIndex))
	return nil
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

func (s *s3ParquetDatastore) WriteBatch(ctx context.Context, findings []ocsf.VulnerabilityFinding, key *string) error {
	allFindings := findings
	if key == nil {
		newkey := filepath.Join(basepath, fmt.Sprintf("%s.parquet", time.Now().Format("20060102T150405Z")))
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

	for _, f := range allFindings {
		s.BaseDatastore.findingIndex[f.FindingInfo.UID] = *key
	}
	s.BaseDatastore.fileIndex[*key] = len(allFindings)

	slog.Info("Wrote Parquet file to S3",
		"bucket", s.s3Bucket,
		"key", *key,
		"findings", len(allFindings),
	)
	return nil
}
