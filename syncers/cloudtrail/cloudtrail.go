package cloudtrail

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/Santiago-Labs/go-ocsf/datastore"
	ocsf "github.com/Santiago-Labs/go-ocsf/ocsf/v1_4_0"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
)

type Syncer struct {
	cloudtrailClient *cloudtrail.Client
	datastore        datastore.Datastore[ocsf.APIActivity]
}

type CloudtrailEvent struct {
	EventVersion string `json:"eventVersion"`
	UserIdentity struct {
		Type           string  `json:"type"`
		PrincipalID    string  `json:"principalId"`
		Arn            string  `json:"arn"`
		AccountID      *string `json:"accountId"`
		AccessKeyID    string  `json:"accessKeyId"`
		SessionContext struct {
			SessionIssuer struct {
				Type        string `json:"type"`
				PrincipalID string `json:"principalId"`
				Arn         string `json:"arn"`
				AccountID   string `json:"accountId"`
				UserName    string `json:"userName"`
			} `json:"sessionIssuer"`
			Attributes struct {
				CreationDate     time.Time `json:"creationDate"`
				MfaAuthenticated string    `json:"mfaAuthenticated"`
			} `json:"attributes"`
		} `json:"sessionContext"`
		InvokedBy string `json:"invokedBy"`
	} `json:"userIdentity"`
	EventTime          time.Time `json:"eventTime"`
	EventSource        string    `json:"eventSource"`
	EventName          string    `json:"eventName"`
	AwsRegion          string    `json:"awsRegion"`
	SourceIPAddress    string    `json:"sourceIPAddress"`
	UserAgent          string    `json:"userAgent"`
	ErrorCode          *string   `json:"errorCode"`
	ErrorMessage       *string   `json:"errorMessage"`
	RequestParameters  any       `json:"requestParameters"`
	ResponseElements   any       `json:"responseElements"`
	RequestID          string    `json:"requestID"`
	EventID            string    `json:"eventID"`
	ReadOnly           bool      `json:"readOnly"`
	EventType          string    `json:"eventType"`
	ManagementEvent    bool      `json:"managementEvent"`
	RecipientAccountID string    `json:"recipientAccountId"`
	EventCategory      string    `json:"eventCategory"`
}

func NewSyncer(ctx context.Context, cloudtrailClient *cloudtrail.Client, storageOpts datastore.StorageOpts) (*Syncer, error) {
	dataStoreInst, err := datastore.SetupStorage[ocsf.APIActivity](ctx, storageOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to setup datastore: %w", err)
	}

	return &Syncer{
		cloudtrailClient: cloudtrailClient,
		datastore:        dataStoreInst,
	}, nil
}

func (s *Syncer) Sync(ctx context.Context) error {
	slog.Info("syncing CloudTrail data")

	// CloudTrail limits to 2 calls per second per account per region
	// We'll use pagination to get all events
	var nextToken *string
	var savedActivities, foundActivities int
	var ocsfEvents []ocsf.APIActivity
	const batchSize = 100

	// Rate limiting setup - CloudTrail has a limit of 2 calls per second
	limiter := time.NewTicker(500 * time.Millisecond) // 500ms = 2 calls per second
	defer limiter.Stop()

	for {
		// Wait for rate limit before making API call
		select {
		case <-limiter.C:
			// Proceed with API call
		case <-ctx.Done():
			return ctx.Err()
		}

		cloudtrailEvents, err := s.cloudtrailClient.LookupEvents(ctx, &cloudtrail.LookupEventsInput{
			NextToken: nextToken,
		})
		if err != nil {
			// Add exponential backoff for rate limiting errors
			if retryErr, ok := err.(interface{ RetryAfter() time.Duration }); ok {
				slog.Warn("rate limited by CloudTrail API, backing off", "retry_after", retryErr.RetryAfter())
				select {
				case <-time.After(retryErr.RetryAfter()):
					continue // retry the request
				case <-ctx.Done():
					return ctx.Err()
				}
			}
			return fmt.Errorf("failed to get event history: %w", err)
		}

		foundActivities += len(cloudtrailEvents.Events)

		// Convert CloudTrail events to OCSF format
		for _, event := range cloudtrailEvents.Events {
			ocsfEvent, err := s.ToOCSF(ctx, event)
			if err != nil {
				slog.Warn("failed to convert event to OCSF", "error", err)
				continue
			}
			ocsfEvents = append(ocsfEvents, ocsfEvent)
		}

		// Save in batches
		if len(ocsfEvents) >= batchSize {
			err = s.datastore.Save(ctx, ocsfEvents)
			if err != nil {
				return fmt.Errorf("failed to save OCSF events batch: %w", err)
			}
			savedActivities += len(ocsfEvents)
			ocsfEvents = nil // Reset the slice
		}

		// Check if we have more pages to fetch
		if cloudtrailEvents.NextToken == nil {
			break
		}
		nextToken = cloudtrailEvents.NextToken
	}

	// Save any remaining events
	if len(ocsfEvents) > 0 {
		err := s.datastore.Save(ctx, ocsfEvents)
		if err != nil {
			return fmt.Errorf("failed to save remaining OCSF events: %w", err)
		}
		savedActivities += len(ocsfEvents)
	}

	slog.Info("Finished syncing CloudTrail data", "saved_activities", savedActivities, "found_activities", foundActivities)
	return nil
}

