package syncers

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/Santiago-Labs/go-ocsf/clients/snyk"
	"github.com/Santiago-Labs/go-ocsf/datastore"
	"github.com/Santiago-Labs/go-ocsf/ocsf"
	"github.com/samsarahq/go/oops"
)

type DataSync interface {
	Sync(ctx context.Context) error
}

type SnykOCSFSyncer struct {
	snykClient *snyk.Client
	datastore  datastore.Datastore
	org        *snyk.Org
}

// NewSnykOCSFSyncer creates a new SnykOCSFSyncer
// It initializes the Snyk client and datastore, and fetches the organization details.
func NewSnykOCSFSyncer(ctx context.Context, snykClient *snyk.Client, datastore datastore.Datastore) (DataSync, error) {
	org, err := snykClient.GetOrg(ctx)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to fetch org")
	}

	return &SnykOCSFSyncer{
		snykClient: snykClient,
		datastore:  datastore,
		org:        org,
	}, nil
}

// Sync synchronizes Snyk data with the OCSF datastore
// It fetches all issues from Snyk, builds OCSF findings, and saves them to the datastore.
func (s *SnykOCSFSyncer) Sync(ctx context.Context) error {
	slog.Info("syncing Snyk data")

	issues, err := s.snykClient.ListIssues(ctx)
	if err != nil {
		return oops.Wrapf(err, "failed to list all issues")
	}

	slog.Info("found Snyk issues", "num_issues", len(issues))

	var findingIDs []string
	for _, issue := range issues {
		findingIDs = append(findingIDs, issue.ID)
	}

	var findingsToSave []ocsf.VulnerabilityFinding
	for _, issue := range issues {
		project, err := s.snykClient.GetProject(ctx, issue.Relationships.ScanItem.Data.ID)
		if err != nil {
			return oops.Wrapf(err, "failed to fetch project for Snyk issue")
		}

		finding, err := s.ToOCSF(ctx, issue, project)
		if err != nil {
			return oops.Wrapf(err, "failed to build OCSF finding")
		}

		findingsToSave = append(findingsToSave, finding)
	}

	err = s.datastore.SaveFindings(ctx, findingsToSave)
	if err != nil {
		return oops.Wrapf(err, "failed to save findings")
	}

	slog.Info("Finished Snyk sync")
	return nil
}

