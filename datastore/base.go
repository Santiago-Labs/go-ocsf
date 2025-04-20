package datastore

import (
	"context"
	"fmt"
	"log/slog"
	"path/filepath"
	"strings"
	"time"

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
	findingsByDay := make(map[string][]ocsf.VulnerabilityFinding)
	for _, finding := range findings {
		partitionPath := d.getPartitionPath(finding)
		findingsByDay[partitionPath] = append(findingsByDay[partitionPath], finding)
	}

	for partitionPath, dayFindings := range findingsByDay {
		if err := d.store.WriteBatch(ctx, dayFindings, partitionPath); err != nil {
			return oops.Wrapf(err, "failed to write batch")
		}
	}

	slog.Info("upserted findings", "findings", len(findings))

	return nil
}

func (d *BaseDatastore) SaveAPIActivities(ctx context.Context, activities []ocsf.APIActivity) error {
	activitiesByDay := make(map[string][]ocsf.APIActivity)
	for _, activity := range activities {
		partitionPath := d.getAPIActivityPartitionPath(activity)
		activitiesByDay[partitionPath] = append(activitiesByDay[partitionPath], activity)
	}

	for partitionPath, dayActivities := range activitiesByDay {
		if err := d.store.WriteAPIActivityBatch(ctx, dayActivities, partitionPath); err != nil {
			return oops.Wrapf(err, "failed to write batch")
		}
	}

	slog.Info("upserted activities", "activities", len(activities))

	return nil
}

// getPartitionPath returns a path for a finding based on its event time
func (d *BaseDatastore) getPartitionPath(finding ocsf.VulnerabilityFinding) string {
	eventDay := time.UnixMilli(finding.Time).Format("2006-01-02")
	return filepath.Join(BasepathFindings, fmt.Sprintf("event_day=%s", eventDay))
}

// getAPIActivityPartitionPath returns a path for an API activity based on its event time
func (d *BaseDatastore) getAPIActivityPartitionPath(activity ocsf.APIActivity) string {
	eventDay := time.UnixMilli(activity.Time).Format("2006-01-02")
	return filepath.Join(BasepathActivities, fmt.Sprintf("event_day=%s", eventDay))
}

func filesExist(pattern string) bool {
	files, err := filepath.Glob(pattern)
	return err == nil && len(files) > 0
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