func (s *Syncer) ToOCSF(ctx context.Context, event types.Event) (ocsf.APIActivity, error) {
	// Parse the event data for OCSF conversion
	classUID := 6003
	categoryUID := 6
	categoryName := "Application Activity"
	className := "API Activity"

	var activityID int
	var activityName string
	var typeUID int
	var typeName string

	var cloudtrailEvent CloudtrailEvent
	err := json.Unmarshal([]byte(*event.CloudTrailEvent), &cloudtrailEvent)
	if err != nil {
		return ocsf.APIActivity{}, fmt.Errorf("failed to unmarshal CloudTrail event: %w", err)
	}

	// Determine the activity type based on the event name
	eventName := toString(event.EventName)
	if strings.HasPrefix(eventName, "Create") || strings.HasPrefix(eventName, "Add") ||
		strings.HasPrefix(eventName, "Put") || strings.HasPrefix(eventName, "Insert") {
		activityID = 1
		activityName = "Create"
		typeUID = classUID*100 + activityID
		typeName = "API Activity: Create"
	} else if strings.HasPrefix(eventName, "Get") || strings.HasPrefix(eventName, "Describe") ||
		strings.HasPrefix(eventName, "List") || strings.HasPrefix(eventName, "Search") {
		activityID = 2
		activityName = "Read"
		typeUID = classUID*100 + activityID
		typeName = "API Activity: Read"
	} else if strings.HasPrefix(eventName, "Update") || strings.HasPrefix(eventName, "Modify") ||
		strings.HasPrefix(eventName, "Set") {
		activityID = 3
		activityName = "Update"
		typeUID = classUID*100 + activityID
		typeName = "API Activity: Update"
	} else if strings.HasPrefix(eventName, "Delete") || strings.HasPrefix(eventName, "Remove") {
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

	// Map event success to OCSF status
	status := "Unknown"
	statusID := 0
	// TODO: each response type is different depending on the event source

	if cloudtrailEvent.ErrorCode == nil || toString(cloudtrailEvent.ErrorCode) == "" {
		status = "Success"
		statusID = 1
	} else {
		status = "Failure"
		statusID = 2
	}

	// Set severity based on error information
	severity := "Informational"
	severityID := 1
	if cloudtrailEvent.ErrorCode != nil {
		severity = "Medium"
		severityID = 3
	}

	// Parse actor information
	var actor ocsf.Actor
	if event.Username != nil {
		actor = ocsf.Actor{
			AppName: stringPtr(toString(event.EventSource)),
			User: &ocsf.User{
				Uid:  stringPtr(toString(event.Username)),
				Name: stringPtr(toString(event.Username)),
			},
		}
		acctID := cloudtrailEvent.UserIdentity.AccountID
		if acctID != nil {
			actor.User.Account = &ocsf.Account{
				TypeId: int32Ptr(10), // AWS Account
				Type:   stringPtr("AWS Account"),
				Uid:    stringPtr(toString(event.Username)),
			}
		}
	} else {
		actor = ocsf.Actor{
			AppName: stringPtr(toString(event.EventSource)),
		}
	}

	// Parse API information
	api := ocsf.API{
		Operation: toString(event.EventName),
		Service: &ocsf.Service{
			Name: stringPtr(toString(event.EventSource)),
		},
	}

	// Parse resource information
	var resources []*ocsf.ResourceDetails
	if event.Resources != nil {
		for _, resource := range event.Resources {
			resources = append(resources, &ocsf.ResourceDetails{
				Name: resource.ResourceName,
				Type: resource.ResourceType,
				Uid:  resource.ResourceName,
			})
		}
	}

	// Parse source endpoint information
	var srcEndpoint ocsf.NetworkEndpoint
	if cloudtrailEvent.SourceIPAddress != "" {
		srcEndpoint = ocsf.NetworkEndpoint{
			Ip: stringPtr(cloudtrailEvent.SourceIPAddress),
		}
	} else {
		srcEndpoint = ocsf.NetworkEndpoint{
			SvcName: stringPtr(cloudtrailEvent.EventSource),
		}
	}

	// Parse timestamp
	var ts time.Time
	if event.EventTime != nil {
		ts = *event.EventTime
	} else {
		ts = time.Now()
	}

	// Create the OCSF API Activity
	activity := ocsf.APIActivity{
		ActivityId:   int32(activityID),
		ActivityName: &activityName,
		Actor:        actor,
		Api:          api,
		CategoryName: &categoryName,
		CategoryUid:  int32(categoryUID),
		ClassName:    &className,
		ClassUid:     int32(classUID),
		Status:       &status,
		StatusId:     int32Ptr(int32(statusID)),

		Resources:  resources,
		Severity:   &severity,
		SeverityId: int32(severityID),

		Metadata: ocsf.Metadata{
			CorrelationUid: stringPtr(toString(event.EventId)),
		},

		SrcEndpoint:    srcEndpoint,
		Time:           ts.UnixMilli(),
		EventDay:       int32(ts.UnixMilli() / 86400000),
		TypeName:       &typeName,
		TypeUid:        int64(typeUID),
		TimezoneOffset: int32Ptr(0),
	}

	return activity, nil
}

// Helper functions
func stringPtr(s string) *string {
	return &s
}

func int32Ptr(i int32) *int32 {
	return &i
}

func toString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func toInt32(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}
