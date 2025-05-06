package datastore

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"path/filepath"
	"reflect"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	goParquet "github.com/parquet-go/parquet-go"
	"github.com/samsarahq/go/oops"
)

type s3ParquetDatastore[T any] struct {
	s3Bucket string
	s3Client *s3.Client

	currentPath string
	basePath    string

	BaseDatastore[T]
}

// NewS3ParquetDatastore creates a new S3 Parquet datastore.
func NewS3ParquetDatastore[T any](ctx context.Context, bucketName string, s3Client *s3.Client) (Datastore[T], error) {
	typeName := reflect.TypeOf((*T)(nil)).Elem().Name()

	s := &s3ParquetDatastore[T]{
		s3Bucket: bucketName,
		s3Client: s3Client,
		basePath: basepaths[typeName],
	}

	s.BaseDatastore = BaseDatastore[T]{
		store: s,
	}

	return s, nil
}

// GetItemsFromFile retrieves all ocsf data from a specific file path.
// It reads the Parquet file and parses it into a slice of ocsf data.
func (s *s3ParquetDatastore[T]) GetItemsFromFile(ctx context.Context, key string) ([]T, error) {
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

	items, err := goParquet.Read[T](bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return nil, oops.Wrapf(err, "failed to read parquet file")
	}

	return items, nil
}

// WriteBatch creates a new Parquet file for storing ocsf data.
// It writes the data to the specified file path
func (s *s3ParquetDatastore[T]) WriteBatch(ctx context.Context, items []T) error {
	allItems := items

	if s.currentPath == "" {
		s.currentPath = filepath.Join(s.basePath, fmt.Sprintf("%s.parquet.gz", time.Now().Format("20060102T150405Z")))
	} else {
		fileItems, err := s.GetItemsFromFile(ctx, s.currentPath)
		if err != nil {
			return oops.Wrapf(err, "failed to get existing items from disk")
		}

		allItems = append(allItems, fileItems...)
	}

	var buf bytes.Buffer
	writer := io.Writer(&buf)
	if err := goParquet.Write[T](writer, allItems, goParquet.Compression(&goParquet.Gzip)); err != nil {
		return oops.Wrapf(err, "failed to write to parquet buffer")
	}

	_, err := s.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:          &s.s3Bucket,
		Key:             &s.currentPath,
		Body:            bytes.NewReader(buf.Bytes()),
		ContentType:     aws.String("application/octet-stream"),
		ContentEncoding: aws.String("gzip"),
	})
	if err != nil {
		return oops.Wrapf(err, "failed to upload Parquet to S3")
	}

	slog.Info("Wrote Parquet file to S3",
		"bucket", s.s3Bucket,
		"key", s.currentPath,
		"items", len(allItems),
	)
	return nil
}
