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
}

// NewLocalParquetDatastore creates a new local Parquet datastore.
// It initializes an in-memory index of finding IDs to file paths.
func NewLocalParquetDatastore(ctx context.Context) (Datastore, error) {
	s := &localParquetDatastore{}

	s.BaseDatastore = BaseDatastore{
		store:                  s,
		findingsTableName:      "vulnerability_finding",
		apiActivitiesTableName: "api_activities",
	}

	if err := os.MkdirAll(BasepathFindings, 0755); err != nil {
		return nil, oops.Wrapf(err, "failed to create directory")
	}
	if err := os.MkdirAll(BasepathActivities, 0755); err != nil {
		return nil, oops.Wrapf(err, "failed to create directory")
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
// It writes the findings to the specified file path and updates the datastore's in-memory index.
func (s *localParquetDatastore) WriteBatch(ctx context.Context, findings []ocsf.VulnerabilityFinding, pathPrefix string) error {
	allFindings := findings

	var fullPath string
	if _, err := os.Stat(pathPrefix); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(pathPrefix, 0755); err != nil {
				return oops.Wrapf(err, "failed to create directory")
			}

			fullPath = filepath.Join(pathPrefix, fmt.Sprintf("%s.parquet.gz", time.Now().Format("20060102T150405Z")))
		} else {
			return oops.Wrapf(err, "failed to check if directory exists")
		}
	} else {
		files, err := filepath.Glob(filepath.Join(pathPrefix, "*.parquet.gz"))
		if err != nil {
			return oops.Wrapf(err, "failed to get files from directory")
		}

		if len(files) > 0 {
			for _, file := range files {
				fileFindings, err := s.GetFindingsFromFile(ctx, file)
				if err != nil {
					return oops.Wrapf(err, "failed to get existing findings from disk")
				}

				allFindings = append(allFindings, fileFindings...)

				fullPath = file
			}
		} else {
			fullPath = filepath.Join(pathPrefix, fmt.Sprintf("%s.parquet.gz", time.Now().Format("20060102T150405Z")))
		}
	}

	err := goParquet.WriteFile(fullPath, allFindings, goParquet.Compression(&goParquet.Gzip))
	if err != nil {
		return oops.Wrapf(err, "failed to write findings to parquet")
	}

	slog.Info("Wrote parquet file to disk",
		"path", fullPath,
		"findings", len(allFindings),
	)

	return nil
}

func (s *localParquetDatastore) WriteAPIActivityBatch(ctx context.Context, activities []ocsf.APIActivity, pathPrefix string) error {
	allActivities := activities

	var fullPath string
	if _, err := os.Stat(pathPrefix); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(pathPrefix, 0755); err != nil {
				return oops.Wrapf(err, "failed to create directory")
			}

			fullPath = filepath.Join(pathPrefix, fmt.Sprintf("%s.parquet.gz", time.Now().Format("20060102T150405Z")))
		} else {
			return oops.Wrapf(err, "failed to check if directory exists")
		}
	} else {
		files, err := filepath.Glob(filepath.Join(pathPrefix, "*.parquet.gz"))
		if err != nil {
			return oops.Wrapf(err, "failed to get files from directory")
		}

		if len(files) > 0 {
			for _, file := range files {
				fileActivities, err := s.GetAPIActivitiesFromFile(ctx, file)
				if err != nil {
					return oops.Wrapf(err, "failed to get existing activities from disk")
				}

				allActivities = append(allActivities, fileActivities...)

				fullPath = file
			}
		} else {
			fullPath = filepath.Join(pathPrefix, fmt.Sprintf("%s.parquet.gz", time.Now().Format("20060102T150405Z")))
		}
	}

	err := goParquet.WriteFile(fullPath, allActivities, goParquet.Compression(&goParquet.Gzip))
	if err != nil {
		return oops.Wrapf(err, "failed to write activities to parquet")
	}

	slog.Info("Wrote parquet file to disk",
		"path", fullPath,
		"activities", len(allActivities),
	)

	return nil
}

func (s *localParquetDatastore) GetAPIActivitiesFromFile(ctx context.Context, path string) ([]ocsf.APIActivity, error) {
	activities, err := goParquet.ReadFile[ocsf.APIActivity](path)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to read parquet file")
	}

	return activities, nil
}
