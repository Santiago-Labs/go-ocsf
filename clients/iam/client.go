package iam

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

type IAMClient struct {
	iamClient *iam.Client
}

func NewIAMClient(ctx context.Context) (*IAMClient, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load default config, %v", err)
	}
	iamClient := iam.NewFromConfig(cfg)
	return &IAMClient{iamClient: iamClient}, nil
}

func (c *IAMClient) CreateLakeFormationAccessRole(ctx context.Context) error {
	roleName := "LakeFormationAccessRole"

	trustPolicy := `{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Effect": "Allow",
				"Principal": {
					"Service": [
						"lakeformation.amazonaws.com",
						"glue.amazonaws.com"
					]
				},
				"Action": "sts:AssumeRole"
			}
		]
	}`

	_, err := c.iamClient.CreateRole(ctx, &iam.CreateRoleInput{
		RoleName:                 aws.String(roleName),
		AssumeRolePolicyDocument: aws.String(trustPolicy),
		Description:              aws.String("Role for Lake Formation to access Glue and S3 resources"),
	})
	if err != nil {
		var eae *types.EntityAlreadyExistsException
		if errors.As(err, &eae) {
			log.Println("Lake Formation access role already exists")
			return nil
		} else {
			return fmt.Errorf("failed to create role, %v", err)
		}
	}

	permissionPolicy := `{
		{
		"Version": "2012-10-17",
		"Statement": [{
			"Sid": "AllowBucketActions",
			"Effect": "Allow",
			"Action": [
				"s3tables:CreateTableBucket",
				"s3tables:PutTableBucketPolicy",
				"s3tables:GetTableBucketPolicy",
				"s3tables:ListTableBuckets",
				"s3tables:GetTableBucket",
				"s3tables:CreateTable",
				"s3tables:PutTableData",
				"s3tables:GetTableData",
				"s3tables:GetTableMetadataLocation",
				"s3tables:UpdateTableMetadataLocation",
				"s3tables:GetNamespace",
				"s3tables:CreateNamespace"
			],

			"Resource": "arn:aws:s3tables:*:*:bucket/*"
		}]
	}`

	_, err = c.iamClient.PutRolePolicy(ctx, &iam.PutRolePolicyInput{
		RoleName:       aws.String(roleName),
		PolicyName:     aws.String("LakeFormationAccessPolicy"),
		PolicyDocument: aws.String(permissionPolicy),
	})
	if err != nil {
		return fmt.Errorf("failed to put role policy, %v", err)
	}

	return nil
}
