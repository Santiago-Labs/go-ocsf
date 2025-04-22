package datastore

import (
	"context"
	"log/slog"
	"strings"

	"github.com/Santiago-Labs/go-ocsf/ocsf"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/samsarahq/go/oops"
)

type BaseDatastore struct {
	store Datastore

	apiActivitiesTableName string
	findingsTableName      string
}

// SaveFindings saves a batch of findings to the datastore. Datastore implementations handle file formats.
func (d *BaseDatastore) SaveFindings(ctx context.Context, findings []ocsf.VulnerabilityFinding) error {
	if err := d.store.WriteBatch(ctx, findings); err != nil {
		return err
	}
	slog.Info("upserted findings", "findings", len(findings))

	return nil
}

func (d *BaseDatastore) SaveAPIActivities(ctx context.Context, activities []ocsf.APIActivity) error {
	if err := d.store.WriteAPIActivityBatch(ctx, activities); err != nil {
		return oops.Wrapf(err, "failed to write batch")
	}

	slog.Info("upserted activities", "activities", len(activities))

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
