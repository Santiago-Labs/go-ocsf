package datastore

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Santiago-Labs/go-ocsf/ocsf"
	goParquet "github.com/parquet-go/parquet-go"
	"github.com/samsarahq/go/oops"
)

type localParquetDatastore struct {
	findingIndex map[string]string
}

func NewLocalParquetDatastore() (Datastore, error) {
	s := &localParquetDatastore{
		findingIndex: make(map[string]string),
	}

	if err := os.MkdirAll(basepath, 0755); err != nil {
		return nil, oops.Wrapf(err, "failed to create directory")
	}

	ctx := context.Background()
	if err := s.buildFindingIndex(ctx); err != nil {
		return nil, oops.Wrapf(err, "failed to build complete finding index")
	}

	return s, nil
}

func (s *localParquetDatastore) buildFindingIndex(ctx context.Context) error {
	files, err := os.ReadDir(basepath)
	if err != nil {
		return oops.Wrapf(err, "failed to read local directory %s", basepath)
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".parquet") {
			continue
		}

		filePath := filepath.Join(basepath, file.Name())
		if err := s.loadFileIntoIndex(ctx, filePath); err != nil {
			slog.Error("failed to load parquet file into index, skipping", "file", filePath, "error", err)
			continue
		}
	}

	slog.Info("built finding index from local files", "count", len(s.findingIndex))
	return nil
}

func (s *localParquetDatastore) GetFinding(ctx context.Context, findingID string) (*ocsf.VulnerabilityFinding, error) {
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

func (s *localParquetDatastore) GetAllFindings(ctx context.Context, path string) ([]ocsf.VulnerabilityFinding, error) {
	findings, err := goParquet.ReadFile[ocsf.VulnerabilityFinding](path)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to read parquet file")
	}

	return findings, nil
}

func (s *localParquetDatastore) SaveFindings(ctx context.Context, findings []ocsf.VulnerabilityFinding) error {
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

	newFindingsFilename := fmt.Sprintf("%s.parquet", time.Now().Format("20060102T150405Z"))
	newFindingsPath := filepath.Join(basepath, newFindingsFilename)

	if len(newFindings) > 0 {
		if err := s.createFile(ctx, newFindingsPath, newFindings); err != nil {
			return oops.Wrapf(err, "failed to write new findings to parquet")
		}
	}

	updatedCount := 0
	for updatedFindingsPath, updatedFindings := range existingFindings {
		if err := s.updateFile(ctx, updatedFindingsPath, updatedFindings); err != nil {
			slog.Error("failed to update parquet file", "file", updatedFindingsPath, "error", err)
			// Fall back to creating a new file for these findings

			newFindingsFilename := fmt.Sprintf("%s.parquet", time.Now().Format("20060102T150405Z"))
			newFindingsPath := filepath.Join(basepath, newFindingsFilename)
			if err := s.createFile(ctx, newFindingsPath, updatedFindings); err != nil {
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

func (s *localParquetDatastore) createFile(ctx context.Context, filepath string, findings []ocsf.VulnerabilityFinding) error {
	err := goParquet.WriteFile(filepath, findings)
	if err != nil {
		return oops.Wrapf(err, "failed to write findings to parquet")
	}

	for _, f := range findings {
		s.findingIndex[f.FindingInfo.UID] = filepath
	}

	slog.Info("Wrote Parquet file to disk",
		"path", filepath,
		"findings", len(findings))
	return nil
}
func (s *localParquetDatastore) loadFileIntoIndex(ctx context.Context, path string) error {
	findings, err := s.GetAllFindings(ctx, path)
	if err != nil {
		return oops.Wrapf(err, "failed to read all findings from parquet file")
	}

	for _, finding := range findings {
		s.findingIndex[finding.FindingInfo.UID] = path
	}

	slog.Debug("indexed findings from file", "path", path, "count", len(findings))
	return nil
}

func (s *localParquetDatastore) updateFile(ctx context.Context, filePath string, updatedFindings []ocsf.VulnerabilityFinding) error {
	updateMap := make(map[string]ocsf.VulnerabilityFinding)
	for _, finding := range updatedFindings {
		updateMap[finding.FindingInfo.UID] = finding
	}

	allFindings, err := s.GetAllFindings(ctx, filePath)
	if err != nil {
		return oops.Wrapf(err, "failed to read all findings from parquet file")
	}

	for i, finding := range allFindings {
		if updated, exists := updateMap[finding.FindingInfo.UID]; exists {
			allFindings[i] = updated
			slog.Debug("updated finding in memory", "id", finding.FindingInfo.UID)
		}
	}

	err = goParquet.WriteFile(filePath, allFindings)
	if err != nil {
		return oops.Wrapf(err, "failed to write findings to parquet")
	}

	slog.Info("replaced parquet file on disk",
		"path", filePath,
		"findings", len(updatedFindings))

	return nil
}
