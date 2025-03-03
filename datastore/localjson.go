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
		findingIndex: make(map[string]string),
		fileIndex:    make(map[string]int),
		store:        s,
	}

	if err := os.MkdirAll(basepath, 0755); err != nil {
		return nil, oops.Wrapf(err, "failed to create directory")
	}

	ctx := context.Background()
	if err := s.buildFindingIndex(ctx); err != nil {
		slog.Warn("failed to build complete finding index", "error", err)
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
		VulnerabilityFindings []ocsf.VulnerabilityFinding `json:"vulnerability_findings"`
	}

	if err := json.Unmarshal(data, &findings); err != nil {
		return nil, oops.Wrapf(err, "failed to parse JSON file")
	}

	return findings.VulnerabilityFindings, nil
}

// WriteBatch creates a new JSON file for storing vulnerability findings.
// It marshals the findings into a JSON object and writes it to the specified file path.
// The datastore also updates its in-memory index of finding IDs to file paths.
func (s *localJsonDatastore) WriteBatch(ctx context.Context, findings []ocsf.VulnerabilityFinding, path *string) error {
	allFindings := findings
	if path == nil {
		newpath := filepath.Join(basepath, fmt.Sprintf("%s.json", time.Now().Format("20060102T150405Z")))
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
		"vulnerability_findings": allFindings,
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
	files, err := os.ReadDir(basepath)
	if err != nil {
		return oops.Wrapf(err, "failed to read local directory %s", basepath)
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		filePath := filepath.Join(basepath, file.Name())
		if err := s.loadFileIntoIndex(ctx, filePath); err != nil {
			slog.Error("failed to load json file into index, skipping", "file", filePath, "error", err)
			continue
		}
	}

	slog.Info("built finding index from local files", "count", len(s.BaseDatastore.findingIndex))
	return nil
}
