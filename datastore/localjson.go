package datastore

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/Santiago-Labs/go-ocsf/clients/duckdb"
	"github.com/Santiago-Labs/go-ocsf/ocsf"
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/samsarahq/go/oops"
)

type localJsonDatastore struct {
	BaseDatastore
}

// localJsonDatastore implements the Datastore interface using local JSON files for storage.
// It provides methods to retrieve, save, and manage vulnerability findings in JSON format.
// The datastore maintains an in-memory index of finding IDs to file paths for quick lookups.
func NewLocalJsonDatastore() (Datastore, error) {
	s := &localJsonDatastore{}
	dbClient, err := duckdb.NewLocalClient(context.Background())
	if err != nil {
		return nil, oops.Wrapf(err, "failed to create local client")
	}

	basePatterns := map[string]string{
		"vulnerability_finding": fmt.Sprintf("%s/*/*.json", BasepathFindings),
		"api_activities":        fmt.Sprintf("%s/*/*.json", BasepathActivities),
	}

	fields := map[string][]arrow.Field{
		"vulnerability_finding": ocsf.VulnerabilityFindingFields,
		"api_activities":        ocsf.APIActivityFields,
	}

	var queries string
	queries += "INSTALL json; LOAD json; "
	for view, pattern := range basePatterns {
		selectFields := duckdb.GenerateDuckDBSelectFields(view, pattern, fields[view])
		columnsDict := duckdb.GenerateDuckDBColumnsDict(fields[view])
		if filesExist(pattern) {
			queries += fmt.Sprintf(`
				CREATE OR REPLACE VIEW %s AS
				SELECT %s FROM read_json_auto(
				'%s',
				ignore_errors=true,
				union_by_name=true,
				hive_partitioning=true,
				%s
				);`,
				view, selectFields, pattern, columnsDict,
			)
		} else {
			queries += duckdb.GenerateDuckDBNullView(view, fields[view])
		}
	}

	_, err = dbClient.Exec(queries)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to create view")
	}

	s.BaseDatastore = BaseDatastore{
		store: s,
		db:    dbClient,
	}

	if err := os.MkdirAll(Basepath, 0755); err != nil {
		return nil, oops.Wrapf(err, "failed to create directory")
	}

	if err := os.MkdirAll(BasepathActivities, 0755); err != nil {
		return nil, oops.Wrapf(err, "failed to create directory")
	}

	return s, nil
}

// GetFindingsFromFile retrieves all vulnerability findings from a specific file path.
// It reads the JSON file and parses it into a slice of vulnerability findings.
func (s *localJsonDatastore) GetFindingsFromFile(ctx context.Context, path string) ([]ocsf.VulnerabilityFinding, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to read JSON file from disk")
	}

	var findings []ocsf.VulnerabilityFinding

	if err := json.Unmarshal(data, &findings); err != nil {
		return nil, oops.Wrapf(err, "failed to parse JSON file")
	}

	return findings, nil
}

// WriteBatch creates a new JSON file for storing vulnerability findings.
// It marshals the findings into a JSON object and writes it to the specified file path.
func (s *localJsonDatastore) WriteBatch(ctx context.Context, findings []ocsf.VulnerabilityFinding, pathPrefix string) error {
	allFindings := findings

	var fullPath string
	if _, err := os.Stat(pathPrefix); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(pathPrefix, 0755); err != nil {
				return oops.Wrapf(err, "failed to create directory")
			}

			fullPath = filepath.Join(pathPrefix, fmt.Sprintf("%s.json", time.Now().Format("20060102T150405Z")))
		} else {
			return oops.Wrapf(err, "failed to check if directory exists")
		}
	} else {
		// Get all files in the directory
		files, err := filepath.Glob(filepath.Join(pathPrefix, "*.json"))
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
			fullPath = filepath.Join(pathPrefix, fmt.Sprintf("%s.json", time.Now().Format("20060102T150405Z")))
		}
	}

	jsonData, err := json.Marshal(allFindings)
	if err != nil {
		return oops.Wrapf(err, "failed to marshal findings to JSON")
	}

	if err := os.WriteFile(fullPath, jsonData, 0644); err != nil {
		return oops.Wrapf(err, "failed to write JSON to disk")
	}

	slog.Info("Wrote JSON file to disk", "path", fullPath, "findings", len(allFindings))

	return nil
}

func (s *localJsonDatastore) WriteAPIActivityBatch(ctx context.Context, activities []ocsf.APIActivity, pathPrefix string) error {
	allActivities := activities

	var fullPath string
	if _, err := os.Stat(pathPrefix); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(pathPrefix, 0755); err != nil {
				return oops.Wrapf(err, "failed to create directory")
			}

			fullPath = filepath.Join(pathPrefix, fmt.Sprintf("%s.json", time.Now().Format("20060102T150405Z")))
		} else {
			return oops.Wrapf(err, "failed to check if directory exists")
		}
	} else {
		files, err := filepath.Glob(filepath.Join(pathPrefix, "*.json"))
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
			fullPath = filepath.Join(pathPrefix, fmt.Sprintf("%s.json", time.Now().Format("20060102T150405Z")))
		}
	}

	jsonData, err := json.Marshal(allActivities)
	if err != nil {
		return oops.Wrapf(err, "failed to marshal activities to JSON")
	}
	if err := os.WriteFile(fullPath, jsonData, 0644); err != nil {
		return oops.Wrapf(err, "failed to write JSON to disk")
	}

	return nil

}

// GetAPIActivitiesFromFile retrieves all API activities from a specific file path.
// It reads the JSON file and parses it into a slice of API activities.
func (s *localJsonDatastore) GetAPIActivitiesFromFile(ctx context.Context, path string) ([]ocsf.APIActivity, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to read JSON file from disk")
	}

	var activities []ocsf.APIActivity

	if err := json.Unmarshal(data, &activities); err != nil {
		return nil, oops.Wrapf(err, "failed to parse JSON file")
	}

	return activities, nil
}
