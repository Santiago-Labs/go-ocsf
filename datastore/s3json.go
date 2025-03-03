package datastore

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"path/filepath"
	"strings"
	"time"

	"github.com/Santiago-Labs/go-ocsf/ocsf"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/samsarahq/go/oops"
)

type s3JsonDatastore struct {
	s3Bucket string
	s3Client *s3.Client

	BaseDatastore
}

// NewS3JsonDatastore creates a new S3 JSON datastore.
// It initializes an in-memory index of finding IDs to file paths.
func NewS3JsonDatastore(bucketName string, s3Client *s3.Client) Datastore {
	s := &s3JsonDatastore{
		s3Bucket: bucketName,
		s3Client: s3Client,
	}
	s.BaseDatastore = BaseDatastore{
		findingIndex: make(map[string]string),
		fileIndex:    make(map[string]int),
		store:        s,
	}

	ctx := context.Background()
	if err := s.buildFindingIndex(ctx); err != nil {
		slog.Warn("failed to build complete finding index", "error", err)
	}

	return s
}

// GetFindingsFromFile retrieves all vulnerability findings from a specific file path.
// It reads the JSON file and parses it into a slice of vulnerability findings.
func (s *s3JsonDatastore) GetFindingsFromFile(ctx context.Context, key string) ([]ocsf.VulnerabilityFinding, error) {
	result, err := s.s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.s3Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, oops.Wrapf(err, "failed to get JSON file from S3")
	}
	defer result.Body.Close()

	data, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to read JSON file data")
	}

	var findings struct {
		VulnerabilityFindings []ocsf.VulnerabilityFinding `json:"vulnerability_findings"`
	}

	if err := json.Unmarshal(data, &findings); err != nil {
		return nil, oops.Wrapf(err, "failed to parse JSON file")
	}

	return findings.VulnerabilityFindings, nil
}

// WriteBatch creates a new JSON file for storing vulnerability findings.
// It marshals the findings into a JSON object and writes it to the specified file path.
func (s *s3JsonDatastore) WriteBatch(ctx context.Context, findings []ocsf.VulnerabilityFinding, key *string) error {
	allFindings := findings
	if key == nil {
		newkey := filepath.Join(basepath, fmt.Sprintf("%s.json", time.Now().Format("20060102T150405Z")))
		key = &newkey
	} else {
		var err error
		allFindings, err = s.GetFindingsFromFile(ctx, *key)
		if err != nil {
			return oops.Wrapf(err, "failed to get existing findings from disk")
		}

		allFindings = append(allFindings, findings...)
	}

	outerSchema := map[string]interface{}{
		"vulnerability_findings": allFindings,
	}
	jsonData, err := json.Marshal(outerSchema)
	if err != nil {
		return oops.Wrapf(err, "failed to marshal findings to JSON")
	}

	_, err = s.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &s.s3Bucket,
		Key:    key,
		Body:   bytes.NewReader(jsonData),
	})
	if err != nil {
		return oops.Wrapf(err, "failed to upload JSON to S3")
	}

	for _, f := range allFindings {
		s.BaseDatastore.findingIndex[f.FindingInfo.UID] = *key
	}
	s.BaseDatastore.fileIndex[*key] = len(allFindings)

	slog.Info("Wrote JSON file to S3",
		"bucket", s.s3Bucket,
		"key", *key,
		"findings", len(allFindings),
	)

	return nil
}

// buildFindingIndex builds the datastore's in-memory index of finding IDs to file paths.
// It reads all JSON files in the S3 bucket and parses them into a slice of vulnerability findings.
func (s *s3JsonDatastore) buildFindingIndex(ctx context.Context) error {
	paginator := s3.NewListObjectsV2Paginator(s.s3Client, &s3.ListObjectsV2Input{
		Bucket: aws.String(s.s3Bucket),
		Prefix: aws.String(basepath),
	})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return oops.Wrapf(err, "failed to list objects in S3")
		}

		for _, object := range output.Contents {
			if strings.HasSuffix(*object.Key, "/") || !strings.HasSuffix(*object.Key, ".json") {
				continue
			}

			if err := s.loadFileIntoIndex(ctx, *object.Key); err != nil {
				slog.Warn("error indexing json file", "key", *object.Key, "error", err)
			}
		}
	}

	slog.Info("built finding index from S3", "count", len(s.BaseDatastore.findingIndex))
	return nil
}
