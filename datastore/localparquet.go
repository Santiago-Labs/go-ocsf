package datastore

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"reflect"
	"time"

	goParquet "github.com/parquet-go/parquet-go"
	"github.com/samsarahq/go/oops"
)

type localParquetDatastore[T any] struct {
	BaseDatastore[T]

	currentPath string
	basepath    string
}

// NewLocalParquetDatastore creates a new local Parquet datastore.
func NewLocalParquetDatastore[T any](ctx context.Context) (Datastore[T], error) {

	typeName := reflect.TypeOf((*T)(nil)).Elem().Name()
	if err := os.MkdirAll(basepaths[typeName], 0755); err != nil {
		return nil, oops.Wrapf(err, "failed to create directory")
	}

	s := &localParquetDatastore[T]{
		basepath: basepaths[typeName],
	}

	s.BaseDatastore = BaseDatastore[T]{
		store: s,
	}

	return s, nil
}

// GetItemsFromFile retrieves all ocsf data from a specific file path.
// It reads the Parquet file and parses it into a slice of ocsf data.
func (s *localParquetDatastore[T]) GetItemsFromFile(ctx context.Context, path string) ([]T, error) {
	items, err := goParquet.ReadFile[T](path)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to read parquet file")
	}

	return items, nil
}

// createFile creates a new Parquet file for storing ocsf data.
// It writes the data to the specified file path.
func (s *localParquetDatastore[T]) WriteBatch(ctx context.Context, items []T) error {
	allItems := items

	if s.currentPath == "" {
		s.currentPath = filepath.Join(s.basepath, fmt.Sprintf("%s.parquet.gz", time.Now().Format("20060102T150405Z")))
	} else {
		fileItems, err := s.GetItemsFromFile(ctx, s.currentPath)
		if err != nil {
			return oops.Wrapf(err, "failed to get existing items from disk")
		}
		allItems = append(allItems, fileItems...)
	}

	err := goParquet.WriteFile(s.currentPath, allItems, goParquet.Compression(&goParquet.Gzip))
	if err != nil {
		return oops.Wrapf(err, "failed to write to parquet")
	}

	slog.Info("Wrote parquet file to disk",
		"path", s.currentPath,
		"items", len(allItems),
	)

	return nil
}
