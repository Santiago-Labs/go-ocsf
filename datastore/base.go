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

func (d *BaseDatastore) loadActivityFileIntoIndex(ctx context.Context, key string) error {
	activities, err := d.store.GetAPIActivitiesFromFile(ctx, key)
	if err != nil {
		return oops.Wrapf(err, "failed to read all activities from parquet file")
	}

	for _, activity := range activities {
		d.activityIndex[*activity.Metadata.CorrelationUID] = key
	}
	d.activityFileIndex[key] = len(activities)

	slog.Debug("indexed activities from file", "key", key, "count", len(activities))
	return nil
}

func (d *BaseDatastore) SaveAPIActivities(ctx context.Context, activities []ocsf.APIActivity) error {
	var (
		currentBatch     []ocsf.APIActivity
		currentBatchSize int
	)

	for _, activity := range activities {
		if currentBatchSize > maxFileSize && len(currentBatch) > 0 {
			pathWithSpace := d.getAPIActivityPathWithSpace(len(currentBatch))
			if err := d.store.WriteAPIActivityBatch(ctx, currentBatch, pathWithSpace); err != nil {
				return err
			}
			currentBatch = []ocsf.APIActivity{}
			currentBatchSize = 0
		}

		currentBatch = append(currentBatch, activity)
		currentBatchSize += avgFindingSize
	}

	if len(currentBatch) > 0 {
		pathWithSpace := d.getAPIActivityPathWithSpace(len(currentBatch))
		if pathWithSpace != nil {
		}
		if err := d.store.WriteAPIActivityBatch(ctx, currentBatch, pathWithSpace); err != nil {
			return err
		}
	}

	slog.Info("upserted activities", "activities", len(activities))

	return nil
}

func (d *BaseDatastore) getAPIActivityPathWithSpace(batchSize int) *string {
	for path, count := range d.activityFileIndex {
		if (count+batchSize)*avgFindingSize < maxFileSize {
			return &path
		}
	}

	return nil
}

// GetAPIActivity returns a single activity from the datastore or ErrNotFound if the activity is not found
func (d *BaseDatastore) GetAPIActivity(ctx context.Context, activityID string) (*ocsf.APIActivity, error) {
	filePath, exists := d.activityIndex[activityID]
	if !exists {
		return nil, ErrNotFound
	}

	activities, err := d.store.GetAPIActivitiesFromFile(ctx, filePath)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to get all activities")
	}

	for _, activity := range activities {
		if *activity.Metadata.CorrelationUID == activityID {
			return &activity, nil
		}
	}

	return nil, ErrNotFound
}
