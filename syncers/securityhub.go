package syncers

import (
	"context"
	"log/slog"
	"time"

	"github.com/Santiago-Labs/go-ocsf/datastore"
	"github.com/Santiago-Labs/go-ocsf/ocsf"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/samsarahq/go/oops"
)

type SecurityHubOCSFSyncer struct {
	securityHubClient *securityhub.Client
	datastore         datastore.Datastore
}

// NewSecurityHubOCSFSyncer creates a new SecurityHubOCSFSyncer
// It initializes the SecurityHub client and datastore.
func NewSecurityHubOCSFSyncer(ctx context.Context, securityHubClient *securityhub.Client, datastore datastore.Datastore) DataSync {
	return &SecurityHubOCSFSyncer{
		securityHubClient: securityHubClient,
		datastore:         datastore,
	}
}

// Sync synchronizes SecurityHub data with the OCSF datastore
// It fetches all findings from SecurityHub, builds OCSF findings, and saves them to the datastore.
func (s *SecurityHubOCSFSyncer) Sync(ctx context.Context) error {
	slog.Info("syncing SecurityHub data")

	var nextToken *string
	for {
		securityHubFindingsOutput, err := s.securityHubClient.GetFindings(
			ctx,
			&securityhub.GetFindingsInput{
				MaxResults: aws.Int32(100),
				SortCriteria: []types.SortCriterion{
					{
						Field:     aws.String("LastObservedAt"),
						SortOrder: types.SortOrderDescending,
					},
				},
				NextToken: nextToken,
			},
		)
		if err != nil {
			return oops.Wrapf(err, "failed to list all findings")
		}

		slog.Info("SecurityHub findings", "num_findings", len(securityHubFindingsOutput.Findings))

		var findingsToSave []ocsf.VulnerabilityFinding
		for _, securityHubFinding := range securityHubFindingsOutput.Findings {
			finding, err := s.ToOCSF(ctx, securityHubFinding)
			if err != nil {
				return oops.Wrapf(err, "failed to build OCSF finding")
			}

			findingsToSave = append(findingsToSave, finding)
		}

		err = s.datastore.SaveFindings(ctx, findingsToSave)
		if err != nil {
			return oops.Wrapf(err, "failed to save findings")
		}

		if securityHubFindingsOutput.NextToken == nil {
			break
		}

		nextToken = securityHubFindingsOutput.NextToken
	}

	slog.Info("Finished SecurityHub sync")
	return nil
}

