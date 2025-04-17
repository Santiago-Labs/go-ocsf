package athena

import (
	"context"
	"fmt"
	"strings"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	DefaultDatabase  = "ocsf_data"
	DefaultWorkgroup = "primary"
)

type Client struct {
	athenaClient *athena.Client
	s3Client     *s3.Client
	database     string
	workgroup    string
}

func NewClient(cfg aws.Config) *Client {
	athenaClient := athena.NewFromConfig(cfg)
	s3Client := s3.NewFromConfig(cfg)

	return &Client{
		athenaClient: athenaClient,
		s3Client:     s3Client,
		database:     DefaultDatabase,
		workgroup:    DefaultWorkgroup,
	}
}

func (c *Client) CreateTable(ctx context.Context, fields []arrow.Field, tableName, s3Location string, partitions []string) error {
	createStmt := GenerateAthenaCreateTable(c.database, tableName, fields)

	startQueryInput := &athena.StartQueryExecutionInput{
		QueryString: aws.String(createStmt),
		QueryExecutionContext: &types.QueryExecutionContext{
			Database: aws.String(c.database),
			Catalog:  aws.String(fmt.Sprintf("s3tablescatalog/%s", getBucketName(s3Location))),
		},
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

func getBucketName(s3Location string) string {
	return strings.Split(s3Location, "/")[2]
}
