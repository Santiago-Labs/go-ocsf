// Package snyk provides a client for interacting with the Snyk API.
// The Snyk API is documented at https://snyk.docs.apiary.io/

package snyk

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"encoding/json"
)

// Client represents a Snyk API client that can interact with Snyk resources.
// It handles authentication and communication with the Snyk API.
type Client struct {
	orgID      string
	apiKey     string
	httpClient *http.Client
}

// NewClient creates a new Snyk client with the provided API key and organization ID.
// It requires an API key and Snyk organization ID.
// Returns an initialized client and any error encountered during initialization.
func NewClient(apiKey, snykOrgID string) (*Client, error) {
	return &Client{
		apiKey:     apiKey,
		httpClient: &http.Client{},
		orgID:      snykOrgID,
	}, nil
}

// GetOrg retrieves information about the Snyk organization associated with this client.
func (c *Client) GetOrg(ctx context.Context) (*Org, error) {
	url := fmt.Sprintf("https://api.snyk.io/rest/orgs/%s?version=2024-10-15", c.orgID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "token "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var org OrgResponse
	if err := json.NewDecoder(resp.Body).Decode(&org); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if org.Data.ID == "" {
		return nil, fmt.Errorf("no org data found")
	}

	return &org.Data, nil
}

// GetProject retrieves information about a specific Snyk project.
func (c *Client) GetProject(ctx context.Context, projectID string) (*Project, error) {
	url := fmt.Sprintf("https://api.snyk.io/rest/orgs/%s/projects/%s?version=2024-10-15", c.orgID, projectID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "token "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var project ProjectResponse
	if err := json.NewDecoder(resp.Body).Decode(&project); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if project.Data.ID == "" {
		return nil, fmt.Errorf("no project data found")
	}

	return &project.Data, nil
}

// ParseSnykTestFile reads and parses a Snyk test result file from the specified filepath.
func (c *Client) ParseSnykTestFile(ctx context.Context, filepath string) (SnykTestAllProjectsResult, error) {
	jsonFile, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var snykTestResult SnykTestAllProjectsResult
	if err := json.Unmarshal(jsonFile, &snykTestResult); err != nil {
		return nil, err
	}
	return snykTestResult, nil
}

// ListIssues retrieves all vulnerability issues from the Snyk organization.
func (c *Client) ListIssues(ctx context.Context) ([]Issue, error) {
	var allIssues []Issue

	nextURL := fmt.Sprintf(
		"https://api.snyk.io/rest/orgs/%s/issues?version=2024-10-15&limit=100",
		c.orgID,
	)

	for {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, nextURL, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}
		req.Header.Set("Authorization", "token "+c.apiKey)

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("failed to make request: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}

		var issueResponse IssuesResponse
		if err := json.NewDecoder(resp.Body).Decode(&issueResponse); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}

		allIssues = append(allIssues, issueResponse.Data...)

		var newNext string
		if issueResponse.Links.Next.StringVal != nil {
			newNext = *issueResponse.Links.Next.StringVal
		} else if issueResponse.Links.Next.ObjectVal != nil {
			newNext = issueResponse.Links.Next.ObjectVal.Href
		}

		if newNext == "" {
			break
		}
		if strings.HasPrefix(newNext, "/") {
			nextURL = fmt.Sprintf("https://api.snyk.io%s", newNext)
		} else {
			nextURL = newNext
		}
	}

	return allIssues, nil
}

// GetIssue retrieves detailed information about a specific vulnerability issue.
func (c *Client) GetIssue(ctx context.Context, issueID string) (*Issue, error) {
	url := fmt.Sprintf("https://api.snyk.io/rest/orgs/%s/issues/%s?version=2024-06-10", c.orgID, issueID)

	// Create request
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Authorization", "token "+c.apiKey)

	// Make request using the client's httpClient
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read error response body: %w", err)
		}
		return nil, fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(body))
	}

	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &issue, nil
}