// ToOCSF converts a SecurityHub finding into an OCSF vulnerability finding.
func (s *SecurityHubOCSFSyncer) ToOCSF(ctx context.Context, securityHubFinding types.AwsSecurityFinding) (ocsf.VulnerabilityFinding, error) {
	severity, severityID := mapSecurityHubSeverity(securityHubFinding.Severity)
	status, statusID := mapSecurityHubStatus(securityHubFinding.Workflow)

	var createdAt *time.Time
	if securityHubFinding.CreatedAt != nil {
		parsedTime, err := time.Parse(time.RFC3339, *securityHubFinding.CreatedAt)
		if err == nil {
			createdAt = &parsedTime
		}
	}

	var endTime *time.Time
	if status == "Closed" {
		if securityHubFinding.UpdatedAt != nil {
			parsedTime, err := time.Parse(time.RFC3339, *securityHubFinding.UpdatedAt)
			if err == nil {
				endTime = &parsedTime
			}
		}
	}

	vendorName := "AWS"
	// SecurityHub Vulnerability doesn't have Exploitability field, so we'll need to check differently
	// or set a default value
	exploitAvailable := false

	var fixAvailable bool
	if securityHubFinding.Remediation != nil && securityHubFinding.Remediation.Recommendation != nil {
		fixAvailable = true
	}

	var remediation *ocsf.Remediation
	if securityHubFinding.Remediation != nil {
		var description string
		if securityHubFinding.Remediation.Recommendation != nil && securityHubFinding.Remediation.Recommendation.Text != nil {
			description = *securityHubFinding.Remediation.Recommendation.Text
		}

		var references []string
		if securityHubFinding.Remediation.Recommendation != nil && securityHubFinding.Remediation.Recommendation.Url != nil {
			references = append(references, *securityHubFinding.Remediation.Recommendation.Url)
		}

		remediation = &ocsf.Remediation{
			Description: description,
			References:  references,
		}
	}

	var title string
	if securityHubFinding.Title != nil {
		title = *securityHubFinding.Title
	}

	// Convert UpdatedAt string to time.Time for LastSeenTime
	var lastSeenTime *time.Time
	if securityHubFinding.UpdatedAt != nil {
		parsedTime, err := time.Parse(time.RFC3339, *securityHubFinding.UpdatedAt)
		if err == nil {
			lastSeenTime = &parsedTime
		}
	}

	vulnerabilities := []ocsf.VulnerabilityDetails{
		{
			UID:                securityHubFinding.Id,
			CWE:                mapSecurityHubCWE(securityHubFinding),
			CVE:                mapSecurityHubCVE(securityHubFinding),
			Desc:               securityHubFinding.Description,
			Title:              &title,
			Severity:           &severity,
			IsExploitAvailable: &exploitAvailable,
			FirstSeenTime:      createdAt,
			IsFixAvailable:     &fixAvailable,
			LastSeenTime:       lastSeenTime,
			VendorName:         &vendorName,
			Remediation:        remediation,
		},
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

	if securityHubFinding.UpdatedAt == securityHubFinding.CreatedAt {
		activityID = int32(1)
		activityName = "Create"
		typeUID = int64(classUID)*100 + int64(activityID)
		typeName = "Vulnerability Finding: Create"
		eventTime = *createdAt
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
		var err error
		eventTime, err = time.Parse(time.RFC3339, *securityHubFinding.UpdatedAt)
		if err != nil {
			return ocsf.VulnerabilityFinding{}, oops.Wrapf(err, "failed to parse time")
		}
	}

	productName := "SecurityHub"

	metadata := ocsf.Metadata{
		Product: ocsf.Product{
			Name:       &productName,
			VendorName: productName,
		},
		Version: "1.4.0",
	}

	var modifiedTime *time.Time
	if securityHubFinding.UpdatedAt != nil {
		parsedTime, err := time.Parse(time.RFC3339, *securityHubFinding.UpdatedAt)
		if err == nil {
			modifiedTime = &parsedTime
		}
	}

	findingInfo := ocsf.FindingInfo{
		UID:           *securityHubFinding.Id,
		Title:         *securityHubFinding.Title,
		Desc:          securityHubFinding.Description,
		CreatedTime:   createdAt,
		FirstSeenTime: createdAt,
		LastSeenTime:  lastSeenTime,
		ModifiedTime:  modifiedTime,
		DataSources:   []string{"securityhub"},
		Types:         []string{"Vulnerability"},
	}

	finding := ocsf.VulnerabilityFinding{
		Time:            eventTime,
		StartTime:       createdAt,
		EventDay:        eventTime,
		EndTime:         endTime,
		ActivityID:      activityID,
		ActivityName:    &activityName,
		CategoryUID:     categoryUID,
		CategoryName:    &categoryName,
		ClassUID:        classUID,
		ClassName:       &className,
		Message:         securityHubFinding.Description,
		Metadata:        metadata,
		Resources:       mapSecurityHubResources(securityHubFinding),
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

func mapSecurityHubSeverity(severity *types.Severity) (string, int) {
	if severity == nil {
		return "Unknown", 0
	}

	// SeverityLabel is an enum, not a pointer
	switch severity.Label {
	case types.SeverityLabelInformational:
		return "Informational", 1
	case types.SeverityLabelLow:
		return "Low", 2
	case types.SeverityLabelMedium:
		return "Medium", 3
	case types.SeverityLabelHigh:
		return "High", 4
	case types.SeverityLabelCritical:
		return "Critical", 5
	default:
		return "Unknown", 0
	}
}

func mapSecurityHubStatus(workflow *types.Workflow) (string, int32) {
	if workflow == nil {
		return "Open", 1
	}

	// WorkflowStatus is an enum, not a pointer
	switch workflow.Status {
	case types.WorkflowStatusNew, types.WorkflowStatusNotified:
		return "Open", 1
	case types.WorkflowStatusSuppressed:
		return "Suppressed", 3
	case types.WorkflowStatusResolved:
		return "Closed", 4
	default:
		return "Unknown", 0
	}
}

func mapSecurityHubResources(finding types.AwsSecurityFinding) []ocsf.ResourceDetails {
	var resources []ocsf.ResourceDetails
	for _, resource := range finding.Resources {
		resourceType := *resource.Type
		resources = append(resources, ocsf.ResourceDetails{
			UID:  resource.Id,
			Type: &resourceType,
		})
	}

	return resources
}

func mapSecurityHubCVE(finding types.AwsSecurityFinding) *ocsf.CVE {
	if finding.Vulnerabilities != nil && len(finding.Vulnerabilities) > 0 {
		for _, vuln := range finding.Vulnerabilities {
			if vuln.Id != nil && vuln.Cvss != nil && len(vuln.Cvss) > 0 {
				var cvss []ocsf.CVSS
				for _, c := range vuln.Cvss {
					if c.BaseScore != nil && c.Version != nil {
						// The field is VectorString, not Vector
						cvss = append(cvss, ocsf.CVSS{
							BaseScore:    *c.BaseScore,
							VectorString: c.BaseVector,
							Version:      *c.Version,
						})
					}
				}

				var references []string
				if vuln.ReferenceUrls != nil {
					references = vuln.ReferenceUrls
				}

				return &ocsf.CVE{
					UID:        *vuln.Id,
					References: references,
					CVSS:       cvss,
				}
			}
		}
	}
	return nil
}

func mapSecurityHubCWE(finding types.AwsSecurityFinding) *ocsf.CWE {
	if finding.Types != nil {
		for _, t := range finding.Types {
			// t is already a string, not a pointer
			if len(t) > 4 && t[:4] == "CWE-" {
				url := "https://cwe.mitre.org/data/definitions/" + t[4:] + ".html"
				return &ocsf.CWE{
					UID:       t,
					SourceURL: &url,
				}
			}
		}
	}
	return nil
}
