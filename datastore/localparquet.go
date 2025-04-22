package datastore

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/Santiago-Labs/go-ocsf/ocsf"
	goParquet "github.com/parquet-go/parquet-go"
	"github.com/samsarahq/go/oops"
)

type localParquetDatastore struct {
	BaseDatastore

	currentFindingsPath   string
	currentActivitiesPath string
}

// NewLocalParquetDatastore creates a new local Parquet datastore.
func NewLocalParquetDatastore(ctx context.Context) (Datastore, error) {
	if err := os.MkdirAll(BasepathFindings, 0755); err != nil {
		return nil, oops.Wrapf(err, "failed to create directory")
	}

	if err := os.MkdirAll(BasepathActivities, 0755); err != nil {
		return nil, oops.Wrapf(err, "failed to create directory")
	}

	s := &localParquetDatastore{}

	s.BaseDatastore = BaseDatastore{
		store:                  s,
		findingsTableName:      "vulnerability_finding",
		apiActivitiesTableName: "api_activities",
	}

	return s, nil
}

// GetFindingsFromFile retrieves all vulnerability findings from a specific file path.
// It reads the Parquet file and parses it into a slice of vulnerability findings.
func (s *localParquetDatastore) GetFindingsFromFile(ctx context.Context, path string) ([]ocsf.VulnerabilityFinding, error) {
	findings, err := goParquet.ReadFile[ocsf.VulnerabilityFinding](path)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to read parquet file")
	}

	return findings, nil
}

// createFile creates a new Parquet file for storing vulnerability findings.
// It writes the findings to the specified file path.
func (s *localParquetDatastore) WriteBatch(ctx context.Context, findings []ocsf.VulnerabilityFinding) error {
	allFindings := findings

	if s.currentFindingsPath == "" {
		s.currentFindingsPath = filepath.Join(BasepathFindings, fmt.Sprintf("%s.parquet.gz", time.Now().Format("20060102T150405Z")))
	} else {
		fileFindings, err := s.GetFindingsFromFile(ctx, s.currentFindingsPath)
		if err != nil {
			return oops.Wrapf(err, "failed to get existing findings from disk")
		}
		allFindings = append(allFindings, fileFindings...)
	}

	err := goParquet.WriteFile(s.currentFindingsPath, allFindings, goParquet.Compression(&goParquet.Gzip))
	if err != nil {
		return oops.Wrapf(err, "failed to write activities to parquet")
	}

	slog.Info("Wrote parquet file to disk",
		"path", s.currentFindingsPath,
		"findings", len(allFindings),
	)

	return nil
}

func (s *localParquetDatastore) WriteAPIActivityBatch(ctx context.Context, activities []ocsf.APIActivity) error {
	allActivities := activities

	if s.currentActivitiesPath == "" {
		s.currentActivitiesPath = filepath.Join(BasepathActivities, fmt.Sprintf("%s.parquet.gz", time.Now().Format("20060102T150405Z")))
	} else {
		fileActivities, err := s.GetAPIActivitiesFromFile(ctx, s.currentActivitiesPath)
		if err != nil {
			return oops.Wrapf(err, "failed to get existing activities from disk")
		}

		allActivities = append(allActivities, fileActivities...)
	}

	err := goParquet.WriteFile(s.currentActivitiesPath, allActivities, goParquet.Compression(&goParquet.Gzip))
	if err != nil {
		return oops.Wrapf(err, "failed to write activities to parquet")
	}

	slog.Info("Wrote parquet file to disk",
		"path", s.currentActivitiesPath,
		"activities", len(allActivities),
	)

	return nil
}

func (s *localParquetDatastore) GetAPIActivitiesFromFile(ctx context.Context, path string) ([]ocsf.APIActivity, error) {
	if _, exists := os.Stat(path); os.IsNotExist(exists) {
		return []ocsf.APIActivity{}, nil
	}
	activities, err := goParquet.ReadFile[ocsf.APIActivity](path)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to read parquet file")
	}

	return activities, nil
}
