package datastore

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3tables"
	"github.com/samsarahq/go/oops"
)

type StorageOpts struct {
	IsParquet      bool
	IsJSON         bool
	BucketName     string
	TableBucketArn string
}

type BaseDatastore[T any] struct {
	store Datastore[T]
}

// SaveFindings saves a batch of findings to the datastore. Datastore implementations handle file formats.
func (d *BaseDatastore[T]) Save(ctx context.Context, items []T) error {
	if err := d.store.WriteBatch(ctx, items); err != nil {
		return err
	}
	slog.Info("upserted items", "items", len(items))

	return nil
}

func filesExistS3(ctx context.Context, s3Client *s3.Client, s3Bucket, path, extension string) (bool, error) {
	resp, err := s3Client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: &s3Bucket,
		Prefix: &path,
	})
	if err != nil {
		return false, oops.Wrapf(err, "failed to list objects in S3")
	}

	files := resp.Contents
	if len(files) > 0 {
		for _, file := range files {
			if strings.HasSuffix(*file.Key, extension) {
				return true, nil
			}
		}
	}

	return false, nil
}

func SetupStorage[T any](ctx context.Context, opts StorageOpts) (Datastore[T], error) {
	var storage Datastore[T]
	var s3Client *s3.Client
	var err error

	if opts.IsParquet {
		if opts.TableBucketArn != "" {

			cfg, err := config.LoadDefaultConfig(ctx)
			if err != nil {
				return nil, oops.Wrapf(err, "failed to load config")
			}
			s3tablesClient := s3tables.NewFromConfig(cfg)

			storage, err = NewS3TablesDatastore[T](ctx, opts.TableBucketArn, s3tablesClient)
			if err != nil {
				return nil, fmt.Errorf("failed to create S3 tables datastore: %v", err)
			}
		} else if opts.BucketName != "" {
			cfg, err := config.LoadDefaultConfig(ctx)
			if err != nil {
				return nil, fmt.Errorf("error loading AWS config: %v", err)
			}

			s3Client = s3.NewFromConfig(cfg)
			storage, err = NewS3ParquetDatastore[T](ctx, opts.BucketName, s3Client)
			if err != nil {
				return nil, fmt.Errorf("failed to create S3 parquet datastore: %v", err)
			}
		} else {
			storage, err = NewLocalParquetDatastore[T](ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to create local parquet datastore: %v", err)
			}
		}
	} else if opts.IsJSON {
		if opts.BucketName != "" {
			cfg, err := config.LoadDefaultConfig(ctx)
			if err != nil {
				return nil, fmt.Errorf("error loading AWS config: %v", err)
			}

			s3Client = s3.NewFromConfig(cfg)
			storage, err = NewS3JsonDatastore[T](ctx, opts.BucketName, s3Client)
			if err != nil {
				return nil, fmt.Errorf("failed to create S3 json datastore: %v", err)
			}
		} else {
			storage, err = NewLocalJsonDatastore[T](ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to create local json datastore: %v", err)
			}
		}
	} else {
		return nil, fmt.Errorf("no storage format specified, use --parquet or --json")
	}

	return storage, nil
}
