package gcpauditlog

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/Santiago-Labs/go-ocsf/clients/gcp"
	"github.com/Santiago-Labs/go-ocsf/datastore"
	"github.com/Santiago-Labs/go-ocsf/ocsf"
	"google.golang.org/api/iterator"
)

type GCPAuditLogSyncer struct {
	datastore datastore.Datastore
	projectID string
	client    *gcp.Client
}

func NewGCPAuditLogSyncer(ctx context.Context, datastore datastore.Datastore, projectID string) (*GCPAuditLogSyncer, error) {
	if projectID == "" {
		return nil, errors.New("projectID is required it can be set via the GCP_PROJECT_ID environment variable")
	}

	client, err := gcp.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCP client: %w", err)
	}

	return &GCPAuditLogSyncer{
		datastore: datastore,
		projectID: projectID,
		client:    client,
	}, nil
}

func (s *GCPAuditLogSyncer) Sync(ctx context.Context) error {
	slog.Info("syncing GCP audit logs")

	// TODO: Make the time range of logs configurable. Default to 1 week for now.
	lastWeek := time.Now().Add(-24 * 7 * time.Hour).Format(time.RFC3339)

	filter := fmt.Sprintf(`protoPayload.@type = "type.googleapis.com/google.cloud.audit.AuditLog" AND
	timestamp >= "%s"
	`, lastWeek)

	it := s.client.ListAuditLogsIterator(ctx, filter)

	var savedActivities, foundActivities int
	var activitiesToSave []ocsf.APIActivity
	const batchSize = 1000

	for {
		log, err := it.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return fmt.Errorf("failed to list audit logs: %w", err)
		}
		foundActivities++

		activity, err := s.ToOCSF(ctx, log)
		if err != nil {
			return fmt.Errorf("failed to build OCSF activity: %w", err)
		}

		existingActivity, err := s.datastore.GetAPIActivity(ctx, *activity.Metadata.CorrelationUID)
		if err != nil && err != datastore.ErrNotFound {
			return fmt.Errorf("failed to get existing activity: %w", err)
		}

		if existingActivity != nil {
			continue
		}

		activitiesToSave = append(activitiesToSave, activity)

		// Save in batches of batchSize
		if len(activitiesToSave) >= batchSize {
			err = s.datastore.SaveAPIActivities(ctx, activitiesToSave)
			if err != nil {
				return fmt.Errorf("failed to save API activities batch: %w", err)
			}
			savedActivities += len(activitiesToSave)
			activitiesToSave = nil // Reset the slice
		}
	}

	// Save any remaining activities
	if len(activitiesToSave) > 0 {
		err := s.datastore.SaveAPIActivities(ctx, activitiesToSave)
		if err != nil {
			return fmt.Errorf("failed to save remaining API activities: %w", err)
		}
		savedActivities += len(activitiesToSave)
	}

	slog.Info("Finished GCP audit log sync", "saved_activities", savedActivities, "found_activities", foundActivities)
	return nil
}

