package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Santiago-Labs/go-ocsf/clients/snyk"
	"github.com/Santiago-Labs/go-ocsf/clients/tenable"
	"github.com/Santiago-Labs/go-ocsf/datastore"
	"github.com/Santiago-Labs/go-ocsf/ocsf"
	"github.com/Santiago-Labs/go-ocsf/syncers"
	"github.com/Santiago-Labs/go-ocsf/syncers/gcpauditlog"
	"github.com/samsarahq/go/oops"

	"github.com/Santiago-Labs/go-ocsf/clients/clickhouse"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
)

func main() {
	isParquet := flag.Bool("parquet", false, "Use parquet format")
	isJSON := flag.Bool("json", false, "Use JSON format")
	isClickhouse := flag.Bool("clickhouse", false, "Use Clickhouse format")
	bucketName := flag.String("bucket-name", "", "S3 bucket name")
	tableBucketName := flag.String("table-bucket-name", "", "Table bucket name")
	// Sync data.
	syncSnykOption := flag.Bool("sync-snyk", false, "Sync Snyk data.")
	syncTenableOption := flag.Bool("sync-tenable", false, "Sync Tenable data.")
	syncSecurityHubOption := flag.Bool("sync-security-hub", false, "Sync SecurityHub data.")
	syncInspectorOption := flag.Bool("sync-inspector", false, "Sync Inspector data.")
	syncGCPAuditLogOption := flag.Bool("sync-gcp-audit-log", false, "Sync GCP AuditLog data.")
	shouldSetupClickhouse := flag.Bool("setup-clickhouse", false, "Setup Clickhouse DB")

	flag.Parse()

	fmt.Println("Starting...", *shouldSetupClickhouse)
	ctx := context.Background()
	var clickhouseClient *clickhouse.Client
	var err error

	if *isClickhouse || *shouldSetupClickhouse {
		clickhouseClient, err = clickhouse.New(ctx, clickhouse.Options{})
		if err != nil {
			log.Fatalf("Failed to create Clickhouse client: %v", err)
		}
		defer clickhouseClient.Close()
	}

	if *shouldSetupClickhouse {
		if err := setupClickhouse(ctx, clickhouseClient); err != nil {
			log.Fatalf("Failed to setup Clickhouse tables: %v", err)
		}
	}

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("Failed to load AWS config: %v", err)
	}

	storage, err := setupStorage(ctx, *isParquet, *isJSON, *bucketName, *tableBucketName, clickhouseClient)
	// snykAPIKey := os.Getenv("SNYK_API_KEY")
	// snykOrganizationID := os.Getenv("SNYK_ORGANIZATION_ID")

	// tenableAPIKey := os.Getenv("TENABLE_API_KEY")
	// tenableSecretKey := os.Getenv("TENABLE_SECRET_KEY")

	// storage, _, err := setupStorage(ctx, *isParquet, *isJSON, *bucketName, clickhouseClient)
	if err != nil {
		log.Fatalf("Failed to setup storage: %v", err)
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
	// if snykAPIKey != "" && snykOrganizationID != "" {
	// 	if err := syncSnyk(ctx, snykAPIKey, snykOrganizationID, storage); err != nil {
	// 		log.Fatalf("Failed to sync Snyk data: %v", err)
	// 	}
	// }

	// if tenableAPIKey != "" && tenableSecretKey != "" {
	// 	if err := syncTenable(ctx, tenableAPIKey, tenableSecretKey, storage); err != nil {
	// 		log.Fatalf("Failed to sync Tenable data: %v", err)
	// 	}
	// }

	// cfg, err := config.LoadDefaultConfig(ctx)
	// if err != nil {
	// 	log.Printf("Warning: Failed to load AWS config: %v. AWS services will be skipped.", err)
	// } else {
	// 	if err := inspectorSync(ctx, storage, cfg); err != nil {
	// 		log.Fatalf("Failed to sync Inspector data: %v", err)
	// 	}

	// 	if err := syncSecurityHub(ctx, storage, cfg); err != nil {
	// 		log.Fatalf("Failed to sync SecurityHub data: %v", err)
	// 	}
	// }

	if *syncGCPAuditLogOption {
		if err := syncGCPAuditLog(ctx, storage); err != nil {
			log.Fatalf("Failed to sync GCPAuditLog data: %v", err)
		}
	}

	if *syncInspectorOption {
		if err := inspectorSync(ctx, storage, cfg); err != nil {
			log.Fatalf("Failed to sync Inspector data: %v", err)
		}
	}

	// if err := syncToClickhouse(ctx, clickhouseClient, storage); err != nil {
	// 	log.Fatalf("Failed to sync to Clickhouse: %v", err)
	// }
}

