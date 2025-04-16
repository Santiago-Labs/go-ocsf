//go:build duckdb_use_lib

package main

/*
#cgo LDFLAGS: -lduckdb -L${SRCDIR}/duckdb_lib -Wl,-rpath,${SRCDIR}/duckdb_lib
*/
import "C"

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Santiago-Labs/go-ocsf/clients/athena"
	"github.com/Santiago-Labs/go-ocsf/clients/glue"
	"github.com/Santiago-Labs/go-ocsf/clients/iam"
	"github.com/Santiago-Labs/go-ocsf/clients/snyk"
	"github.com/Santiago-Labs/go-ocsf/clients/tenable"
	"github.com/Santiago-Labs/go-ocsf/datastore"
	"github.com/Santiago-Labs/go-ocsf/ocsf"
	"github.com/Santiago-Labs/go-ocsf/syncers"
	"github.com/Santiago-Labs/go-ocsf/syncers/gcpauditlog"
	"github.com/apache/arrow/go/v15/arrow"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go-v2/service/s3tables"
	s3tablesTypes "github.com/aws/aws-sdk-go-v2/service/s3tables/types"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
)

func main() {
	isParquet := flag.Bool("parquet", false, "Use parquet format")
	isJSON := flag.Bool("json", false, "Use JSON format")
	bucketName := flag.String("bucket-name", "", "S3 bucket name")

	// Create tables.
	setupS3TablesOption := flag.Bool("setup-s3-tables", false, "Setup S3 tables. Only required once per bucket.")

	// Sync data.
	syncSnykOption := flag.Bool("sync-snyk", false, "Sync Snyk data.")
	syncTenableOption := flag.Bool("sync-tenable", false, "Sync Tenable data.")
	syncSecurityHubOption := flag.Bool("sync-security-hub", false, "Sync SecurityHub data.")
	syncInspectorOption := flag.Bool("sync-inspector", false, "Sync Inspector data.")
	syncGCPAuditLogOption := flag.Bool("sync-gcp-audit-log", false, "Sync GCP AuditLog data.")
	flag.Parse()

	ctx := context.Background()

	snykAPIKey := os.Getenv("SNYK_API_KEY")
	snykOrganizationID := os.Getenv("SNYK_ORGANIZATION_ID")

	tenableAPIKey := os.Getenv("TENABLE_API_KEY")
	tenableSecretKey := os.Getenv("TENABLE_SECRET_KEY")

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("Failed to load AWS config: %v", err)
	}

	if *setupS3TablesOption {
		if *bucketName == "" {
			log.Fatal("--bucket-name is required when --setup-s3-tables is set")
		}

		s3tablesClient := s3tables.NewFromConfig(cfg)
		athenaClient := athena.NewClient(cfg)
		if err := setupS3Tables(ctx, *bucketName, s3tablesClient, athenaClient); err != nil {
			log.Fatalf("Failed to setup S3 tables: %v", err)
		}
	}

	storage, err := setupStorage(ctx, *isParquet, *isJSON, *bucketName)
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
}

func setupStorage(ctx context.Context, isParquet, isJSON bool, bucketName string) (datastore.Datastore, error) {
	var storage datastore.Datastore
	var s3Client *s3.Client
	var err error

	if isParquet {
		if bucketName != "" {
			cfg, err := config.LoadDefaultConfig(ctx)
			if err != nil {
				return nil, fmt.Errorf("error loading AWS config: %v", err)
			}

			s3Client = s3.NewFromConfig(cfg)
			storage, err = datastore.NewS3ParquetDatastore(bucketName, s3Client)
			if err != nil {
				return nil, fmt.Errorf("failed to create S3 parquet datastore: %v", err)
			}
		} else {
			storage, err = datastore.NewLocalParquetDatastore()
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
			storage, err = datastore.NewS3JsonDatastore(bucketName, s3Client)
			if err != nil {
				return nil, fmt.Errorf("failed to create S3 json datastore: %v", err)
			}
		} else {
			storage, err = datastore.NewLocalJsonDatastore()
			if err != nil {
				return nil, fmt.Errorf("failed to create local json datastore: %v", err)
			}
		}
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

func setupS3Tables(ctx context.Context, bucketName string, s3tablesClient *s3tables.Client, athenaClient *athena.Client) error {
	log.Println("Creating table bucket...")
	resp, err := s3tablesClient.CreateTableBucket(ctx, &s3tables.CreateTableBucketInput{
		Name: aws.String(bucketName),
	})

	var bne *types.BucketAlreadyExists
	var bucketArn string
	if err != nil {
		if !errors.As(err, &bne) {
			getResp, err := s3tablesClient.ListTableBuckets(ctx, &s3tables.ListTableBucketsInput{})
			if err != nil {
				return err
			}

			for _, bucket := range getResp.TableBuckets {
				if *bucket.Name == bucketName {
					log.Println("Bucket already exists:", bucket)
					bucketArn = *bucket.Arn
					break
				}
			}
		} else {
			return err
		}
	} else {
		bucketArn = *resp.Arn
	}

	log.Println("Creating namespace...")
	desiredNamespace := "ocsf_data"

	listNamespacesResp, err := s3tablesClient.ListNamespaces(ctx, &s3tables.ListNamespacesInput{
		TableBucketARN: aws.String(bucketArn),
	})
	if err != nil {
		return fmt.Errorf("failed to list namespaces: %v", err)
	}

	namespaceExists := false
	for _, namespace := range listNamespacesResp.Namespaces {
		if strings.Join(namespace.Namespace, ".") == desiredNamespace {
			log.Println("Namespace already exists:", namespace)
			namespaceExists = true
			break
		}
	}

	if !namespaceExists {
		_, err = s3tablesClient.CreateNamespace(ctx, &s3tables.CreateNamespaceInput{
			TableBucketARN: aws.String(bucketArn),
			Namespace:      []string{desiredNamespace},
		})
		if err != nil {
			return fmt.Errorf("failed to create namespace: %v", err)
		}
	}

	var nne *s3tablesTypes.ConflictException
	if err != nil && !errors.As(err, &nne) {
		return err
	}

	log.Println("Creating Lake Formation access role...")
	iamClient, err := iam.NewIAMClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create IAM client: %v", err)
	}

	err = iamClient.CreateLakeFormationAccessRole(ctx)
	if err != nil {
		return fmt.Errorf("failed to create Lake Formation access role: %v", err)
	}

	log.Println("Creating Glue resource link...")
	glueClient, err := glue.NewGlueClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create Glue client: %v", err)
	}
	err = glueClient.CreateDatabase(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("failed to create Glue resource link: %v", err)
	}

	log.Println("Creating Iceberg tables...")
	tables := map[string][]arrow.Field{
		"vulnerability_finding": ocsf.VulnerabilityFindingFields,
		"api_activity":          ocsf.APIActivityFields,
	}

	for tableName, fields := range tables {
		log.Println("Creating table:", tableName)
		s3location := fmt.Sprintf("s3://%s/data/%s", bucketName, tableName)
		err = athenaClient.CreateTable(ctx, fields, tableName, s3location, []string{})
		if err != nil {
			return fmt.Errorf("failed to create Athena table: %v", err)
		}

		log.Println("Table created successfully:", tableName)
	}

	return nil
}