// ToOCSF converts a Snyk issue into an OCSF vulnerability finding.
func (s *SnykOCSFSyncer) ToOCSF(ctx context.Context, issue snyk.Issue, project *snyk.Project) (ocsf.VulnerabilityFinding, error) {
	severity, severityID := mapSnykSeverity(issue.Attributes.EffectiveSeverityLevel)
	status, statusID := mapSnykStatus(issue.Attributes.Status)
	createdAt := issue.Attributes.CreatedAt
	var endTime *time.Time
	if status == "Closed" {
		updatedAt := issue.Attributes.UpdatedAt
		endTime = &updatedAt
	}

	var lastSeenTime time.Time
	if status == "Open" {
		lastSeenTime = issue.Attributes.UpdatedAt
	} else {
		// This technically isn't correct because its when the issue was closed,
		// but we don't have a way to know when the issue was last seen.
		lastSeenTime = issue.Attributes.UpdatedAt
	}

	projectName := project.Attributes.Name
	vendorName := "Snyk"

	var vulnerabilities []ocsf.VulnerabilityDetails
	exploitAvailable := issue.Attributes.ExploitDetails != nil

	var fixAvailable bool
	var remediation *ocsf.Remediation
	for _, coordinate := range issue.Attributes.Coordinates {
		fixAvailable = fixAvailable || coordinate.IsFixableManually || coordinate.IsFixableSnyk ||
			coordinate.IsFixableUpstream || coordinate.IsPatchable || coordinate.IsPinnable || coordinate.IsUpgradeable

		for _, remedy := range coordinate.Remedies {
			if remediation == nil {
				remediation = &ocsf.Remediation{
					Description: remedy.Description,
				}
			} else {
				// Snyk may have multiple remediations for a single issue.
				remediation.Description = fmt.Sprintf("%s\n\nor\n\n%s", remediation.Description, remedy.Description)
			}
		}
	}

	issueURL := fmt.Sprintf("https://app.snyk.io/org/%s/project/%s#issue-%s", s.org.Attributes.Slug, project.ID, issue.Attributes.Key)
	cwe := snykIssueCWE(issue)
	if len(issue.Attributes.Problems) == 0 {
		vulnerabilities = append(vulnerabilities, ocsf.VulnerabilityDetails{
			UID:                &issue.ID,
			CWE:                cwe,
			Desc:               &issue.Attributes.Description,
			Title:              &issue.Attributes.Title,
			Severity:           &severity,
			IsExploitAvailable: &exploitAvailable,
			FirstSeenTime:      &createdAt,
			IsFixAvailable:     &fixAvailable,
			LastSeenTime:       &lastSeenTime,
			VendorName:         &vendorName,
			AffectedCode:       snykAffectedCode(issue, project),
			AffectedPackages:   snykAffectedPackages(issue),
			Remediation:        remediation,
			References:         []string{issueURL},
		})
	} else {
		for _, problem := range issue.Attributes.Problems {
			reference := issueURL
			if problem.URL != nil {
				reference = *problem.URL
			}

			vulnerabilities = append(vulnerabilities, ocsf.VulnerabilityDetails{
				UID:                &problem.ID,
				CVE:                snykProblemToCVE(problem),
				CWE:                cwe,
				AffectedCode:       snykAffectedCode(issue, project),
				AffectedPackages:   snykAffectedPackages(issue),
				Desc:               &issue.Attributes.Description,
				Title:              &issue.Attributes.Title,
				Severity:           &severity,
				IsExploitAvailable: &exploitAvailable,
				FirstSeenTime:      &createdAt,
				IsFixAvailable:     &fixAvailable,
				LastSeenTime:       &lastSeenTime,
				VendorName:         &vendorName,
				Remediation:        remediation,
				References:         []string{reference},
			})
		}
	}

	resourceType := project.Attributes.Type
	resource := ocsf.ResourceDetails{
		UID:  &issue.ID,
		Name: &projectName,
		Type: &resourceType,
	}

	var activityID int32
	var activityName string
	var typeUID int64
	var typeName string
	var eventTime time.Time
	className := "Vulnerability Finding"
	categoryUID := int32(2)
	categoryName := "Findings"
	classUID := int32(2002)

	if createdAt == issue.Attributes.UpdatedAt {
		activityID = int32(1)
		activityName = "Create"
		typeUID = int64(classUID)*100 + int64(activityID)
		typeName = "Vulnerability Finding: Create"
		eventTime = createdAt
	} else if status == "Closed" {
		activityID = int32(3)
		activityName = "Close"
		typeUID = int64(classUID)*100 + int64(activityID)
		typeName = "Vulnerability Finding: Close"
		eventTime = *endTime
	} else {
		activityID = int32(2)
		activityName = "Update"
		typeUID = int64(classUID)*100 + int64(activityID)
		typeName = "Vulnerability Finding: Update"
		eventTime = lastSeenTime
	}

	productName := "Snyk"

	metadata := ocsf.Metadata{
		Product: ocsf.Product{
			Name:       &productName,
			VendorName: productName,
		},
		Version: "1.4.0",
	}

	findingInfo := ocsf.FindingInfo{
		UID:           issue.ID,
		Title:         issue.Attributes.Title,
		Desc:          &issue.Attributes.Description,
		CreatedTime:   &createdAt,
		FirstSeenTime: &createdAt,
		LastSeenTime:  &lastSeenTime,
		ModifiedTime:  &issue.Attributes.UpdatedAt,
		DataSources:   []string{"snyk"},
		Types:         []string{"Vulnerability"},
	}

	finding := ocsf.VulnerabilityFinding{
		Time:            eventTime,
		EventDay:        eventTime,
		StartTime:       &createdAt,
		EndTime:         endTime,
		ActivityID:      activityID,
		ActivityName:    &activityName,
		CategoryUID:     categoryUID,
		CategoryName:    &categoryName,
		ClassUID:        classUID,
		ClassName:       &className,
		Message:         &issue.Attributes.Description,
		Metadata:        metadata,
		Resources:       []ocsf.ResourceDetails{resource},
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
		return "Closed", 4
	default:
		return "Open", 1
	}
}

func snykProblemToCVE(problem snyk.Problem) *ocsf.CVE {
	if problem.Source == "NVD" {
		return &ocsf.CVE{
			UID: problem.ID,
			References: []string{
				*problem.URL,
			},
		}
	}
	return nil
}

func snykIssueCWE(issue snyk.Issue) *ocsf.CWE {
	for _, class := range issue.Attributes.Classes {
		if class.Source == "CWE" {
			return &ocsf.CWE{
				UID:       class.ID,
				SourceURL: class.URL,
			}
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
