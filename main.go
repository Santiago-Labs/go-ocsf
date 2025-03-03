package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/Santiago-Labs/go-ocsf/clients/snyk"
	"github.com/Santiago-Labs/go-ocsf/datastore"
	"github.com/Santiago-Labs/go-ocsf/syncers"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	isParquet := flag.Bool("parquet", false, "Use parquet format")
	isJSON := flag.Bool("json", false, "Use JSON format")
	bucketName := flag.String("bucket-name", "", "S3 bucket name")

	flag.Parse()

	snykAPIKey := os.Getenv("SNYK_API_KEY")
	snykOrganizationID := os.Getenv("SNYK_ORGANIZATION_ID")

	ctx := context.Background()

	snykClient, err := snyk.NewClient(snykAPIKey, snykOrganizationID)
	if err != nil {
		log.Fatalf("Failed to create Snyk client: %v", err)
	}

	var s3Client *s3.Client
	if *bucketName != "" {
		cfg, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			log.Fatalf("Error loading AWS config: %v", err)
		}

		s3Client = s3.NewFromConfig(cfg)
	}

	var storage datastore.Datastore
	if *isParquet {
		if *bucketName != "" {
			storage = datastore.NewS3ParquetDatastore(*bucketName, s3Client)
		} else {
			storage, err = datastore.NewLocalParquetDatastore()
			if err != nil {
				log.Fatalf("Failed to create local parquet datastore: %v", err)
			}
		}
	} else if *isJSON {
		if *bucketName != "" {
			storage = datastore.NewS3JsonDatastore(*bucketName, s3Client)
		} else {
			storage, err = datastore.NewLocalJsonDatastore()
			if err != nil {
				log.Fatalf("Failed to create local json datastore: %v", err)
			}
		}
	}

	snykSyncer, err := syncers.NewSnykOCSFSyncer(ctx, snykClient, storage)
	if err != nil {
		log.Fatalf("Failed to create Snyk syncer: %v", err)
	}

	err = snykSyncer.Sync(ctx)
	if err != nil {
		log.Fatalf("Failed to sync Snyk data: %v", err)
	}
}
