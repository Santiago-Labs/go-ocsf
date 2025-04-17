package datastore

import (
	"context"
	"fmt"
	"log/slog"
	"sync"

	"github.com/Santiago-Labs/go-ocsf/ocsf"
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/apache/iceberg-go"
	"github.com/apache/iceberg-go/catalog"
	"github.com/apache/iceberg-go/table"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/samsarahq/go/oops"

	_ "github.com/apache/iceberg-go/catalog/glue"
)

var (
	onceFindingsSchema sync.Once
	findingsSchema     *arrow.Schema

	onceActivitiesSchema sync.Once
	activitiesSchema     *arrow.Schema
)

func initFindingsSchema() {
	onceFindingsSchema.Do(initFindingsSchema)
}

func initActivitiesSchema() {
	onceActivitiesSchema.Do(initActivitiesSchema)
}

type s3TablesDatastore struct {
	s3Bucket           string
	apiActivitiesTable *table.Table
	findingsTable      *table.Table

	BaseDatastore
}

// NewS3TablesDatastore creates a new S3 Tables datastore.
// It initializes an in-memory index of finding IDs to file paths.
func NewS3TablesDatastore(ctx context.Context, bucketName string) (Datastore, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to load config")
	}
	s3Client := s3.NewFromConfig(cfg)

	bucketRegion, err := s3Client.GetBucketLocation(ctx, &s3.GetBucketLocationInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return nil, oops.Wrapf(err, "failed to get bucket region")
	}

	props := iceberg.Properties{
		"type":         "glue",
		"glue.region":  string(bucketRegion.LocationConstraint),
		"glue.catalog": fmt.Sprintf("s3tablescatalog/%s", bucketName),
	}

	cat, err := catalog.Load(ctx, "glue", props)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to create catalog")
	}

	findingsIdent := table.Identifier([]string{"ocsf_data", "vulnerability_finding"})
	findingsTable, err := cat.LoadTable(ctx, findingsIdent, props)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to load table")
	}

	apiActivitiesIdent := table.Identifier([]string{"ocsf_data", "api_activities"})
	apiActivitiesTable, err := cat.LoadTable(ctx, apiActivitiesIdent, props)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to load table")
	}

	s := &s3TablesDatastore{
		s3Bucket:           bucketName,
		apiActivitiesTable: apiActivitiesTable,
		findingsTable:      findingsTable,
	}

	s.BaseDatastore = BaseDatastore{
		store: s,
	}

	return s, nil
}

// WriteBatch creates a new Parquet file for storing vulnerability findings.
// It writes the findings to the specified file path and updates the datastore's in-memory index.
func (s *s3TablesDatastore) WriteBatch(ctx context.Context, findings []ocsf.VulnerabilityFinding, keyPrefix string) error {
	onceFindingsSchema.Do(initFindingsSchema)

	mem := memory.NewGoAllocator()
	rb := array.NewRecordBuilder(mem, findingsSchema)
	defer rb.Release()

	for _, row := range findings {
		if err := buildRecord(rb, row); err != nil {
			return err
		}
	}
	rec := rb.NewRecord()
	defer rec.Release()

	tbl := array.NewTableFromRecords(findingsSchema, []arrow.Record{rec})
	defer tbl.Release()

	txn := s.findingsTable.NewTransaction()
	if err := txn.AppendTable(ctx, tbl, 1024, s.findingsTable.Properties()); err != nil {
		return err
	}
	newTbl, err := txn.Commit(ctx)
	if err != nil {
		return err
	}
	s.findingsTable = newTbl
	slog.Info("Inserted findings using Athena",
		"bucket", s.s3Bucket,
		"findings", len(findings),
	)
	return nil
}

func (s *s3TablesDatastore) WriteAPIActivityBatch(ctx context.Context, activities []ocsf.APIActivity, keyPrefix string) error {

	onceActivitiesSchema.Do(initActivitiesSchema)

	mem := memory.NewGoAllocator()
	rb := array.NewRecordBuilder(mem, activitiesSchema)
	defer rb.Release()

	for _, row := range activities {
		if err := buildRecord(rb, row); err != nil {
			return err
		}
	}
	rec := rb.NewRecord()
	defer rec.Release()

	tbl := array.NewTableFromRecords(activitiesSchema, []arrow.Record{rec})
	defer tbl.Release()

	txn := s.apiActivitiesTable.NewTransaction()
	if err := txn.AppendTable(ctx, tbl, 1024, s.apiActivitiesTable.Properties()); err != nil {
		return err
	}
	newTbl, err := txn.Commit(ctx)
	if err != nil {
		return err
	}
	s.apiActivitiesTable = newTbl

	slog.Info("Inserted activities using Athena",
		"bucket", s.s3Bucket,
		"activities", len(activities),
	)
	return nil
}
