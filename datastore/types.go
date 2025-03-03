package datastore

import (
	"context"
	"errors"

	"github.com/Santiago-Labs/go-ocsf/ocsf"
)

var basepath = "data/findings"
var ErrNotFound = errors.New("not found")

const maxFileSize = 128 * 1024 * 1024 // 128 MB
const avgFindingSize = 5 * 1024       // 5 KB, rough estimate

// Datastore defines the interface for vulnerability finding storage.
// It provides methods to retrieve, save, and manage vulnerability findings.
// Each implementation of Datastore is responsible for reading and writing to a file, and building the in-memory index of finding IDs to file paths.
type Datastore interface {
	// GetFinding retrieves a vulnerability finding by its unique identifier.
	GetFinding(ctx context.Context, findingID string) (*ocsf.VulnerabilityFinding, error)

	// GetFindingsFromFile retrieves all vulnerability findings from a specific file path.
	GetFindingsFromFile(ctx context.Context, path string) ([]ocsf.VulnerabilityFinding, error)

	// SaveFindings saves a list of vulnerability findings to the datastore.
	SaveFindings(ctx context.Context, findings []ocsf.VulnerabilityFinding) error

	// WriteBatch writes a batch of vulnerability findings to a specific file in a specific format.
	WriteBatch(ctx context.Context, findings []ocsf.VulnerabilityFinding, path *string) error
}
