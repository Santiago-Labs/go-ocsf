package datastore

import (
	"context"
	"log/slog"

	"github.com/Santiago-Labs/go-ocsf/ocsf"
	"github.com/samsarahq/go/oops"
)

type BaseDatastore struct {
	findingIndex map[string]string // findingID -> filePath
	fileIndex    map[string]int    // filePath -> len(findings)

	store Datastore
}

// GetFinding returns a single finding from the datastore or ErrNotFound if the finding is not found
func (d *BaseDatastore) GetFinding(ctx context.Context, findingID string) (*ocsf.VulnerabilityFinding, error) {
	filePath, exists := d.findingIndex[findingID]
	if !exists {
		return nil, ErrNotFound
	}

	findings, err := d.store.GetFindingsFromFile(ctx, filePath)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to get all findings")
	}

	for _, finding := range findings {
		if finding.FindingInfo.UID == findingID {
			return &finding, nil
		}
	}

	return nil, ErrNotFound
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

	slog.Info("saved findings", "findings", len(findings))

	return nil
}

func (d *BaseDatastore) getPathWithSpace(batchSize int) *string {
	for path, count := range d.fileIndex {
		if (count+batchSize)*avgFindingSize < maxFileSize {
			return &path
		}
	}

	return nil
}

func (d *BaseDatastore) loadFileIntoIndex(ctx context.Context, key string) error {
	findings, err := d.store.GetFindingsFromFile(ctx, key)
	if err != nil {
		return oops.Wrapf(err, "failed to read all findings from parquet file")
	}

	for _, finding := range findings {
		d.findingIndex[finding.FindingInfo.UID] = key
	}
	d.fileIndex[key] = len(findings)

	slog.Debug("indexed findings from file", "key", key, "count", len(findings))
	return nil
}
