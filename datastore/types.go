package datastore

import (
	"context"
	"errors"
	"path/filepath"
)

var Basepath = "data"

var (
	basepaths = map[string]string{
		"VulnerabilityFinding": filepath.Join(Basepath, "vulnerability_finding"),
		"APIActivity":          filepath.Join(Basepath, "api_activity"),
	}
)

var ErrNotFound = errors.New("not found")

const maxFileSize = 128 * 1024 * 1024 // 128 MB
const avgFindingSize = 5 * 1024       // 5 KB, rough estimate

// Datastore defines the interface for vulnerability finding storage.
// It provides methods to retrieve, save, and manage vulnerability findings.
// Each implementation of Datastore is responsible for reading and writing to a file, and building the in-memory index of finding IDs to file paths.
type Datastore[T any] interface {
	// Save saves a list to the datastore.
	Save(ctx context.Context, items []T) error

	// WriteBatch writes a batch of items to a specific file in a specific format.
	WriteBatch(ctx context.Context, items []T) error
}
