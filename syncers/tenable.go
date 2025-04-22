package syncers

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/Santiago-Labs/go-ocsf/clients/tenable"
	"github.com/Santiago-Labs/go-ocsf/datastore"
	"github.com/Santiago-Labs/go-ocsf/ocsf"
	"github.com/samsarahq/go/oops"
)

// TenableOCSFSyncer is responsible for syncing Tenable vulnerability findings to OCSF format
type TenableOCSFSyncer struct {
	tenableClient *tenable.Client
	datastore     datastore.Datastore
}

// NewTenableOCSFSyncer creates a new TenableOCSFSyncer
func NewTenableOCSFSyncer(ctx context.Context, tenableClient *tenable.Client, datastore datastore.Datastore) (DataSync, error) {
	return &TenableOCSFSyncer{
		tenableClient: tenableClient,
		datastore:     datastore,
	}, nil
}

// Sync synchronizes Tenable data with the OCSF datastore
func (s *TenableOCSFSyncer) Sync(ctx context.Context) error {
	slog.Info("syncing Tenable data")

	filters := map[string]interface{}{}

	findings, err := s.tenableClient.GetAllFindingsFromExport(ctx, filters)
	if err != nil {
		return oops.Wrapf(err, "failed to get all findings")
	}

	slog.Info("found Tenable findings", "num_findings", len(findings))

	var findingsToSave []ocsf.VulnerabilityFinding
	for _, finding := range findings {
		ocsfFinding, err := s.ToOCSF(ctx, finding)
		if err != nil {
			return oops.Wrapf(err, "failed to build OCSF finding")
		}

		findingsToSave = append(findingsToSave, ocsfFinding)
	}

	err = s.datastore.SaveFindings(ctx, findingsToSave)
	if err != nil {
		return oops.Wrapf(err, "failed to save findings")
	}

	slog.Info("Finished Tenable sync")
	return nil
}

