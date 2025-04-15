package datastore

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
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

func (s *localJsonDatastore) buildActivityIndex(ctx context.Context) error {
	files, err := os.ReadDir(basepathActivities)
	if err != nil {
		return oops.Wrapf(err, "failed to read local directory %s", basepathActivities)
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		filePath := filepath.Join(basepathActivities, file.Name())
		if err := s.loadActivityFileIntoIndex(ctx, filePath); err != nil {
			slog.Error("failed to load json file into index, skipping", "file", filePath, "error", err)
			continue
		}
	}

	slog.Info("built activity index from local files", "count", len(s.BaseDatastore.activityIndex))
	return nil
}

func (s *localJsonDatastore) WriteAPIActivityBatch(ctx context.Context, activities []ocsf.APIActivity, path *string) error {
	allActivities := activities
	if path == nil {
		newpath := filepath.Join(basepathActivities, fmt.Sprintf("%s.json", time.Now().Format("20060102T150405Z")))
		path = &newpath
	} else {
		var err error
		allActivities, err = s.GetAPIActivitiesFromFile(ctx, *path)
		if err != nil {
			return oops.Wrapf(err, "failed to get existing activities from disk")
		}

		allActivities = append(allActivities, activities...)
	}

	outerSchema := map[string]interface{}{
		"api_activities": allActivities,
	}
	jsonData, err := json.Marshal(outerSchema)
	if err != nil {
		return oops.Wrapf(err, "failed to marshal activities to JSON")
	}
	if err := os.WriteFile(*path, jsonData, 0644); err != nil {
		return oops.Wrapf(err, "failed to write JSON to disk")
	}

	for _, activity := range allActivities {
		s.BaseDatastore.activityIndex[*activity.Metadata.CorrelationUID] = *path
	}
	s.BaseDatastore.activityFileIndex[*path] = len(allActivities)

	return nil

}

// GetAPIActivitiesFromFile retrieves all API activities from a specific file path.
// It reads the JSON file and parses it into a slice of API activities.
func (s *localJsonDatastore) GetAPIActivitiesFromFile(ctx context.Context, path string) ([]ocsf.APIActivity, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to read JSON file from disk")
	}

	var activities struct {
		APIActivities []ocsf.APIActivity `json:"api_activities"`
	}

	if err := json.Unmarshal(data, &activities); err != nil {
		return nil, oops.Wrapf(err, "failed to parse JSON file")
	}

	return activities.APIActivities, nil
}