// ToOCSF converts a GCP audit log into an OCSF API activity.
func (s *GCPAuditLogSyncer) ToOCSF(ctx context.Context, log *gcp.AuditLog) (ocsf.APIActivity, error) {
	// Parse the timestamp
	methodName := log.AuditLog.MethodName
	classUID := 6003
	categoryUID := 6
	categoryName := "Application Activity"
	className := "API Activity"

	var activityID int
	var activityName string
	var typeUID int
	var typeName string

	// Check for common method name patterns
	if strings.Contains(methodName, "create") || strings.Contains(methodName, "Create") ||
		strings.Contains(methodName, "insert") || strings.Contains(methodName, "Insert") ||
		strings.Contains(methodName, "cloudsql.instances.automatedBackup") {
		activityID = 1
		activityName = "Create"
		typeUID = classUID*100 + activityID
		typeName = "API Activity: Create"
	} else if strings.Contains(methodName, "get") || strings.Contains(methodName, "Get") ||
		strings.Contains(methodName, "list") || strings.Contains(methodName, "List") {
		activityID = 2
		activityName = "Read"
		typeUID = classUID*100 + activityID
		typeName = "API Activity: Read"
	} else if strings.Contains(methodName, "update") || strings.Contains(methodName, "Update") ||
		strings.Contains(methodName, "modify") || strings.Contains(methodName, "Modify") {
		activityID = 3
		activityName = "Update"
		typeUID = classUID*100 + activityID
		typeName = "API Activity: Update"
	} else if strings.Contains(methodName, "delete") || strings.Contains(methodName, "Delete") {
		activityID = 4
		activityName = "Delete"
		typeUID = classUID*100 + activityID
		typeName = "API Activity: Delete"
	} else {
		activityID = 0
		activityName = "Unknown"
		typeUID = classUID*100 + activityID
		typeName = "API Activity: Unknown"
	}

	status, statusID := mapGCPStatusToOCSFStatus(int32(log.Log.Severity))
	severity, severityID := mapGCPSeverityToOCSFStatus(int32(log.Severity))

	var actor ocsf.Actor
	// For the actor, we expect a process or user for the most part. We have the
	// strict constraint that at least one of these attributes must be present:
	// app_name, app_uid, invoked_by, process, session, user.

	// If there is authentication info, we have a user. Otherwise, we have a process.
	if log.AuditLog.GetAuthenticationInfo() != nil && log.AuditLog.GetAuthenticationInfo().GetPrincipalEmail() != "" {
		actor = ocsf.Actor{
			User: &ocsf.User{
				Account: &ocsf.Account{
					TypeID: intPtr(11),
					Type:   stringPtr("GCP Project"),
					// It's possible in the config that we could get the actual project name.
					UID: stringPtr(projectIDFromLogName(log.Log.LogName)),
				},
			},
			Process: &ocsf.Process{
				Name: stringPtr(log.AuditLog.GetAuthenticationInfo().GetPrincipalEmail()),
				UID:  stringPtr(log.AuditLog.GetAuthenticationInfo().GetPrincipalEmail()),
			},
		}
	} else {
		actor = ocsf.Actor{
			// The service will have invoked the process.
			InvokedBy: stringPtr(log.AuditLog.ServiceName),
		}
	}

	var api ocsf.API
	if log.AuditLog.GetRequestMetadata() != nil && log.AuditLog.GetRequestMetadata().RequestAttributes != nil {
		var operation string
		if log.AuditLog.GetRequestMetadata().RequestAttributes.Method != "" {
			operation = log.AuditLog.GetRequestMetadata().RequestAttributes.Method
		} else if log.AuditLog.MethodName != "" {
			operation = log.AuditLog.MethodName
		} else {
			operation = "Unknown"
		}

		api = ocsf.API{
			Operation: operation,
			Service: &ocsf.Service{
				Name: stringPtr(log.AuditLog.ServiceName),
			},
		}
	}

	var resources []*ocsf.ResourceDetails
	if log.AuditLog.GetResourceName() != "" {
		resources = append(resources, &ocsf.ResourceDetails{
			Name: stringPtr(log.AuditLog.GetResourceName()),
			Type: stringPtr(log.Log.Resource.Type),
		})
	}

	ts := log.Timestamp

	// SrcEndpoint is either a service or something with a mappable IP address.
	// We need to map this to a network endpoint.
	var srcEndpoint ocsf.NetworkEndpoint

	// If there is a caller IP, we can use that.
	if log.AuditLog.RequestMetadata.CallerIp != "" {
		srcEndpoint = ocsf.NetworkEndpoint{
			IP: stringPtr(log.AuditLog.RequestMetadata.CallerIp),
		}
	} else if log.AuditLog.ServiceName != "" {
		srcEndpoint = ocsf.NetworkEndpoint{
			SvcName: stringPtr(log.AuditLog.ServiceName),
		}
	}

	// Create the API Activity
	activity := ocsf.APIActivity{
		ActivityID:   activityID,
		ActivityName: &activityName,
		Actor:        actor,
		API:          api,
		CategoryName: &categoryName,
		CategoryUID:  categoryUID,
		ClassName:    &className,
		ClassUID:     classUID,
		Status:       &status,
		StatusID:     int(statusID),

		Resources:  resources,
		Severity:   &severity,
		SeverityID: severityID,

		Metadata: ocsf.Metadata{
			CorrelationUID: stringPtr(log.ID),
		},

		SrcEndpoint:    srcEndpoint,
		Time:           ts,
		TypeName:       &typeName,
		TypeUID:        typeUID,
		TimezoneOffset: 0,
	}

	return activity, nil
}

func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}

func projectIDFromLogName(logName string) string {
	// projects/my-project/logs/compute.googleapis.com%2Faudit.log
	parts := strings.Split(logName, "/")
	return parts[1]
}

func mapGCPSeverityToOCSFStatus(severity int32) (string, int) {
	switch severity {
	case 0: // DEFAULT
		return "Unknown", 0
	case 100, 200: // DEBUG, INFO
		return "Informational", 1
	case 300, 400: // NOTICE, WARNING
		return "Low", 2
	case 500: // ERROR
		return "Medium", 3
	case 600: // CRITICAL
		return "High", 4
	case 700: // ALERT
		return "Critical", 5
	case 800: // EMERGENCY
		return "Fatal", 6
	default:
		return "Unknown", 0
	}
}

func mapGCPStatusToOCSFStatus(severity int32) (string, int) {
	switch severity {
	case 0: // DEFAULT
		return "Unknown", 0
	case 100, 200: // DEBUG, INFO
		return "Success", 1
	case 300, 400, 500, 600, 700, 800: // NOTICE, WARNING, ERROR, CRITICAL, ALERT, EMERGENCY
		return "Failure", 2
	default:
		return "Unknown", 0
	}
}