// ToOCSF converts a Tenable finding into an OCSF vulnerability finding
func (s *TenableOCSFSyncer) ToOCSF(ctx context.Context, finding tenable.Finding) (ocsf.VulnerabilityFinding, error) {
	severity, severityID := mapTenableSeverity(finding.SeverityID)
	status, statusID := mapTenableState(finding.State)

	// Parse first seen time
	var firstSeenTime int64
	parsedTime, err := time.Parse(time.RFC3339, finding.FirstFound)
	if err != nil {
		// If parsing fails, create a time object from the string
		t, _ := time.Parse("2006-01-02 15:04:05", finding.FirstFound)
		firstSeenTime = t.UnixMilli()
	} else {
		firstSeenTime = parsedTime.UnixMilli()
	}

	// Parse last seen time
	var lastSeenTime int64
	parsedTime, err = time.Parse(time.RFC3339, finding.LastFound)
	if err != nil {
		// If parsing fails, create a time object from the string
		t, _ := time.Parse("2006-01-02 15:04:05", finding.LastFound)
		lastSeenTime = t.UnixMilli()
	} else {
		lastSeenTime = parsedTime.UnixMilli()
	}

	var endTime *int64
	if status == "Closed" {
		endTime = &lastSeenTime
	}

	findingID := fmt.Sprintf("tenable-%s", finding.FindingID)

	resourceName := finding.Asset.HostName
	if resourceName == "" {
		resourceName = finding.Asset.IPV4
	}

	var resourceID string
	if finding.Asset.ID == "" {
		resourceID = fmt.Sprintf("tenable-%s", finding.FindingID)
	} else {
		resourceID = fmt.Sprintf("tenable-%s", finding.Asset.ID)
	}
	resourceType := "host"
	vendorName := "Tenable"

	var vulnerabilities []*ocsf.VulnerabilityDetails
	exploitAvailable := finding.Plugin.ExploitAvailable

	var remediation *ocsf.Remediation
	if finding.Plugin.Solution != "" {
		remediation = &ocsf.Remediation{
			Description: finding.Plugin.Solution,
		}
	}

	// Create references
	var references []string
	references = append(references, finding.Plugin.SeeAlso...)

	// Create CVE details
	var cve *ocsf.CVE
	// Check if there are any CVEs in the plugin
	if len(finding.Plugin.Cpe) > 0 {
		// This is a placeholder - we need to extract CVEs from the plugin data
		// In a real implementation, you would need to extract CVEs from the plugin data
		cveID := fmt.Sprintf("PLUGIN-%d", finding.Plugin.ID)
		cve = &ocsf.CVE{
			UID:        cveID,
			References: references,
		}
	}

	vulnerabilities = append(vulnerabilities, &ocsf.VulnerabilityDetails{
		UID:                &findingID,
		CVE:                cve,
		Desc:               &finding.Plugin.Description,
		Title:              &finding.Plugin.Name,
		Severity:           &severity,
		IsExploitAvailable: &exploitAvailable,
		FirstSeenTime:      &firstSeenTime,
		IsFixAvailable:     &finding.Plugin.HasPatch,
		LastSeenTime:       &lastSeenTime,
		VendorName:         &vendorName,
		Remediation:        remediation,
		References:         references,
	})

	resource := ocsf.ResourceDetails{
		UID:  &resourceID,
		Name: &resourceName,
		Type: &resourceType,
	}

	var activityID int32
	var activityName string
	var typeUID int64
	var typeName string
	var eventTime int64
	className := "Vulnerability Finding"
	categoryUID := int32(2)
	categoryName := "Findings"
	classUID := int32(2002)

	if finding.FirstFound == finding.LastFound {
		activityID = int32(1)
		activityName = "Create"
		typeUID = int64(classUID)*100 + int64(activityID)
		typeName = "Vulnerability Finding: Create"
		eventTime = firstSeenTime
	} else if status == "Closed" {
		activityID = int32(3)
		activityName = "Close"
		typeUID = int64(classUID)*100 + int64(activityID)
		typeName = "Vulnerability Finding: Close"
		eventTime = lastSeenTime
	} else {
		activityID = int32(2)
		activityName = "Update"
		typeUID = int64(classUID)*100 + int64(activityID)
		typeName = "Vulnerability Finding: Update"
		eventTime = lastSeenTime
	}

	productName := "Tenable"

	metadata := ocsf.Metadata{
		Product: ocsf.Product{
			Name:       &productName,
			VendorName: productName,
		},
		Version: "1.1.0",
	}

	findingInfo := ocsf.FindingInfo{
		UID:           findingID,
		Title:         finding.Plugin.Name,
		Desc:          &finding.Plugin.Description,
		CreatedTime:   &firstSeenTime,
		FirstSeenTime: &firstSeenTime,
		LastSeenTime:  &lastSeenTime,
		ModifiedTime:  &lastSeenTime,
		DataSources:   []string{"tenable"},
		Types:         []string{"Vulnerability"},
	}

	ocsfFinding := ocsf.VulnerabilityFinding{
		Time:            eventTime,
		EventDay:        int32(eventTime / 86400000),
		StartTime:       &firstSeenTime,
		EndTime:         endTime,
		ActivityID:      activityID,
		ActivityName:    &activityName,
		CategoryUID:     categoryUID,
		CategoryName:    &categoryName,
		ClassUID:        classUID,
		ClassName:       &className,
		Message:         &finding.Plugin.Description,
		Metadata:        metadata,
		Resources:       []*ocsf.ResourceDetails{&resource},
		Status:          &status,
		StatusID:        &statusID,
		TypeUID:         typeUID,
		TypeName:        &typeName,
		Vulnerabilities: vulnerabilities,
		FindingInfo:     findingInfo,
		SeverityID:      int32(severityID),
	}

	return ocsfFinding, nil
}

// mapTenableSeverity maps Tenable severity levels to OCSF severity levels
func mapTenableSeverity(tenableSeverity int) (string, int) {
	switch tenableSeverity {
	case 0:
		return "Informational", 1
	case 1:
		return "Low", 2
	case 2:
		return "Medium", 3
	case 3:
		return "High", 4
	case 4:
		return "Critical", 5
	default:
		return "Unknown", 0
	}
}

// mapTenableState maps Tenable vulnerability states to OCSF status
func mapTenableState(tenableState string) (string, int32) {
	switch tenableState {
	case "fixed":
		return "Closed", 2
	default:
		return "Open", 1
	}
}
