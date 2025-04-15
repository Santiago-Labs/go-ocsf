package gcp

import (
	"context"
	"fmt"
	"time"

	logging "cloud.google.com/go/logging/apiv2"
	"cloud.google.com/go/logging/apiv2/loggingpb"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/genproto/googleapis/api/monitoredres"
	"google.golang.org/genproto/googleapis/cloud/audit"
)

// Client represents a GCP client for interacting with Google Cloud services
type Client struct {
	projectID string
	logClient *logging.Client
}

// NewClient creates a new GCP client
func NewClient(ctx context.Context, projectID string, opts ...option.ClientOption) (*Client, error) {
	logClient, err := logging.NewClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create logging client: %w", err)
	}

	return &Client{
		projectID: projectID,
		logClient: logClient,
	}, nil
}

// Close closes the client connections
func (c *Client) Close() error {
	return c.logClient.Close()
}

// AuditLog represents a simplified audit log entry
type AuditLog struct {
	LogName   string
	Severity  int32
	AuditLog  *audit.AuditLog
	Resource  *monitoredres.MonitoredResource
	Timestamp time.Time
	ID        string
	Log       *loggingpb.LogEntry
}

// LogEntryIterator wraps the logging iterator to provide a cleaner interface
type LogEntryIterator struct {
	iterator *logging.LogEntryIterator
}

// Next returns the next log entry or nil if done
func (it *LogEntryIterator) Next() (*AuditLog, error) {
	entry, err := it.iterator.Next()
	if err != nil {
		return nil, err
	}

	payload, ok := entry.Payload.(*loggingpb.LogEntry_ProtoPayload)
	if !ok {
		return nil, fmt.Errorf("unexpected payload type: %T", entry.Payload)
	}

	auditLog := &audit.AuditLog{}
	if err := payload.ProtoPayload.UnmarshalTo(auditLog); err != nil {
		return nil, fmt.Errorf("failed to unmarshal protoPayload: %w", err)
	}

	return &AuditLog{
		AuditLog:  auditLog,
		Severity:  int32(entry.Severity),
		Timestamp: entry.Timestamp.AsTime(),
		ID:        entry.InsertId,
		Log:       entry,
	}, nil
}

// ListAuditLogsIterator returns an iterator for audit logs
func (c *Client) ListAuditLogsIterator(ctx context.Context, filter string) *LogEntryIterator {
	req := &loggingpb.ListLogEntriesRequest{
		ResourceNames: []string{fmt.Sprintf("projects/%s", c.projectID)},
		Filter:        filter,
		OrderBy:       "timestamp desc",
	}

	it := c.logClient.ListLogEntries(ctx, req)
	return &LogEntryIterator{iterator: it}
}

// ListAuditLogs retrieves audit logs from the specified time range
func (c *Client) ListAuditLogs(ctx context.Context, filter string) ([]*AuditLog, error) {
	it := c.ListAuditLogsIterator(ctx, filter)

	var auditLogs []*AuditLog
	for {
		auditLog, err := it.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, fmt.Errorf("failed to get next log entry: %w", err)
		}

		auditLogs = append(auditLogs, auditLog)
	}

	return auditLogs, nil
}
