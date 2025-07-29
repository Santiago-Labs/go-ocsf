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

	basePath string

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

// WriteBatch creates a new Parquet file for storing ocsf data.
// It writes the data to the specified file path
func (s *s3ParquetDatastore[T]) WriteBatch(ctx context.Context, items []T) error {
	savePath := filepath.Join(s.basePath, fmt.Sprintf("%s.parquet.gz", time.Now().Format("20060102T150405Z")))

	var buf bytes.Buffer
	writer := io.Writer(&buf)
	if err := goParquet.Write[T](writer, items, goParquet.Compression(&goParquet.Gzip)); err != nil {
		return oops.Wrapf(err, "failed to write to parquet buffer")
	}

	_, err := s.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:          &s.s3Bucket,
		Key:             &savePath,
		Body:            bytes.NewReader(buf.Bytes()),
		ContentType:     aws.String("application/octet-stream"),
		ContentEncoding: aws.String("gzip"),
	})
	if err != nil {
		return oops.Wrapf(err, "failed to upload Parquet to S3")
	}

	slog.Info("Wrote Parquet file to S3",
		"bucket", s.s3Bucket,
		"key", savePath,
		"items", len(items),
	)
	return nil
}
