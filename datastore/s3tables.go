package datastore

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/Santiago-Labs/go-ocsf/ocsf"
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/iceberg-go"
	"github.com/apache/iceberg-go/catalog"
	"github.com/apache/iceberg-go/table"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3tables"
	"github.com/samsarahq/go/oops"

	_ "github.com/apache/iceberg-go/catalog/rest"
)

var (
	findingIdent  = table.Identifier([]string{"ocsf_data", "vulnerability_finding"})
	activityIdent = table.Identifier([]string{"ocsf_data", "api_activity"})
)

type s3TablesDatastore struct {
	s3Bucket           string
	apiActivitiesTable *table.Table
	findingsTable      *table.Table

	BaseDatastore
}

// NewS3TablesDatastore creates a new S3 Tables datastore.
func NewS3TablesDatastore(ctx context.Context, bucketName string, s3Client *s3tables.Client) (Datastore, error) {
	var bucketRegion string
	var bucketArn string
	var nextToken *string
	for {
		allbuckets, err := s3Client.ListTableBuckets(ctx, &s3tables.ListTableBucketsInput{
			ContinuationToken: nextToken,
		})
		if err != nil {
			return nil, oops.Wrapf(err, "failed to list buckets")
		}
		for _, bucket := range allbuckets.TableBuckets {
			if *bucket.Name == bucketName {
				bucketRegion = strings.Split(*bucket.Arn, ":")[3]
				bucketArn = *bucket.Arn
				break
			}
		}
		if allbuckets.ContinuationToken == nil {
			break
		}
		nextToken = allbuckets.ContinuationToken
	}

	props := iceberg.Properties{
		"type":                "rest",
		"warehouse":           bucketArn,
		"uri":                 fmt.Sprintf("https://s3tables.%s.amazonaws.com/iceberg", bucketRegion),
		"rest.sigv4-enabled":  "true",
		"rest.signing-name":   "s3tables",
		"rest.signing-region": bucketRegion,
	}

	cat, err := catalog.Load(ctx, "rest", props)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to create catalog")
	}

	// err = setup(ctx, s3Client, cat, bucketArn)
	// if err != nil {
	// 	return nil, oops.Wrapf(err, "failed to setup tables")
	// }

	findingsTable, err := cat.LoadTable(ctx, findingIdent, props)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to load table")
	}

	activitiesTable, err := cat.LoadTable(ctx, activityIdent, props)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to load table")
	}

	s := &s3TablesDatastore{
		s3Bucket:           bucketName,
		apiActivitiesTable: activitiesTable,
		findingsTable:      findingsTable,
	}

	s.BaseDatastore = BaseDatastore{
		store: s,
	}

	return s, nil
}

func setup(ctx context.Context, s3TablesClient *s3tables.Client, cat catalog.Catalog, bucketArn string) error {
	_, err := s3TablesClient.CreateNamespace(ctx, &s3tables.CreateNamespaceInput{
		Namespace:      []string{"ocsf_data"},
		TableBucketARN: aws.String(bucketArn),
	})
	if err != nil {
		return oops.Wrapf(err, "failed to create namespace")
	}

	err = createIcebergTable(ctx, cat, ocsf.VulnerabilityFindingSchema, findingIdent)
	if err != nil {
		return oops.Wrapf(err, "failed to create findings table")
	}

	err = createIcebergTable(ctx, cat, ocsf.APIActivitySchema, activityIdent)
	if err != nil {
		return oops.Wrapf(err, "failed to create api activity table")
	}

	return nil
}

func createIcebergTable(ctx context.Context, cat catalog.Catalog, arrowSchema *arrow.Schema, tableName table.Identifier) error {
	iceSchema, err := ArrowSchemaToIceberg(arrowSchema)
	if err != nil {
		return oops.Wrapf(err, "failed to create iceberg schema")
	}

	options := []catalog.CreateTableOpt{
		catalog.WithProperties(map[string]string{
			"type": "iceberg",
		}),
	}

	_, err = cat.CreateTable(ctx, tableName, iceSchema, options...)
	if err != nil {
		return oops.Wrapf(err, "failed to create table")
	}

	return nil
}

