package athena

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
)

// Client represents an Athena API client
type Client struct {
	athenaClient *athena.Client
	database     string
	workgroup    string
}

// NewClient creates a new Athena client
func NewClient(athenaClient *athena.Client, database, workgroup string) *Client {
	return &Client{
		athenaClient: athenaClient,
		database:     database,
		workgroup:    workgroup,
	}
}

func (c *Client) CreateTable(ctx context.Context, fields []arrow.Field, tableName, s3Location string, partitions []string) error {
	createStmt := GenerateAthenaTable(fields, tableName, s3Location, partitions)

	startQueryInput := &athena.StartQueryExecutionInput{
		QueryString: aws.String(createStmt),
		QueryExecutionContext: &types.QueryExecutionContext{
			Database: aws.String(c.database),
		},
		WorkGroup: aws.String(c.workgroup),
		ResultConfiguration: &types.ResultConfiguration{
			OutputLocation: aws.String(s3Location),
		},
	}

	result, err := c.athenaClient.StartQueryExecution(ctx, startQueryInput)
	if err != nil {
		return fmt.Errorf("failed to start query execution: %w", err)
	}

	queryExecutionID := result.QueryExecutionId
	for {
		status, err := c.athenaClient.GetQueryExecution(ctx, &athena.GetQueryExecutionInput{
			QueryExecutionId: queryExecutionID,
		})
		if err != nil {
			return fmt.Errorf("failed to get query execution status: %w", err)
		}

		state := status.QueryExecution.Status.State
		if state == types.QueryExecutionStateSucceeded {
			return nil
		}
		if state == types.QueryExecutionStateFailed || state == types.QueryExecutionStateCancelled {
			return fmt.Errorf("query execution failed: %s", *status.QueryExecution.Status.StateChangeReason)
		}
	}
}
