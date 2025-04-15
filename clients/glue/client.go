package glue

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	gluetypes "github.com/aws/aws-sdk-go-v2/service/glue/types"
)

var (
	Namespace = "ocsf_data"
	Database  = "ocsf_data_database"
)

type GlueClient struct {
	glueClient *glue.Client
	accountID  string
	region     string
}

func NewGlueClient(ctx context.Context) (*GlueClient, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	var creds aws.Credentials
	creds, err = cfg.Credentials.Retrieve(context.TODO())
	if err != nil {
		return nil, err
	}

	return &GlueClient{
		glueClient: glue.NewFromConfig(cfg),
		accountID:  creds.AccountID,
		region:     cfg.Region,
	}, nil
}

func (c *GlueClient) CreateDatabase(ctx context.Context, tableBucketName string) error {
	input := &glue.CreateDatabaseInput{
		DatabaseInput: &gluetypes.DatabaseInput{
			Name: aws.String(Database),
			TargetDatabase: &gluetypes.DatabaseIdentifier{
				CatalogId:    aws.String(c.accountID + ":s3tablescatalog/" + tableBucketName),
				DatabaseName: aws.String(Namespace),
				Region:       aws.String(c.region),
			},
		},
	}

	_, err := c.glueClient.CreateDatabase(ctx, input)
	if err != nil {
		var ae *gluetypes.AlreadyExistsException
		if errors.As(err, &ae) {
			log.Printf("Glue database (%s) already exists.", Database)
			return nil
		}
		return err
	}

	log.Println("Glue database created successfully.")
	return nil
}
