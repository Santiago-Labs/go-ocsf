package tenable

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	defaultBaseURL = "https://cloud.tenable.com"
	maxRetries     = 5
	retryDelay     = 5 * time.Second
)

// Client represents a Tenable API client
type Client struct {
	baseURL    string
	accessKey  string
	secretKey  string
	httpClient *http.Client
}

// NewClient creates a new Tenable client
func NewClient(accessKey, secretKey string) (*Client, error) {
	if accessKey == "" || secretKey == "" {
		return nil, fmt.Errorf("access key and secret key are required")
	}

	return &Client{
		baseURL:    defaultBaseURL,
		accessKey:  accessKey,
		secretKey:  secretKey,
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}, nil
}

// SetBaseURL sets a custom base URL for the Tenable API
func (c *Client) SetBaseURL(url string) {
	c.baseURL = url
}

// doRequest performs an HTTP request with retries
func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}) ([]byte, error) {
	var bodyReader io.Reader
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}

	url := fmt.Sprintf("%s%s", c.baseURL, path)
	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-ApiKeys", fmt.Sprintf("accessKey=%s; secretKey=%s", c.accessKey, c.secretKey))

	var resp *http.Response
	var respBody []byte

	// Implement retry logic
	for attempt := 0; attempt < maxRetries; attempt++ {
		resp, err = c.httpClient.Do(req)
		if err != nil {
			if attempt < maxRetries-1 {
				time.Sleep(retryDelay)
				continue
			}
			return nil, fmt.Errorf("failed to make request after %d attempts: %w", maxRetries, err)
		}
		defer resp.Body.Close()

		respBody, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}

		// Check for rate limiting (429) or server errors (5xx)
		if resp.StatusCode == http.StatusTooManyRequests || (resp.StatusCode >= 500 && resp.StatusCode < 600) {
			if attempt < maxRetries-1 {
				time.Sleep(retryDelay)
				continue
			}
		}

		break
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var errResp ErrorResponse
		if err := json.Unmarshal(respBody, &errResp); err == nil && errResp.Error != "" {
			return nil, fmt.Errorf("API error: %s - %s (status code: %d)", errResp.Error, errResp.Message, resp.StatusCode)
		}
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

// ExportVulnerabilities initiates an export of vulnerability findings
func (c *Client) ExportVulnerabilities(ctx context.Context, filters map[string]interface{}) (string, error) {
	if filters == nil {
		filters = make(map[string]interface{})
	}

	// Set default filters if not provided
	if _, ok := filters["severity"]; !ok {
		filters["severity"] = []string{"low", "medium", "high", "critical"}
	}
	if _, ok := filters["state"]; !ok {
		filters["state"] = []string{"open", "reopened", "fixed"}
	}

	payload := map[string]interface{}{
		"filters":    filters,
		"num_assets": 5000,
	}

	respBody, err := c.doRequest(ctx, http.MethodPost, "/vulns/export", payload)
	if err != nil {
		return "", fmt.Errorf("failed to initiate vulnerability export: %w", err)
	}

	var exportResp ExportVulnsResponse
	if err := json.Unmarshal(respBody, &exportResp); err != nil {
		return "", fmt.Errorf("failed to parse export response: %w", err)
	}

	return exportResp.ExportUUID, nil
}

// GetExportStatus checks the status of an export job
func (c *Client) GetExportStatus(ctx context.Context, exportUUID string) (*ExportStatusResponse, error) {
	respBody, err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/vulns/export/%s/status", exportUUID), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get export status: %w", err)
	}

	var statusResp ExportStatusResponse
	if err := json.Unmarshal(respBody, &statusResp); err != nil {
		return nil, fmt.Errorf("failed to parse export status response: %w", err)
	}

	return &statusResp, nil
}

// GetExportChunk retrieves a specific chunk of exported data
func (c *Client) GetExportChunk(ctx context.Context, exportUUID string, chunkID int) ([]Finding, error) {
	respBody, err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/vulns/export/%s/chunks/%d", exportUUID, chunkID), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get export chunk: %w", err)
	}

	var chunkResp []Finding
	if err := json.Unmarshal(respBody, &chunkResp); err != nil {
		return nil, fmt.Errorf("failed to parse export chunk response: %w chunkResp: \n%+v\n", err, string(respBody))
	}

	return chunkResp, nil
}

// GetAllFindingsFromExport retrieves all vulnerability findings by handling the export process
func (c *Client) GetAllFindingsFromExport(ctx context.Context, filters map[string]interface{}) ([]Finding, error) {
	exportUUID, err := c.ExportVulnerabilities(ctx, filters)
	if err != nil {
		return nil, err
	}

	// Poll for export completion
	var status *ExportStatusResponse
	for {
		status, err = c.GetExportStatus(ctx, exportUUID)
		if err != nil {
			return nil, err
		}

		if status.Status == "FINISHED" {
			break
		}

		if status.Status == "ERROR" || status.Status == "CANCELLED" {
			return nil, fmt.Errorf("export job failed")
		}

		fmt.Printf("Export status: %s, sleeping for 10 seconds\n", status.Status)

		// Wait before polling again
		time.Sleep(10 * time.Second)
	}

	// Collect all findings from all chunks
	var allFindings []Finding
	for _, chunkID := range status.ChunksAvailable {
		findings, err := c.GetExportChunk(ctx, exportUUID, chunkID)
		if err != nil {
			return nil, err
		}
		allFindings = append(allFindings, findings...)
	}

	return allFindings, nil
}
