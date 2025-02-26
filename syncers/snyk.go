package syncers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/Santiago-Labs/go-ocsf/clients/snyk"
	"github.com/Santiago-Labs/go-ocsf/ocsf"
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
	"github.com/apache/arrow/go/v15/arrow/memory"
	"github.com/apache/arrow/go/v15/parquet"
	"github.com/apache/arrow/go/v15/parquet/compress"
	"github.com/apache/arrow/go/v15/parquet/pqarrow"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/samsarahq/go/oops"
)

type DataSync interface {
	Sync(ctx context.Context) error
}

// ----------------------------------------------------------------------------
// SnykOCSFSyncer
// ----------------------------------------------------------------------------

type SnykOCSFSyncer struct {
	snykClient *snyk.Client
	s3Client   *s3.Client
	s3Bucket   string
	parquet    bool
	json       bool
}

func NewSnykOCSFSyncer(ctx context.Context, snykClient *snyk.Client, s3Client *s3.Client, s3Bucket string, parquet, json bool) DataSync {
	return &SnykOCSFSyncer{
		snykClient: snykClient,
		s3Client:   s3Client,
		s3Bucket:   s3Bucket,
		parquet:    parquet,
		json:       json,
	}
}

func (s *SnykOCSFSyncer) Sync(ctx context.Context) error {
	slog.Info("syncing Snyk data")

	issues, err := s.snykClient.ListIssues(ctx)
	if err != nil {
		return oops.Wrapf(err, "failed to list all issues")
	}

	slog.Info("found Snyk issues", "num_issues", len(issues))

	var findings []ocsf.VulnerabilityFinding
	for _, issue := range issues {
		project, err := s.snykClient.GetProject(ctx, issue.Relationships.ScanItem.Data.ID)
		if err != nil {
			return oops.Wrapf(err, "failed to fetch project for Snyk issue")
		}

		finding, err := s.ToOCSF(issue, project)
		if err != nil {
			return oops.Wrapf(err, "failed to build OCSF finding")
		}
		findings = append(findings, finding)
	}

	if s.parquet {
		err = s.writeVulnerabilityFindingToParquet(ctx, findings)
		if err != nil {
			return oops.Wrapf(err, "failed to write findings to Parquet")
		}
	}

	if s.json {
		err = s.writeVulnerabilityFindingToJSON(ctx, findings)
		if err != nil {
			return oops.Wrapf(err, "failed to write findings to JSON")
		}
	}

	slog.Info("Finished Snyk sync")
	return nil
}

