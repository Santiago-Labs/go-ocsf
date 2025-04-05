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
}

// localJsonDatastore implements the Datastore interface using local JSON files for storage.
// It provides methods to retrieve, save, and manage vulnerability findings in JSON format.
// The datastore maintains an in-memory index of finding IDs to file paths for quick lookups.
func NewLocalJsonDatastore() (Datastore, error) {
	s := &localJsonDatastore{}
	s.BaseDatastore = BaseDatastore{
		findingIndex:      make(map[string]string),
		fileIndex:         make(map[string]int),
		activityIndex:     make(map[string]string),
		activityFileIndex: make(map[string]int),
		store:             s,
	}

	if err := os.MkdirAll(Basepath, 0755); err != nil {
		return nil, oops.Wrapf(err, "failed to create directory")
	}

	if err := os.MkdirAll(BasepathActivities, 0755); err != nil {
		return nil, oops.Wrapf(err, "failed to create directory")
	}

	ctx := context.Background()
	if err := s.buildFindingIndex(ctx); err != nil {
		slog.Warn("failed to build complete finding index", "error", err)
	}

	if err := s.buildActivityIndex(ctx); err != nil {
		slog.Warn("failed to build complete activity index", "error", err)
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

	var findings struct {
		VulnerabilityFindings []ocsf.VulnerabilityFinding `json:"vulnerability_finding"`
	}

	if err := json.Unmarshal(data, &findings); err != nil {
		return nil, oops.Wrapf(err, "failed to parse JSON file")
	}

	return findings.VulnerabilityFindings, nil
}

// WriteBatch creates a new JSON file for storing vulnerability findings.
// It marshals the findings into a JSON object and writes it to the specified file path.
func (s *localJsonDatastore) WriteBatch(ctx context.Context, findings []ocsf.VulnerabilityFinding, path *string) error {
	allFindings := findings
	if path == nil {
		newpath := filepath.Join(Basepath, fmt.Sprintf("%s.json", time.Now().Format("20060102T150405Z")))
		path = &newpath
	} else {
		var err error
		allFindings, err = s.GetFindingsFromFile(ctx, *path)
		if err != nil {
			return oops.Wrapf(err, "failed to get existing findings from disk")
		}

		allFindings = append(allFindings, findings...)
	}

	outerSchema := map[string]interface{}{
		"vulnerability_finding": allFindings,
	}
	jsonData, err := json.Marshal(outerSchema)
	if err != nil {
		return oops.Wrapf(err, "failed to marshal findings to JSON")
	}

	if err := os.WriteFile(*path, jsonData, 0644); err != nil {
		return oops.Wrapf(err, "failed to write JSON to disk")
	}

	for _, finding := range allFindings {
		s.BaseDatastore.findingIndex[finding.FindingInfo.UID] = *path
	}
	s.BaseDatastore.fileIndex[*path] = len(allFindings)

	slog.Info("Wrote JSON file to disk", "path", *path, "findings", len(allFindings))

	return nil
}

// buildFindingIndex builds the datastore's in-memory index of finding IDs to file paths.
// It reads all JSON files in the base directory and parses them into a slice of vulnerability findings.
func (s *localJsonDatastore) buildFindingIndex(ctx context.Context) error {
	files, err := os.ReadDir(BasepathFindings)
	if err != nil {
		return oops.Wrapf(err, "failed to read local directory %s", BasepathFindings)
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		filePath := filepath.Join(BasepathFindings, file.Name())
		if err := s.loadFileIntoIndex(ctx, filePath); err != nil {
			slog.Error("failed to load json file into index, skipping", "file", filePath, "error", err)
			continue
		}
	}

	slog.Info("built finding index from local files", "count", len(s.BaseDatastore.findingIndex))
	return nil
}

func (s *localJsonDatastore) buildActivityIndex(ctx context.Context) error {
	files, err := os.ReadDir(BasepathActivities)
	if err != nil {
		return oops.Wrapf(err, "failed to read local directory %s", BasepathActivities)
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		filePath := filepath.Join(BasepathActivities, file.Name())
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
		newpath := filepath.Join(BasepathActivities, fmt.Sprintf("%s.json", time.Now().Format("20060102T150405Z")))
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
