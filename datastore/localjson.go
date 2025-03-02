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
	findingIndex map[string]string
}

func NewLocalJsonDatastore() (Datastore, error) {
	s := &localJsonDatastore{
		findingIndex: make(map[string]string),
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

func (s *localJsonDatastore) GetFinding(ctx context.Context, findingID string) (*ocsf.VulnerabilityFinding, error) {
	filePath, exists := s.findingIndex[findingID]
	if !exists {
		return nil, ErrNotFound
	}

	findings, err := s.GetAllFindings(ctx, filePath)
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

func (s *localJsonDatastore) GetAllFindings(ctx context.Context, path string) ([]ocsf.VulnerabilityFinding, error) {
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

func (s *localJsonDatastore) SaveFindings(ctx context.Context, findings []ocsf.VulnerabilityFinding) error {
	newFindings := []ocsf.VulnerabilityFinding{}
	existingFindings := make(map[string][]ocsf.VulnerabilityFinding)

	for _, finding := range findings {
		filePath, exists := s.findingIndex[finding.FindingInfo.UID]
		if !exists {
			newFindings = append(newFindings, finding)
		} else {
			existingFindings[filePath] = append(existingFindings[filePath], finding)
		}
	}

	newFindingsFilename := fmt.Sprintf("%s.json", time.Now().Format("20060102T150405Z"))
	newFindingsKey := filepath.Join(basepath, newFindingsFilename)
	if len(newFindings) > 0 {
		if err := s.createFile(ctx, newFindingsKey, newFindings); err != nil {
			return oops.Wrapf(err, "failed to write new findings to parquet")
		}
	}

	updatedCount := 0
	for updatedFindingsPath, updatedFindings := range existingFindings {
		if err := s.updateFile(ctx, updatedFindingsPath, updatedFindings); err != nil {
			slog.Error("failed to update json file", "file", updatedFindingsPath, "error", err)
			// Fall back to creating a new file for these findings

			newFindingsFilename := fmt.Sprintf("%s.json", time.Now().Format("20060102T150405Z"))
			newFindingsKey := filepath.Join(basepath, newFindingsFilename)
			if err := s.createFile(ctx, newFindingsKey, updatedFindings); err != nil {
				return oops.Wrapf(err, "failed to write fallback parquet file")
			}
		} else {
			updatedCount += len(updatedFindings)
		}
	}

	slog.Info("saved findings",
		"new", len(newFindings),
		"updated", updatedCount)

	return nil
}

func (s *localJsonDatastore) createFile(ctx context.Context, path string, findings []ocsf.VulnerabilityFinding) error {
	outerSchema := map[string]interface{}{
		"vulnerability_findings": findings,
	}
	jsonData, err := json.Marshal(outerSchema)
	if err != nil {
		return oops.Wrapf(err, "failed to marshal findings to JSON")
	}

	if err := os.WriteFile(path, jsonData, 0644); err != nil {
		return oops.Wrapf(err, "failed to write JSON to disk")
	}

	for _, finding := range findings {
		s.findingIndex[finding.FindingInfo.UID] = path
	}

	slog.Info("Wrote JSON file to disk", "path", path)

	return nil
}

func (s *localJsonDatastore) loadFileIntoIndex(ctx context.Context, path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return oops.Wrapf(err, "failed to read JSON file from disk")
	}

	var findings struct {
		VulnerabilityFindings []ocsf.VulnerabilityFinding `json:"vulnerability_findings"`
	}

	if err := json.Unmarshal(data, &findings); err != nil {
		return oops.Wrapf(err, "failed to parse JSON file")
	}

	for _, finding := range findings.VulnerabilityFindings {
		s.findingIndex[finding.FindingInfo.UID] = path
	}

	return nil
}

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

	slog.Info("built finding index from local files", "count", len(s.findingIndex))
	return nil
}

func (s *localJsonDatastore) updateFile(ctx context.Context, filePath string, updatedFindings []ocsf.VulnerabilityFinding) error {
	updateMap := make(map[string]ocsf.VulnerabilityFinding)
	for _, finding := range updatedFindings {
		updateMap[finding.FindingInfo.UID] = finding
	}

	fullPath := filepath.Join("/findings", filePath)
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return oops.Wrapf(err, "failed to read JSON file from disk")
	}

	var findings struct {
		VulnerabilityFindings []ocsf.VulnerabilityFinding `json:"vulnerability_findings"`
	}

	if err := json.Unmarshal(data, &findings); err != nil {
		return oops.Wrapf(err, "failed to parse JSON file")
	}

	allFindings := findings.VulnerabilityFindings

	for i, finding := range allFindings {
		if updated, exists := updateMap[finding.FindingInfo.UID]; exists {
			allFindings[i] = updated
			slog.Debug("updated finding in memory", "id", finding.FindingInfo.UID)
		}
	}

	outerSchema := map[string]interface{}{
		"vulnerability_findings": allFindings,
	}
	jsonData, err := json.Marshal(outerSchema)
	if err != nil {
		return oops.Wrapf(err, "failed to marshal findings to JSON")
	}

	if err := os.WriteFile(fullPath, jsonData, 0644); err != nil {
		return oops.Wrapf(err, "failed to write updated JSON to disk")
	}

	slog.Info("updated JSON file on disk",
		"path", fullPath,
		"size", len(jsonData),
		"findings", len(updatedFindings))

	return nil
}
