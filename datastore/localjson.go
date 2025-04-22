package datastore

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/Santiago-Labs/go-ocsf/ocsf"
	"github.com/samsarahq/go/oops"
)

type localJsonDatastore struct {
	BaseDatastore

	currentFindingsPath   string
	currentActivitiesPath string
}

// localJsonDatastore implements the Datastore interface using local JSON files for storage.
// It provides methods to retrieve, save, and manage vulnerability findings in JSON format.
func NewLocalJsonDatastore(ctx context.Context) (Datastore, error) {
	if err := os.MkdirAll(BasepathFindings, 0755); err != nil {
		return nil, oops.Wrapf(err, "failed to create directory")
	}

	if err := os.MkdirAll(BasepathActivities, 0755); err != nil {
		return nil, oops.Wrapf(err, "failed to create directory")
	}

	s := &localJsonDatastore{}

	s.BaseDatastore = BaseDatastore{
		store: s,
	}

	return s, nil
}

// GetFindingsFromFile retrieves all vulnerability findings from a specific file path.
// It reads the gzipped JSON file and parses it into a slice of vulnerability findings.
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
func (s *localJsonDatastore) WriteBatch(ctx context.Context, findings []ocsf.VulnerabilityFinding) error {
	allFindings := findings

	if s.currentFindingsPath == "" {
		s.currentFindingsPath = filepath.Join(BasepathFindings, fmt.Sprintf("%s.json", time.Now().Format("20060102T150405Z")))
	} else {
		fileFindings, err := s.GetFindingsFromFile(ctx, s.currentFindingsPath)
		if err != nil {
			return oops.Wrapf(err, "failed to get existing findings from disk")
		}

		allFindings = append(allFindings, fileFindings...)
	}

	jsonData, err := json.Marshal(allFindings)
	if err != nil {
		return oops.Wrapf(err, "failed to marshal findings to JSON")
	}

	if err := os.WriteFile(s.currentFindingsPath, jsonData, 0644); err != nil {
		return oops.Wrapf(err, "failed to write JSON to disk")
	}

	slog.Info("Wrote JSON file to disk", "path", s.currentFindingsPath, "findings", len(allFindings))

	return nil
}

func (s *localJsonDatastore) WriteAPIActivityBatch(ctx context.Context, activities []ocsf.APIActivity) error {
	allActivities := activities

	if s.currentActivitiesPath == "" {
		s.currentActivitiesPath = filepath.Join(BasepathActivities, fmt.Sprintf("%s.json", time.Now().Format("20060102T150405Z")))
	} else {
		fileActivities, err := s.GetAPIActivitiesFromFile(ctx, s.currentActivitiesPath)
		if err != nil {
			return oops.Wrapf(err, "failed to get existing activities from disk")
		}

		allActivities = append(allActivities, fileActivities...)
	}

	jsonData, err := json.Marshal(allActivities)
	if err != nil {
		return oops.Wrapf(err, "failed to marshal findings to JSON")
	}

	if err := os.WriteFile(s.currentActivitiesPath, jsonData, 0644); err != nil {
		return oops.Wrapf(err, "failed to write JSON to disk")
	}
	slog.Info("Wrote JSON file to disk",
		"path", s.currentActivitiesPath,
		"activities", len(allActivities),
	)

	return nil

}

// GetAPIActivitiesFromFile retrieves all API activities from a specific file path.
// It reads the gzipped JSON file and parses it into a slice of API activities.
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
