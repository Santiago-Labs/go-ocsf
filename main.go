package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Santiago-Labs/go-ocsf/clients/athena"
	"github.com/Santiago-Labs/go-ocsf/clients/snyk"
	"github.com/Santiago-Labs/go-ocsf/clients/tenable"
	"github.com/Santiago-Labs/go-ocsf/datastore"
	"github.com/Santiago-Labs/go-ocsf/ocsf"
	"github.com/Santiago-Labs/go-ocsf/syncers"
	"github.com/Santiago-Labs/go-ocsf/syncers/gcpauditlog"
	"github.com/apache/arrow/go/v15/arrow"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	awsathena "github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
)

func main() {
	isParquet := flag.Bool("parquet", false, "Use parquet format")
	isJSON := flag.Bool("json", false, "Use JSON format")
	bucketName := flag.String("bucket-name", "", "S3 bucket name")

	// Create tables.
	createS3Table := flag.Bool("create-s3-table", false, "Create S3 table. Only required once per bucket.")
	createVulnFindingTable := flag.Bool("create-vuln-finding-table", false, "Create vulnerability finding table. Only required once per bucket.")
	createAPIActivityTable := flag.Bool("create-api-activity-table", false, "Create API activity table. Only required once per bucket.")

	// Sync data.
	syncSnykOption := flag.Bool("sync-snyk", false, "Sync Snyk data.")
	syncTenableOption := flag.Bool("sync-tenable", false, "Sync Tenable data.")
	syncSecurityHubOption := flag.Bool("sync-security-hub", false, "Sync SecurityHub data.")
	syncInspectorOption := flag.Bool("sync-inspector", false, "Sync Inspector data.")

	flag.Parse()

	ctx := context.Background()

	snykAPIKey := os.Getenv("SNYK_API_KEY")
	snykOrganizationID := os.Getenv("SNYK_ORGANIZATION_ID")

	tenableAPIKey := os.Getenv("TENABLE_API_KEY")
	tenableSecretKey := os.Getenv("TENABLE_SECRET_KEY")

	storage, _, err := setupStorage(ctx, *isParquet, *isJSON, *bucketName)
	if err != nil {
		log.Fatalf("Failed to setup storage: %v", err)
	}

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("Failed to load AWS config: %v", err)
	}

	if *createS3Table {
		if *bucketName == "" {
			log.Fatal("--bucket-name is required when --create-s3-table is set")
		}

		if err := createAthenaTables(ctx, cfg, *bucketName, *createVulnFindingTable, *createAPIActivityTable); err != nil {
			log.Fatalf("Failed to create Athena table: %v", err)
		}
	}

	if *syncSnykOption {
		if snykAPIKey == "" || snykOrganizationID == "" {
			log.Fatal("SNYK_API_KEY and SNYK_ORGANIZATION_ID must be set when --sync-snyk is set")
		}

		if err := syncSnyk(ctx, snykAPIKey, snykOrganizationID, storage); err != nil {
			log.Fatalf("Failed to sync Snyk data: %v", err)
		}
	}

	if *syncTenableOption {
		if tenableAPIKey == "" || tenableSecretKey == "" {
			log.Fatal("TENABLE_API_KEY and TENABLE_SECRET_KEY must be set when --sync-tenable is set")
		}

		if err := syncTenable(ctx, tenableAPIKey, tenableSecretKey, storage); err != nil {
			log.Fatalf("Failed to sync Tenable data: %v", err)
		}
	}

	if *syncSecurityHubOption {
		if err := syncSecurityHub(ctx, storage, cfg); err != nil {
			log.Fatalf("Failed to sync SecurityHub data: %v", err)
		}
	}

	if err := syncGCPAuditLog(ctx, storage); err != nil {
		log.Fatalf("Failed to sync GCPAuditLog data: %v", err)
	}

	if *syncInspectorOption {
		if err := inspectorSync(ctx, storage, cfg); err != nil {
			log.Fatalf("Failed to sync Inspector data: %v", err)
		}
	}
}

