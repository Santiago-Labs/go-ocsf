package datastore

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Santiago-Labs/go-ocsf/ocsf"
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/iceberg-go"
	"github.com/apache/iceberg-go/catalog"
	"github.com/apache/iceberg-go/table"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/parquet-go/parquet-go"
	"github.com/samsarahq/go/oops"

	_ "github.com/apache/iceberg-go/catalog/rest"
)

type s3TablesDatastore struct {
	s3Bucket           string
	apiActivitiesTable *table.Table
	findingsTable      *table.Table

	findingsTableSchema   *parquet.Schema
	activitiesTableSchema *parquet.Schema

	BaseDatastore
}

// NewS3TablesDatastore creates a new S3 Tables datastore.
func NewS3TablesDatastore(ctx context.Context, bucketName string, s3Client *s3.Client) (Datastore, error) {
	bucketRegion, err := s3Client.GetBucketLocation(ctx, &s3.GetBucketLocationInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return nil, oops.Wrapf(err, "failed to get bucket region")
	}

	region := string(bucketRegion.LocationConstraint)
	if region == "" {
		region = "us-east-1"
	}

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	creds, err := cfg.Credentials.Retrieve(ctx)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to fetch credentials")
	}

	props := iceberg.Properties{
		"type":                "rest",
		"warehouse":           fmt.Sprintf("arn:aws:s3tables:%s:%s:bucket/%s", region, creds.AccountID, bucketName),
		"uri":                 fmt.Sprintf("https://s3tables.%s.amazonaws.com/iceberg", region),
		"rest.sigv4-enabled":  "true",
		"rest.signing-name":   "s3tables",
		"rest.signing-region": region,
	}

	cat, err := catalog.Load(ctx, "rest", props)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to create catalog")
	}

	findingIdent := table.Identifier([]string{"ocsf_data", "vulnerability_finding"})
	// err = createIcebergTable(ctx, cat, ocsf.VulnerabilityFindingSchema, findingIdent)
	// if err != nil {
	// 	return nil, oops.Wrapf(err, "failed to create table")
	// }

	findingsTable, err := cat.LoadTable(ctx, findingIdent, props)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to load table")
	}

	patchedfindingsSchema, err := patchParquetSchema(*findingsTable, parquet.SchemaOf(new(ocsf.VulnerabilityFinding)), "VulnerabilityFinding")
	if err != nil {
		return nil, oops.Wrapf(err, "failed to patch parquet schema")
	}

	activityIdent := table.Identifier([]string{"ocsf_data", "api_activity"})
	// err = createIcebergTable(ctx, cat, ocsf.APIActivitySchema, activityIdent)
	// if err != nil {
	// 	return nil, oops.Wrapf(err, "failed to create table")
	// }
	activitiesTable, err := cat.LoadTable(ctx, activityIdent, props)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to load table")
	}

	patchedActivitiesSchema, err := patchParquetSchema(*activitiesTable, parquet.SchemaOf(new(ocsf.APIActivity)), "APIActivity")
	if err != nil {
		return nil, oops.Wrapf(err, "failed to patch parquet schema")
	}

	s := &s3TablesDatastore{
		s3Bucket:              bucketName,
		apiActivitiesTable:    activitiesTable,
		findingsTable:         findingsTable,
		findingsTableSchema:   patchedfindingsSchema,
		activitiesTableSchema: patchedActivitiesSchema,
	}

	s.BaseDatastore = BaseDatastore{
		store: s,
	}

	return s, nil
}

func patchParquetSchema(icebergTable table.Table, parquetSchema *parquet.Schema, tableName string) (*parquet.Schema, error) {
	icebergSchema := buildParquetSchemaFromIceberg(icebergTable)

	patchedRoot, err := applyFieldIDs(icebergSchema, parquetSchema)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to apply field IDs")
	}

	patchedSchema := parquet.NewSchema(tableName, patchedRoot)

	return patchedSchema, nil
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