// Note: iceberg-go does not support writing to tables with partition specs yet.
func buildPartitionSpec(s *iceberg.Schema) (*iceberg.PartitionSpec, error) {
	col, ok := s.FindFieldByName("event_day")
	if !ok {
		return nil, fmt.Errorf(`field "event_day" not found in schema`)
	}

	const specFieldID = 1000

	spec := iceberg.NewPartitionSpec(
		iceberg.PartitionField{
			SourceID:  col.ID,
			FieldID:   specFieldID,
			Name:      "event_day",
			Transform: iceberg.IdentityTransform{},
		},
	)
	return &spec, nil
}

// WriteBatch creates a new Parquet file for storing vulnerability findings.
// It writes the findings to the specified file path.
func (s *s3TablesDatastore) WriteBatch(ctx context.Context, findings []ocsf.VulnerabilityFinding) error {

	recFixed, err := SliceToRecordBatch(findings, ocsf.VulnerabilityFindingSchema)
	if err != nil {
		return oops.Wrapf(err, "failed to create record batch")
	}
	defer recFixed.Release()

	annot, err := attachFieldIDs(recFixed.Schema(), s.findingsTable.Schema())
	if err != nil {
		return oops.Wrapf(err, "failed to attach field IDs")
	}

	cols := recFixed.Columns()
	for _, col := range cols {
		col.Retain()
	}
	newRec := array.NewRecord(annot, cols, recFixed.NumRows())
	defer newRec.Release()

	columns := make([]arrow.Column, newRec.NumCols())
	for i := 0; i < int(newRec.NumCols()); i++ {
		arr := newRec.Column(i)
		arr.Retain()

		chunked := arrow.NewChunked(arr.DataType(), []arrow.Array{arr})

		columns[i] = *arrow.NewColumn(
			annot.Field(i),
			chunked,
		)
	}

	tbl := array.NewTable(annot, columns, newRec.NumRows())
	defer tbl.Release()

	txn := s.findingsTable.NewTransaction()
	if err := txn.AppendTable(ctx, tbl, 1024, s.findingsTable.Properties()); err != nil {
		return oops.Wrapf(err, "failed to append table")
	}
	updated, err := txn.Commit(ctx)
	if err != nil {
		return oops.Wrapf(err, "failed to commit")
	}
	s.findingsTable = updated

	slog.Info("inserted findings",
		"bucket", s.s3Bucket,
		"rows", len(findings),
	)
	return nil
}

func (s *s3TablesDatastore) WriteAPIActivityBatch(ctx context.Context, activities []ocsf.APIActivity) error {
	recFixed, err := SliceToRecordBatch(activities, ocsf.APIActivitySchema)
	if err != nil {
		return oops.Wrapf(err, "failed to create record batch")
	}
	defer recFixed.Release()

	annot, err := attachFieldIDs(recFixed.Schema(), s.apiActivitiesTable.Schema())
	if err != nil {
		return oops.Wrapf(err, "failed to attach field IDs")
	}

	cols := recFixed.Columns()
	for _, col := range cols {
		col.Retain()
	}
	newRec := array.NewRecord(annot, cols, recFixed.NumRows())
	defer newRec.Release()

	columns := make([]arrow.Column, newRec.NumCols())
	for i := 0; i < int(newRec.NumCols()); i++ {
		arr := newRec.Column(i)
		arr.Retain()

		chunked := arrow.NewChunked(arr.DataType(), []arrow.Array{arr})

		columns[i] = *arrow.NewColumn(
			annot.Field(i),
			chunked,
		)
	}

	tbl := array.NewTable(annot, columns, newRec.NumRows())
	defer tbl.Release()

	txn := s.apiActivitiesTable.NewTransaction()
	if err := txn.AppendTable(ctx, tbl, 1024, s.apiActivitiesTable.Properties()); err != nil {
		return oops.Wrapf(err, "failed to append table")
	}
	updated, err := txn.Commit(ctx)
	if err != nil {
		return oops.Wrapf(err, "failed to commit")
	}
	s.apiActivitiesTable = updated

	slog.Info("Inserted activities using Athena",
		"bucket", s.s3Bucket,
		"activities", len(activities),
	)
	return nil
}
