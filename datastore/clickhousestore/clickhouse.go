package clickhousestore

import (
	"context"
	"errors"
	"fmt"

	"github.com/Santiago-Labs/go-ocsf/clients/clickhouse"
	"github.com/Santiago-Labs/go-ocsf/ocsf"
)

// So we need to implement

type ClickhouseStore struct {
	clickhouseClient *clickhouse.Client
}

func NewClickhouseStore(clickhouseClient *clickhouse.Client) *ClickhouseStore {
	return &ClickhouseStore{
		clickhouseClient: clickhouseClient,
	}
}

// GetFinding retrieves a vulnerability finding by its unique identifier.
func (s *ClickhouseStore) GetFinding(ctx context.Context, findingID string) (*ocsf.VulnerabilityFinding, error) {
	query := fmt.Sprintf("SELECT * FROM %s.vulnerability_findings WHERE uid = '%s'", s.clickhouseClient.DBName(), findingID)
	rows, err := s.clickhouseClient.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var finding ocsf.VulnerabilityFinding
		err = rows.Scan(&finding)
		if err != nil {
			return nil, err
		}
		return &finding, nil
	}

	return nil, nil
}

// GetAPIActivity retrieves a API activity by its unique identifier.
func (s *ClickhouseStore) GetAPIActivity(ctx context.Context, activityID string) (*ocsf.APIActivity, error) {
	// query should look at api_activity metadata.correlation_uid
	query := fmt.Sprintf("SELECT * FROM %s.api_activities WHERE Metadata.CorrelationUID = '%s'", s.clickhouseClient.DBName(), activityID)
	rows, err := s.clickhouseClient.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var activity ocsf.APIActivity
		err = rows.Scan(&activity)
		if err != nil {
			return nil, err
		}
		return &activity, nil
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return nil, errors.New("not found")
}

// GetFindingsFromFile retrieves all vulnerability findings from a specific file path.
func (s *ClickhouseStore) GetFindingsFromFile(ctx context.Context, path string) ([]ocsf.VulnerabilityFinding, error) {
	return nil, nil
}

// GetAPIActivitiesFromFile retrieves all API activities from a specific file path.
func (s *ClickhouseStore) GetAPIActivitiesFromFile(ctx context.Context, path string) ([]ocsf.APIActivity, error) {
	return nil, nil
}

// SaveFindings saves a list of vulnerability findings to the datastore.
func (s *ClickhouseStore) SaveFindings(ctx context.Context, findings []ocsf.VulnerabilityFinding) error {
	return s.clickhouseClient.InsertFindings(ctx, findings)
}

// SaveAPIActivities saves a list of API activities to the datastore.
func (s *ClickhouseStore) SaveAPIActivities(ctx context.Context, activities []ocsf.APIActivity) error {
	return nil
}

// WriteBatch writes a batch of vulnerability findings to a specific file in a specific format.
func (s *ClickhouseStore) WriteBatch(ctx context.Context, findings []ocsf.VulnerabilityFinding, path *string) error {
	return nil
}

// WriteAPIActivityBatch writes a batch of API activities to a specific file in a specific format.
func (s *ClickhouseStore) WriteAPIActivityBatch(ctx context.Context, activities []ocsf.APIActivity, path *string) error {
	return nil
}
