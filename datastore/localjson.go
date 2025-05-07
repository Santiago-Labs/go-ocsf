package datastore

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"reflect"
	"time"

	"github.com/samsarahq/go/oops"
)

type localJsonDatastore[T any] struct {
	BaseDatastore[T]

	currentPath string
	basepath    string
}

// localJsonDatastore implements the Datastore interface using local JSON files for storage.
// It provides methods to retrieve, save, and manage ocsf data in JSON format.
func NewLocalJsonDatastore[T any](ctx context.Context) (Datastore[T], error) {

	typeName := reflect.TypeOf((*T)(nil)).Elem().Name()
	if err := os.MkdirAll(basepaths[typeName], 0755); err != nil {
		return nil, oops.Wrapf(err, "failed to create directory")
	}

	s := &localJsonDatastore[T]{
		basepath: basepaths[typeName],
	}

	s.BaseDatastore = BaseDatastore[T]{
		store: s,
	}

	return s, nil
}

// GetItemsFromFile retrieves all ocsf data from a specific file path.
// It reads the gzipped JSON file and parses it into a slice of ocsf data.
func (s *localJsonDatastore[T]) GetItemsFromFile(ctx context.Context, path string) ([]T, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to read JSON file from disk")
	}

	var items []T
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, oops.Wrapf(err, "failed to parse JSON file")
	}

	return items, nil
}

// WriteBatch creates a new JSON file for storing ocsf data.
// It marshals the data into a JSON object and writes it to the specified file path.
func (s *localJsonDatastore[T]) WriteBatch(ctx context.Context, items []T) error {
	allItems := items

	if s.currentPath == "" {
		s.currentPath = filepath.Join(s.basepath, fmt.Sprintf("%s.json", time.Now().Format("20060102T150405Z")))
	} else {
		fileItems, err := s.GetItemsFromFile(ctx, s.currentPath)
		if err != nil {
			return oops.Wrapf(err, "failed to get existing items from disk")
		}

		allItems = append(allItems, fileItems...)
	}

	jsonData, err := json.Marshal(allItems)
	if err != nil {
		return oops.Wrapf(err, "failed to marshal items to JSON")
	}

	if err := os.WriteFile(s.currentPath, jsonData, 0644); err != nil {
		return oops.Wrapf(err, "failed to write JSON to disk")
	}

	slog.Info("Wrote JSON file to disk", "path", s.currentPath, "items", len(allItems))

	return nil
}
