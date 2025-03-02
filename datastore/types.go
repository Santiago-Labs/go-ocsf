package datastore

import (
	"context"
	"errors"

	"github.com/Santiago-Labs/go-ocsf/ocsf"
)

var basepath = "data/findings"
var ErrNotFound = errors.New("not found")

type Datastore interface {
	GetFinding(ctx context.Context, findingID string) (*ocsf.VulnerabilityFinding, error)
	GetAllFindings(ctx context.Context, path string) ([]ocsf.VulnerabilityFinding, error)
	SaveFindings(ctx context.Context, findings []ocsf.VulnerabilityFinding) error
}
