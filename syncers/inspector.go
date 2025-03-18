package syncers

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/Santiago-Labs/go-ocsf/datastore"
	"github.com/Santiago-Labs/go-ocsf/ocsf"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/samsarahq/go/oops"
)

type InspectorOCSFSyncer struct {
	inspectorClient *inspector2.Client
	datastore       datastore.Datastore
}

// NewInspectorOCSFSyncer creates a new InspectorOCSFSyncer
// It initializes the Inspector client and datastore.
func NewInspectorOCSFSyncer(ctx context.Context, inspectorClient *inspector2.Client, datastore datastore.Datastore) DataSync {
	return &InspectorOCSFSyncer{
		inspectorClient: inspectorClient,
		datastore:       datastore,
	}
}

// Sync synchronizes Inspector data with the OCSF datastore
// It fetches all findings from Inspector, builds OCSF findings, and saves them to the datastore.
func (s *InspectorOCSFSyncer) Sync(ctx context.Context) error {
	slog.Info("syncing Inspector data")

	var nextToken *string
	for {
		inspectorFindingsOutput, err := s.inspectorClient.ListFindings(
			ctx,
			&inspector2.ListFindingsInput{
				MaxResults: aws.Int32(100),
				SortCriteria: &types.SortCriteria{
					Field:     types.SortFieldLastObservedAt,
					SortOrder: types.SortOrderDesc,
				},
				NextToken: nextToken,
			},
		)
		if err != nil {
			return oops.Wrapf(err, "failed to list all findings")
		}

		slog.Info("Inspector findings", "num_findings", len(inspectorFindingsOutput.Findings))

		var findings []ocsf.VulnerabilityFinding
		for _, inspectorFinding := range inspectorFindingsOutput.Findings {
			existingFinding, err := s.datastore.GetFinding(ctx, *inspectorFinding.FindingArn)
			if err != nil && err != datastore.ErrNotFound {
				return oops.Wrapf(err, "failed to get existing finding")
			}

			finding, err := s.ToOCSF(ctx, inspectorFinding, existingFinding)
			if err != nil {
				return oops.Wrapf(err, "failed to build OCSF finding")
			}

			// Only save the finding if it is new or has changed.
			if existingFinding == nil || existingFinding.SeverityID != finding.SeverityID ||
				existingFinding.StatusID != nil && finding.StatusID == nil ||
				existingFinding.StatusID == nil && finding.StatusID != nil ||
				*existingFinding.StatusID != *finding.StatusID {

				findings = append(findings, finding)
			}
		}

		err = s.datastore.SaveFindings(ctx, findings)
		if err != nil {
			return oops.Wrapf(err, "failed to save findings")
		}

		if inspectorFindingsOutput.NextToken == nil {
			break
		}

		nextToken = inspectorFindingsOutput.NextToken
	}

	slog.Info("Finished Inspector sync")
	return nil
}

