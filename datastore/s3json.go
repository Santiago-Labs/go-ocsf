package datastore

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"path/filepath"
	"reflect"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/samsarahq/go/oops"
)

type s3JsonDatastore[T any] struct {
	s3Bucket string
	s3Client *s3.Client

	currentPath string
	basePath    string

	BaseDatastore[T]
}

// NewS3JsonDatastore creates a new S3 JSON datastore.
func NewS3JsonDatastore[T any](ctx context.Context, bucketName string, s3Client *s3.Client) (Datastore[T], error) {

	typeName := reflect.TypeOf((*T)(nil)).Elem().Name()

	s := &s3JsonDatastore[T]{
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
// It reads the gzipped JSON file and parses it into a slice of ocsf data.
func (s *s3JsonDatastore[T]) GetItemsFile(ctx context.Context, key string) ([]T, error) {
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

	var items []T
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, oops.Wrapf(err, "failed to parse JSON file")
	}

	return items, nil
}

// WriteBatch creates a new JSON file for storing ocsf data.
// It marshals the data into a JSON object and writes it to the specified file path.
func (s *s3JsonDatastore[T]) WriteBatch(ctx context.Context, items []T) error {
	allItems := items

	if s.currentPath == "" {
		s.currentPath = filepath.Join(s.basePath, fmt.Sprintf("%s.json", time.Now().Format("20060102T150405Z")))
	} else {
		fileItems, err := s.GetItemsFile(ctx, s.currentPath)
		if err != nil {
			return oops.Wrapf(err, "failed to get existing items from s3")
		}

		allItems = append(allItems, fileItems...)
	}

	jsonData, err := json.Marshal(allItems)
	if err != nil {
		return oops.Wrapf(err, "failed to marshal to JSON")
	}

	_, err = s.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      &s.s3Bucket,
		Key:         &s.currentPath,
		Body:        bytes.NewReader(jsonData),
		ContentType: aws.String("application/json"),
	})
	if err != nil {
		return oops.Wrapf(err, "failed to upload JSON to S3")
	}

	slog.Info("Wrote JSON file to S3",
		"bucket", s.s3Bucket,
		"key", s.currentPath,
		"items", len(allItems),
	)

	return nil
}
