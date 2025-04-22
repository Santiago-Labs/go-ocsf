package datastore

import (
	"compress/gzip"
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
	currentFindingsPath := filepath.Join(BasepathFindings, fmt.Sprintf("%s.json.gz", time.Now().Format("20060102T150405Z")))
	if _, err := os.Stat(BasepathFindings); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(BasepathFindings, 0755); err != nil {
				return nil, oops.Wrapf(err, "failed to create directory")
			}
		}
	}

	_, err := os.Create(currentFindingsPath)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to create file")
	}

	currentActivitiesPath := filepath.Join(BasepathActivities, fmt.Sprintf("%s.json.gz", time.Now().Format("20060102T150405Z")))
	if _, err := os.Stat(currentActivitiesPath); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(BasepathActivities, 0755); err != nil {
				return nil, oops.Wrapf(err, "failed to create directory")
			}
		}
	}

	_, err = os.Create(currentActivitiesPath)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to create file")
	}

	s := &localJsonDatastore{
		currentFindingsPath:   filepath.Join(BasepathFindings, fmt.Sprintf("%s.json.gz", time.Now().Format("20060102T150405Z"))),
		currentActivitiesPath: filepath.Join(BasepathActivities, fmt.Sprintf("%s.json.gz", time.Now().Format("20060102T150405Z"))),
	}

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
// It reads the gzipped JSON file and parses it into a slice of vulnerability findings.
func (s *localJsonDatastore) GetFindingsFromFile(ctx context.Context, path string) ([]ocsf.VulnerabilityFinding, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to open gzip file")
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to create gzip reader")
	}
	defer gzipReader.Close()

	var findings []ocsf.VulnerabilityFinding
	if err := json.NewDecoder(gzipReader).Decode(&findings); err != nil {
		return nil, oops.Wrapf(err, "failed to parse gzipped JSON file")
	}

	return findings, nil
}

// WriteBatch creates a new JSON file for storing vulnerability findings.
// It marshals the findings into a JSON object and writes it to the specified file path.
func (s *localJsonDatastore) WriteBatch(ctx context.Context, findings []ocsf.VulnerabilityFinding) error {
	allFindings := findings

	fileFindings, err := s.GetFindingsFromFile(ctx, s.currentFindingsPath)
	if err != nil {
		return oops.Wrapf(err, "failed to get existing findings from disk")
	}

	allFindings = append(allFindings, fileFindings...)

	jsonData, err := json.Marshal(allFindings)
	if err != nil {
		return oops.Wrapf(err, "failed to marshal findings to JSON")
	}

	file, err := os.OpenFile(s.currentFindingsPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return oops.Wrapf(err, "failed to create gzip file")
	}
	defer file.Close()

	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()

	if _, err := gzipWriter.Write(jsonData); err != nil {
		return oops.Wrapf(err, "failed to write gzip data")
	}

	slog.Info("Wrote JSON file to disk", "path", s.currentFindingsPath, "findings", len(allFindings))

	return nil
}

func (s *localJsonDatastore) WriteAPIActivityBatch(ctx context.Context, activities []ocsf.APIActivity) error {
	allActivities := activities

	fileActivities, err := s.GetAPIActivitiesFromFile(ctx, s.currentActivitiesPath)
	if err != nil {
		return oops.Wrapf(err, "failed to get existing activities from disk")
	}

	allActivities = append(allActivities, fileActivities...)

	jsonData, err := json.Marshal(allActivities)
	if err != nil {
		return oops.Wrapf(err, "failed to marshal activities to JSON")
	}

	file, err := os.OpenFile(s.currentActivitiesPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return oops.Wrapf(err, "failed to create gzip file")
	}
	defer file.Close()

	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()

	if _, err := gzipWriter.Write(jsonData); err != nil {
		return oops.Wrapf(err, "failed to write gzip data")
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
	file, err := os.Open(path)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to open gzip file")
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to create gzip reader")
	}
	defer gzipReader.Close()

	var activities []ocsf.APIActivity
	if err := json.NewDecoder(gzipReader).Decode(&activities); err != nil {
		return nil, oops.Wrapf(err, "failed to parse gzipped JSON file")
	}

	return activities, nil
}