func (s *SnykOCSFSyncer) ToOCSF(issue snyk.Issue, project *snyk.Project) (ocsf.VulnerabilityFinding, error) {
	severity, severityID := mapSnykSeverity(issue.Attributes.EffectiveSeverityLevel)
	status, statusID := mapSnykStatus(issue.Attributes.Status)
	createdAt := issue.Attributes.CreatedAt
	findingID := issue.Relationships.ScanItem.Data.ID
	projectName := project.Attributes.Name

	vendorName := "Snyk"

	var vulnerabilities []ocsf.VulnerabilityDetails
	exploitAvailable := issue.Attributes.ExploitDetails != nil
	if len(issue.Attributes.Problems) == 0 {
		vulnerabilities = append(vulnerabilities, ocsf.VulnerabilityDetails{
			UID:                &issue.ID,
			Desc:               &issue.Attributes.Description,
			Title:              &issue.Attributes.Title,
			Severity:           &severity,
			IsExploitAvailable: &exploitAvailable,
			FirstSeenTime:      &createdAt,
			IsFixAvailable:     false,
			LastSeenTime:       &createdAt,
			VendorName:         &vendorName,
			AffectedCode:       snykAffectedCode(issue, project),
			AffectedPackages:   snykAffectedPackages(issue),
		})
	} else {
		for _, problem := range issue.Attributes.Problems {
			vulnerabilities = append(vulnerabilities, ocsf.VulnerabilityDetails{
				UID:                &issue.ID,
				CVE:                snykProblemToCVE(problem),
				AffectedCode:       snykAffectedCode(issue, project),
				AffectedPackages:   snykAffectedPackages(issue),
				Desc:               &issue.Attributes.Description,
				Title:              &issue.Attributes.Title,
				Severity:           &severity,
				IsExploitAvailable: &exploitAvailable,
				FirstSeenTime:      &createdAt,
				IsFixAvailable:     false,
				LastSeenTime:       &createdAt,
				VendorName:         &vendorName,
			})
		}
	}

	resourceType := "Code"
	resource := ocsf.ResourceDetails{
		UID:  &findingID,
		Name: &projectName,
		Type: &resourceType,
	}

	activityID := int16(1)
	activityName := "Create"
	className := "Vulnerability Finding"
	categoryUID := int32(2)
	categoryName := "Findings"
	classUID := int32(4002)
	typeUID := int64(classUID)*100 + int64(activityID)
	typeName := "Vulnerability Finding: Create"
	productName := "Snyk"

	metadata := ocsf.Metadata{
		Product: ocsf.Product{
			Name:       &productName,
			VendorName: productName,
		},
		Version: "1.1.0",
	}

	now := time.Now()

	findingInfo := ocsf.FindingInfo{
		UID:           issue.ID,
		Title:         issue.Attributes.Title,
		Desc:          &issue.Attributes.Description,
		CreatedTime:   &createdAt,
		FirstSeenTime: &createdAt,
		LastSeenTime:  &now,
		DataSources:   []string{"snyk"},
		Types:         []string{"Vulnerability"},
	}

	finding := ocsf.VulnerabilityFinding{
		Time:            createdAt,
		StartTime:       &createdAt,
		ActivityID:      activityID,
		ActivityName:    &activityName,
		CategoryUID:     categoryUID,
		CategoryName:    &categoryName,
		ClassUID:        classUID,
		ClassName:       &className,
		Message:         &issue.Attributes.Description,
		Metadata:        metadata,
		Resource:        &resource,
		Status:          &status,
		StatusID:        &statusID,
		TypeUID:         typeUID,
		TypeName:        &typeName,
		Vulnerabilities: vulnerabilities,
		FindingInfo:     findingInfo,
		SeverityID:      int32(severityID),
	}

	return finding, nil
}

func (s *SnykOCSFSyncer) writeVulnerabilityFindingToParquet(ctx context.Context, findings []ocsf.VulnerabilityFinding) error {
	pool := memory.NewGoAllocator()
	builder := array.NewRecordBuilder(pool, ocsf.VulnerabilityFindingSchema)
	defer builder.Release()

	for _, f := range findings {
		f.WriteToParquet(builder)
	}

	rec := builder.NewRecord()
	defer rec.Release()

	table := array.NewTableFromRecords(rec.Schema(), []arrow.Record{rec})
	defer table.Release()

	buf := new(bytes.Buffer)
	props := parquet.NewWriterProperties(parquet.WithCompression(compress.Codecs.Snappy))
	arrowProps := pqarrow.NewArrowWriterProperties()

	if err := pqarrow.WriteTable(table, buf, 1024*1024, props, arrowProps); err != nil {
		return oops.Wrapf(err, "failed to write Parquet")
	}

	if s.s3Bucket != "" {
		key := fmt.Sprintf("snyk/%s.parquet", time.Now().Format("20060102T150405Z"))
		_, err := s.s3Client.PutObject(ctx, &s3.PutObjectInput{
			Bucket: &s.s3Bucket,
			Key:    &key,
			Body:   bytes.NewReader(buf.Bytes()),
		})
		if err != nil {
			return oops.Wrapf(err, "failed to upload Parquet to S3")
		}

		slog.Info("Wrote Parquet file to S3",
			"bucket", s.s3Bucket,
			"key", key,
			"size", buf.Len(),
		)
	} else {
		f, err := os.Create(fmt.Sprintf("snyk-%s.parquet", time.Now().Format("20060102T150405Z")))
		if err != nil {
			return oops.Wrapf(err, "failed to create Parquet file")
		}
		defer f.Close()
		f.Write(buf.Bytes())

		slog.Info("Wrote Parquet file to disk",
			"file", f.Name(),
			"size", buf.Len(),
		)
	}
	return nil
}

