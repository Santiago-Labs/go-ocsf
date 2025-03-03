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
	BaseDatastore
}

// NewLocalParquetDatastore creates a new local Parquet datastore.
// It initializes an in-memory index of finding IDs to file paths.
func NewLocalParquetDatastore() (Datastore, error) {
	s := &localParquetDatastore{}
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
		return nil, oops.Wrapf(err, "failed to build complete finding index")
	}

	return s, nil
}

// buildFindingIndex builds the datastore's in-memory index of finding IDs to file paths.
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

	slog.Info("built finding index from local files", "count", len(s.BaseDatastore.findingIndex))
	return nil
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
func (s *localParquetDatastore) WriteBatch(ctx context.Context, findings []ocsf.VulnerabilityFinding, path *string) error {
	allFindings := findings
	if path == nil {
		newpath := filepath.Join(basepath, fmt.Sprintf("%s.parquet", time.Now().Format("20060102T150405Z")))
		path = &newpath
	} else {
		var err error
		allFindings, err = s.GetFindingsFromFile(ctx, *path)
		if err != nil {
			return oops.Wrapf(err, "failed to get existing findings from disk")
		}

		allFindings = append(allFindings, findings...)
	}

	err := goParquet.WriteFile(*path, allFindings)
	if err != nil {
		return oops.Wrapf(err, "failed to write findings to parquet")
	}

	for _, f := range allFindings {
		s.BaseDatastore.findingIndex[f.FindingInfo.UID] = *path
	}
	s.BaseDatastore.fileIndex[*path] = len(allFindings)

	slog.Info("Wrote Parquet file to disk",
		"path", *path,
		"findings", len(allFindings),
	)
	return nil
}
