package datastore

import (
	"context"
	"errors"
	"path/filepath"

	"github.com/Santiago-Labs/go-ocsf/ocsf"
)

var Basepath = "data"

var BasepathFindings = filepath.Join(Basepath, "vulnerability_finding")
var BasepathActivities = filepath.Join(Basepath, "api_activity")

var ErrNotFound = errors.New("not found")

const maxFileSize = 128 * 1024 * 1024 // 128 MB
const avgFindingSize = 5 * 1024       // 5 KB, rough estimate

// Datastore defines the interface for vulnerability finding storage.
// It provides methods to retrieve, save, and manage vulnerability findings.
// Each implementation of Datastore is responsible for reading and writing to a file, and building the in-memory index of finding IDs to file paths.
type Datastore interface {
	// GetFinding retrieves a vulnerability finding by its unique identifier.
	GetFinding(ctx context.Context, findingID string) (*ocsf.VulnerabilityFinding, error)

	// GetAPIActivity retrieves a API activity by its unique identifier.
	GetAPIActivity(ctx context.Context, activityID string) (*ocsf.APIActivity, error)

	// GetFindingsFromFile retrieves all vulnerability findings from a specific file path.
	GetFindingsFromFile(ctx context.Context, path string) ([]ocsf.VulnerabilityFinding, error)

	// GetAPIActivitiesFromFile retrieves all API activities from a specific file path.
	GetAPIActivitiesFromFile(ctx context.Context, path string) ([]ocsf.APIActivity, error)

	// SaveFindings saves a list of vulnerability findings to the datastore.
	SaveFindings(ctx context.Context, findings []ocsf.VulnerabilityFinding) error

	// SaveAPIActivities saves a list of API activities to the datastore.
	SaveAPIActivities(ctx context.Context, activities []ocsf.APIActivity) error

	// WriteBatch writes a batch of vulnerability findings to a specific file in a specific format.
	WriteBatch(ctx context.Context, findings []ocsf.VulnerabilityFinding, path *string) error

	// WriteAPIActivityBatch writes a batch of API activities to a specific file in a specific format.
	WriteAPIActivityBatch(ctx context.Context, activities []ocsf.APIActivity, path *string) error
}