func setupStorage(ctx context.Context, isParquet, isJSON bool, bucketName string) (datastore.Datastore, *s3.Client, error) {
	var storage datastore.Datastore
	var s3Client *s3.Client
	var err error

	if bucketName != "" {
		cfg, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			return nil, nil, fmt.Errorf("error loading AWS config: %v", err)
		}

		s3Client = s3.NewFromConfig(cfg)
	}

	if isParquet {
		if bucketName != "" {
			storage = datastore.NewS3ParquetDatastore(bucketName, s3Client)
		} else {
			storage, err = datastore.NewLocalParquetDatastore()
			if err != nil {
				return nil, nil, fmt.Errorf("failed to create local parquet datastore: %v", err)
			}
		}
	} else if isJSON {
		if bucketName != "" {
			storage = datastore.NewS3JsonDatastore(bucketName, s3Client)
		} else {
			storage, err = datastore.NewLocalJsonDatastore()
			if err != nil {
				return nil, nil, fmt.Errorf("failed to create local json datastore: %v", err)
			}
		}
	} else {
		return nil, nil, fmt.Errorf("no storage format specified, use --parquet or --json")
	}

	return storage, s3Client, nil
}

func syncSnyk(ctx context.Context, apiKey, orgID string, storage datastore.Datastore) error {
	snykClient, err := snyk.NewClient(apiKey, orgID)
	if err != nil {
		return fmt.Errorf("failed to create Snyk client: %v", err)
	}

	snykSyncer, err := syncers.NewSnykOCSFSyncer(ctx, snykClient, storage)
	if err != nil {
		return fmt.Errorf("failed to create Snyk syncer: %v", err)
	}

	err = snykSyncer.Sync(ctx)
	if err != nil {
		log.Fatalf("Failed to sync Snyk data: %v", err)
	}

	return snykSyncer.Sync(ctx)
}

func inspectorSync(ctx context.Context, storage datastore.Datastore, cfg aws.Config) error {
	inspectorClient := inspector2.NewFromConfig(cfg)

	inspectorSyncer := syncers.NewInspectorOCSFSyncer(ctx, inspectorClient, storage)
	return inspectorSyncer.Sync(ctx)
}

func syncTenable(ctx context.Context, apiKey, secretKey string, storage datastore.Datastore) error {
	tenableClient, err := tenable.NewClient(apiKey, secretKey)
	if err != nil {
		return fmt.Errorf("failed to create Tenable client: %v", err)
	}

	tenableSyncer, err := syncers.NewTenableOCSFSyncer(ctx, tenableClient, storage)
	if err != nil {
		return fmt.Errorf("failed to create Tenable syncer: %v", err)
	}

	return tenableSyncer.Sync(ctx)
}

func syncSecurityHub(ctx context.Context, storage datastore.Datastore, cfg aws.Config) error {
	securityHubClient := securityhub.NewFromConfig(cfg)

	securityHubSyncer := syncers.NewSecurityHubOCSFSyncer(ctx, securityHubClient, storage)
	return securityHubSyncer.Sync(ctx)
}

func syncGCPAuditLog(ctx context.Context, storage datastore.Datastore) error {
	gcpauditlogSyncer, err := gcpauditlog.NewGCPAuditLogSyncer(ctx, storage, os.Getenv("GCP_PROJECT_ID"))
	if err != nil {
		return fmt.Errorf("failed to create GCPAuditLog syncer: %v", err)
	}

	return gcpauditlogSyncer.Sync(ctx)
}

func createAthenaTables(ctx context.Context, cfg aws.Config, bucketName string, createVulnFindingTable, createAPIActivityTable bool) error {
	athenaClient := awsathena.NewFromConfig(cfg)
	client := athena.NewClient(athenaClient, "default", "primary")

	tables := make(map[string][]arrow.Field)
	if createVulnFindingTable {
		tables["vulnerability_finding"] = ocsf.VulnerabilityFindingFields
	}
	if createAPIActivityTable {
		tables["api_activity"] = ocsf.APIActivityFields
	}

	for tableName, fields := range tables {
		s3Location := fmt.Sprintf("s3://%s/%s/%s", bucketName, datastore.Basepath, tableName)
		err := client.CreateTable(ctx, fields, tableName, s3Location, []string{}) // TODO: partition by source and day.
		if err != nil {
			return fmt.Errorf("failed to create Athena table: %w", err)
		}

		log.Printf("Successfully created Athena table '%s' in database 'default'", tableName)
	}
	return nil
}