func setupStorage(ctx context.Context, isParquet, isJSON bool, bucketName, tableBucketName string, clickhouseClient *clieckhouse.Client) (datastore.Datastore, error) {
	var storage datastore.Datastore
	var s3Client *s3.Client
	var err error

	if isParquet {
		if tableBucketName != "" {

			cfg, err := config.LoadDefaultConfig(ctx)
			if err != nil {
				return nil, oops.Wrapf(err, "failed to load config")
			}
			s3Client := s3.NewFromConfig(cfg)

			storage, err = datastore.NewS3TablesDatastore(ctx, tableBucketName, s3Client)
			if err != nil {
				return nil, fmt.Errorf("failed to create S3 tables datastore: %v", err)
			}
		} else if bucketName != "" {
			cfg, err := config.LoadDefaultConfig(ctx)
			if err != nil {
				return nil, fmt.Errorf("error loading AWS config: %v", err)
			}

			s3Client = s3.NewFromConfig(cfg)
			storage, err = datastore.NewS3ParquetDatastore(ctx, bucketName, s3Client)
			if err != nil {
				return nil, fmt.Errorf("failed to create S3 parquet datastore: %v", err)
			}
		} else {
			storage, err = datastore.NewLocalParquetDatastore(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to create local parquet datastore: %v", err)
			}
		}
	} else if isJSON {
		if bucketName != "" {
			cfg, err := config.LoadDefaultConfig(ctx)
			if err != nil {
				return nil, fmt.Errorf("error loading AWS config: %v", err)
			}

			s3Client = s3.NewFromConfig(cfg)
			storage, err = datastore.NewS3JsonDatastore(ctx, bucketName, s3Client)
			if err != nil {
				return nil, fmt.Errorf("failed to create S3 json datastore: %v", err)
			}
		} else {
			storage, err = datastore.NewLocalJsonDatastore(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to create local json datastore: %v", err)
			}
		}
	} else if clickhouseClient != nil {
		storage = datastore.NewClickhouseStore(clickhouseClient)
	} else {
		return nil, fmt.Errorf("no storage format specified, use --parquet or --json")
	}

	return storage, nil
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

func setupClickhouse(ctx context.Context, clickhouseClient *clickhouse.Client) error {
	err := clickhouseClient.CreateTableFromStruct(ctx, "vulnerability_findings", "ActivityID", "", ocsf.VulnerabilityFinding{})
	if err != nil {
		return fmt.Errorf("failed to create vulnerability_findings table: %v", err)
	}

	err = clickhouseClient.CreateTableFromStruct(ctx, "api_activities", "CategoryUID", "", ocsf.APIActivity{})
	if err != nil {
		return fmt.Errorf("failed to create api_activities table: %v", err)
	}

	return nil
}

// func syncToClickhouse(ctx context.Context, clickhouseClient *clickhouse.Client, vulnerabilityFindings []ocsf.VulnerabilityFinding, apiActivities []ocsf.APIActivity) error {
// 	// err := clickhouseClient.InsertFindings(ctx, vulnerabilityFindings)
// 	// if err != nil {
// 	// 	return fmt.Errorf("failed to insert vulnerability finding: %v", err)
// 	// }

// 	var err error
// 	err = clickhouseClient.InsertAPIActivities(ctx, apiActivities)
// 	if err != nil {
// 		return fmt.Errorf("failed to insert api activities: %v", err)
// 	}

// 	return nil
// }
