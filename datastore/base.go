package datastore

import (
	"context"
	"log/slog"
	"path/filepath"

	"github.com/Santiago-Labs/go-ocsf/ocsf"
	"github.com/jmoiron/sqlx"
	"github.com/samsarahq/go/oops"
)

type BaseDatastore struct {
	store Datastore

	db *sqlx.DB
}

// GetFinding returns a single finding from the datastore or ErrNotFound if the finding is not found
func (d *BaseDatastore) GetFinding(ctx context.Context, findingID string) (*ocsf.VulnerabilityFinding, error) {
	findingDTO := ocsf.VulnerabilityFindingDTO{}
	rows, err := d.db.Queryx("SELECT * FROM vulnerability_finding WHERE finding_info.uid = ?", findingID)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to get finding")
	}

	if !rows.Next() {
		return nil, ErrNotFound
	}

	if err := rows.StructScan(&findingDTO); err != nil {
		return nil, oops.Wrapf(err, "failed to scan finding")
	}

	finding, err := findingDTO.ToStruct()
	if err != nil {
		return nil, oops.Wrapf(err, "failed to convert struct")
	}

	return finding, nil
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
			return err
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
			return err
		}
	}

	slog.Info("upserted activities", "activities", len(activities))

	return nil
}

// GetAPIActivity returns a single activity from the datastore or ErrNotFound if the activity is not found
func (d *BaseDatastore) GetAPIActivity(ctx context.Context, activityID string) (*ocsf.APIActivity, error) {
	activity := ocsf.APIActivity{}
	rows, err := d.db.Queryx("SELECT * FROM api_activities WHERE uid = ?", activityID)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to get activity")
	}

	if !rows.Next() {
		return nil, ErrNotFound
	}

	if err := rows.StructScan(&activity); err != nil {
		return nil, oops.Wrapf(err, "failed to scan activity")
	}

	return &activity, nil
}

// getPartitionPath returns a path for a finding based on its event time
func (d *BaseDatastore) getPartitionPath(finding ocsf.VulnerabilityFinding) string {
	eventDay := finding.Time.Format("2006-01-02")
	return filepath.Join(BasepathFindings, eventDay)
}

// getAPIActivityPartitionPath returns a path for an API activity based on its event time
func (d *BaseDatastore) getAPIActivityPartitionPath(activity ocsf.APIActivity) string {
	eventDay := activity.Time.Format("2006-01-02")
	return filepath.Join(BasepathActivities, eventDay)
}

func filesExist(pattern string) bool {
	files, err := filepath.Glob(pattern)
	return err == nil && len(files) > 0
}
