package datastore

import (
	"context"
	"log/slog"

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
	finding := ocsf.VulnerabilityFinding{}
	rows, err := d.db.Queryx("SELECT * FROM vulnerability_findings WHERE uid = ?", findingID)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to get finding")
	}

	if !rows.Next() {
		return nil, ErrNotFound
	}

	if err := rows.StructScan(&finding); err != nil {
		return nil, oops.Wrapf(err, "failed to scan finding")
	}

	return &finding, nil
}

// SaveFindings saves a batch of findings to the datastore. Datastore implementations handle file formats.
func (d *BaseDatastore) SaveFindings(ctx context.Context, findings []ocsf.VulnerabilityFinding) error {
	var (
		currentBatch     []ocsf.VulnerabilityFinding
		currentBatchSize int
	)

	for _, finding := range findings {
		if currentBatchSize > maxFileSize && len(currentBatch) > 0 {
			pathWithSpace := d.getPathWithSpace(len(currentBatch))
			if err := d.store.WriteBatch(ctx, currentBatch, pathWithSpace); err != nil {
				return err
			}
			currentBatch = []ocsf.VulnerabilityFinding{}
			currentBatchSize = 0
		}

		currentBatch = append(currentBatch, finding)
		currentBatchSize += avgFindingSize
	}

	if len(currentBatch) > 0 {
		pathWithSpace := d.getPathWithSpace(len(currentBatch))
		if err := d.store.WriteBatch(ctx, currentBatch, pathWithSpace); err != nil {
			return err
		}
	}

	slog.Info("upserted findings", "findings", len(findings))

	return nil
}

func (d *BaseDatastore) getPathWithSpace(batchSize int) *string {
	return nil
}

func (d *BaseDatastore) getAPIActivityPathWithSpace(batchSize int) *string {
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