func (s *SnykOCSFSyncer) writeVulnerabilityFindingToJSON(ctx context.Context, findings []ocsf.VulnerabilityFinding) error {
	outerSchema := map[string]interface{}{
		"vulnerability_findings": findings,
	}
	jsonData, err := json.Marshal(outerSchema)
	if err != nil {
		return oops.Wrapf(err, "failed to marshal findings to JSON")
	}

	if s.s3Bucket != "" {
		key := fmt.Sprintf("snyk/%s.json", time.Now().Format("20060102T150405Z"))
		_, err := s.s3Client.PutObject(ctx, &s3.PutObjectInput{
			Bucket: &s.s3Bucket,
			Key:    &key,
			Body:   bytes.NewReader(jsonData),
		})
		if err != nil {
			return oops.Wrapf(err, "failed to upload JSON to S3")
		}

		slog.Info("Wrote JSON file to S3",
			"bucket", s.s3Bucket,
			"key", key,
		)
	} else {
		f, err := os.Create(fmt.Sprintf("snyk-%s.json", time.Now().Format("20060102T150405Z")))
		if err != nil {
			return oops.Wrapf(err, "failed to create JSON file")
		}
		defer f.Close()
		f.Write(jsonData)

		slog.Info("Wrote JSON file to disk",
			"file", f.Name(),
			"size", len(jsonData),
		)
	}

	return nil
}

// ----------------------------------------------------------------------------
// Helper Functions
// ----------------------------------------------------------------------------

func mapSnykSeverity(snykSeverity string) (string, int) {
	switch snykSeverity {
	case "info":
		return "Informational", 1
	case "low":
		return "Low", 2
	case "medium":
		return "Medium", 3
	case "high":
		return "High", 4
	case "critical":
		return "Critical", 5
	default:
		return "Unknown", 0
	}
}

func mapSnykStatus(snykStatus string) (string, int32) {
	switch snykStatus {
	case "resolved":
		return "Closed", 2
	default:
		return "Open", 1
	}
}

func snykProblemToCVE(problem snyk.Problem) *ocsf.CVE {
	if problem.Source == "CVE" {
		return &ocsf.CVE{
			UID: problem.ID,
			References: []string{
				*problem.URL,
			},
		}
	}
	return nil
}

func snykAffectedCode(issue snyk.Issue, project *snyk.Project) []ocsf.AffectedCode {
	var affectedCode []ocsf.AffectedCode
	for _, coordinate := range issue.Attributes.Coordinates {
		for _, representation := range coordinate.Representations {
			fileName := project.Attributes.TargetFile
			lineNumber := int32(0)
			endLine := int32(0)

			if representation.SourceLocation == nil {
				continue
			}

			if representation.SourceLocation.Region.Start.Line > 0 {
				lineNumber = int32(representation.SourceLocation.Region.Start.Line)
			}
			if representation.SourceLocation.Region.End.Line > 0 {
				endLine = int32(representation.SourceLocation.Region.End.Line)
			}

			fileObj := ocsf.File{
				Path: fileName,
			}

			affectedCode = append(affectedCode, ocsf.AffectedCode{
				File:      fileObj,
				StartLine: lineNumber,
				EndLine:   endLine,
			})
		}
	}

	return affectedCode
}

func snykAffectedPackages(issue snyk.Issue) []ocsf.AffectedSoftwarePackage {
	var affectedPackage []ocsf.AffectedSoftwarePackage
	for _, coordinate := range issue.Attributes.Coordinates {
		for _, representation := range coordinate.Representations {
			if representation.Dependency == nil {
				continue
			}

			affectedPackage = append(affectedPackage, ocsf.AffectedSoftwarePackage{
				Name:    representation.Dependency.PackageName,
				Version: representation.Dependency.PackageVersion,
			})
		}
	}
	return affectedPackage
}