// ToOCSF converts a Inspector finding into an OCSF vulnerability finding.
func (s *InspectorOCSFSyncer) ToOCSF(ctx context.Context, inspectorFinding types.Finding, existingFinding *ocsf.VulnerabilityFinding) (ocsf.VulnerabilityFinding, error) {
	severity, severityID := mapInspectorSeverity(inspectorFinding.Severity)
	status, statusID := mapInspectorStatus(inspectorFinding.Status)
	createdAt := inspectorFinding.FirstObservedAt
	var endTime *time.Time

	if status == string(types.FindingStatusClosed) {
		endTime = inspectorFinding.UpdatedAt
	}

	vendorName := "AWS"
	var exploitAvailable bool
	if inspectorFinding.ExploitAvailable == types.ExploitAvailableYes {
		exploitAvailable = true
	} else {
		exploitAvailable = false
	}

	var fixAvailable bool
	if inspectorFinding.FixAvailable == types.FixAvailableYes {
		fixAvailable = true
	} else {
		fixAvailable = false
	}

	var remediation *ocsf.Remediation
	if inspectorFinding.Remediation != nil {
		var description string
		if inspectorFinding.Remediation.Recommendation != nil && inspectorFinding.Remediation.Recommendation.Text != nil {
			description = *inspectorFinding.Remediation.Recommendation.Text
		}

		var references []string
		if inspectorFinding.Remediation.Recommendation != nil && inspectorFinding.Remediation.Recommendation.Url != nil {
			references = append(references, *inspectorFinding.Remediation.Recommendation.Url)
		}

		remediation = &ocsf.Remediation{
			Description: description,
			References:  references,
		}
	}

	var title string
	if inspectorFinding.Title != nil {
		title = *inspectorFinding.Title
	}

	vulnerabilities := []ocsf.VulnerabilityDetails{
		{
			UID:                inspectorFinding.FindingArn,
			CWE:                mapInspectorCWE(inspectorFinding),
			CVE:                mapInspectorCVE(inspectorFinding),
			Desc:               inspectorFinding.Description,
			Title:              &title,
			Severity:           &severity,
			IsExploitAvailable: &exploitAvailable,
			FirstSeenTime:      createdAt,
			IsFixAvailable:     fixAvailable,
			LastSeenTime:       inspectorFinding.LastObservedAt,
			VendorName:         &vendorName,
			AffectedCode:       mapInspectorAffectedCode(inspectorFinding),
			AffectedPackages:   mapInspectorAffectedPackages(inspectorFinding),
			Remediation:        remediation,
		},
	}

	var activityID int32
	var activityName string
	var typeUID int64
	var typeName string
	className := "Vulnerability Finding"
	categoryUID := int32(2)
	categoryName := "Findings"
	classUID := int32(2002)

	if existingFinding == nil {
		activityID = int32(1)
		activityName = "Create"
		typeUID = int64(classUID)*100 + int64(activityID)
		typeName = "Vulnerability Finding: Create"
	} else {
		if status == "Closed" {
			activityID = int32(3)
			activityName = "Close"
			typeUID = int64(classUID)*100 + int64(activityID)
			typeName = "Vulnerability Finding: Close"
		} else {
			activityID = int32(2)
			activityName = "Update"
			typeUID = int64(classUID)*100 + int64(activityID)
			typeName = "Vulnerability Finding: Update"
		}
	}

	productName := "Inspector"

	metadata := ocsf.Metadata{
		Product: ocsf.Product{
			Name:       &productName,
			VendorName: productName,
		},
		Version: "1.4.0",
	}

	now := time.Now()

	findingInfo := ocsf.FindingInfo{
		UID:           *inspectorFinding.FindingArn,
		Title:         *inspectorFinding.Title,
		Desc:          inspectorFinding.Description,
		CreatedTime:   &now,
		FirstSeenTime: inspectorFinding.FirstObservedAt,
		LastSeenTime:  inspectorFinding.LastObservedAt,
		ModifiedTime:  inspectorFinding.UpdatedAt,
		DataSources:   []string{"inspector"},
		Types:         []string{"Vulnerability"},
	}

	finding := ocsf.VulnerabilityFinding{
		Time:            time.Now(),
		StartTime:       inspectorFinding.FirstObservedAt,
		EndTime:         endTime,
		ActivityID:      activityID,
		ActivityName:    &activityName,
		CategoryUID:     categoryUID,
		CategoryName:    &categoryName,
		ClassUID:        classUID,
		ClassName:       &className,
		Message:         inspectorFinding.Description,
		Metadata:        metadata,
		Resources:       mapInspectorResources(inspectorFinding),
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

func mapInspectorSeverity(severity types.Severity) (string, int) {
	switch severity {
	case types.SeverityInformational:
		return "Informational", 1
	case types.SeverityLow:
		return "Low", 2
	case types.SeverityMedium:
		return "Medium", 3
	case types.SeverityHigh:
		return "High", 4
	case types.SeverityCritical:
		return "Critical", 5
	default:
		return "Unknown", 0
	}
}

func mapInspectorStatus(status types.FindingStatus) (string, int32) {
	switch status {
	case types.FindingStatusActive:
		return "Open", 1
	case types.FindingStatusSuppressed:
		return "Suppressed", 3
	case types.FindingStatusClosed:
		return "Closed", 4
	default:
		return "Unknown", 0
	}
}

func mapInspectorResources(finding types.Finding) []ocsf.ResourceDetails {
	var resources []ocsf.ResourceDetails
	for _, resource := range finding.Resources {

		resourceType := string(resource.Type)
		resources = append(resources, ocsf.ResourceDetails{
			UID:  resource.Id,
			Type: &resourceType,
		})
	}

	return resources
}

func mapInspectorCVE(finding types.Finding) *ocsf.CVE {
	if finding.PackageVulnerabilityDetails != nil && finding.PackageVulnerabilityDetails.VulnerabilityId != nil {
		var cvss []ocsf.CVSS
		for _, c := range finding.PackageVulnerabilityDetails.Cvss {
			cvss = append(cvss, ocsf.CVSS{
				BaseScore:    *c.BaseScore,
				VectorString: c.ScoringVector,
				Version:      *c.Version,
			})
		}

		return &ocsf.CVE{
			UID:        *finding.PackageVulnerabilityDetails.VulnerabilityId,
			References: finding.PackageVulnerabilityDetails.ReferenceUrls,
			CVSS:       cvss,
		}
	}
	return nil
}

func mapInspectorCWE(finding types.Finding) *ocsf.CWE {
	if finding.CodeVulnerabilityDetails != nil && finding.CodeVulnerabilityDetails.Cwes != nil {
		for _, cwe := range finding.CodeVulnerabilityDetails.Cwes {

			url := fmt.Sprintf("https://cwe.mitre.org/data/definitions/%s.html", strings.TrimPrefix(cwe, "CWE-"))
			return &ocsf.CWE{
				UID:       cwe,
				SourceURL: &url,
			}
		}
	}
	return nil
}

func mapInspectorAffectedCode(finding types.Finding) []ocsf.AffectedCode {
	var affectedCode []ocsf.AffectedCode

	if finding.CodeVulnerabilityDetails != nil {
		startLine := int32(0)
		endLine := int32(0)
		var filePath string

		if finding.CodeVulnerabilityDetails.FilePath != nil {
			if finding.CodeVulnerabilityDetails.FilePath.StartLine != nil {
				startLine = *finding.CodeVulnerabilityDetails.FilePath.StartLine
			}
			if finding.CodeVulnerabilityDetails.FilePath.EndLine != nil {
				endLine = *finding.CodeVulnerabilityDetails.FilePath.EndLine
			}

			if finding.CodeVulnerabilityDetails.FilePath.FilePath != nil {
				filePath = *finding.CodeVulnerabilityDetails.FilePath.FilePath
			}
		}

		affectedCode = append(affectedCode, ocsf.AffectedCode{
			File: ocsf.File{
				Path: filePath,
			},
			StartLine: startLine,
			EndLine:   endLine,
		})
	}

	return affectedCode
}

func mapInspectorAffectedPackages(finding types.Finding) []ocsf.AffectedSoftwarePackage {
	var affectedPackages []ocsf.AffectedSoftwarePackage

	if finding.PackageVulnerabilityDetails != nil {
		pkg := finding.PackageVulnerabilityDetails.VulnerablePackages
		for _, p := range pkg {

			packageManager := string(p.PackageManager)
			epoch := p.Epoch

			var remediation *ocsf.Remediation
			if p.Remediation != nil {
				var remediationDescription string
				if p.Remediation != nil {
					remediationDescription = *p.Remediation
				}
				remediation = &ocsf.Remediation{
					Description: remediationDescription,
				}
			}

			affectedPackages = append(affectedPackages, ocsf.AffectedSoftwarePackage{
				Name:           *p.Name,
				Version:        *p.Version,
				Architecture:   p.Arch,
				PackageManager: &packageManager,
				Release:        p.Release,
				Path:           p.FilePath,
				FixedInVersion: p.FixedInVersion,
				Epoch:          &epoch,
				Remediation:    remediation,
			})
		}
	}

	return affectedPackages
}
