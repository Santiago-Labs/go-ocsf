package salesforce

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	SFVersion         = "51.0"
	SFTimestampFormat = "2006-01-02T15:04:05.000-0700"
	SOQLEventLogFile  = "SELECT Id, ApiVersion, EventType, CreatedDate, LogDate, LogFile FROM EventLogFile WHERE EventType = '%s' AND LogDate >= %s"
	DefaultClientID   = "simple-salesforce"
)

var (
	SFOperations = []string{"Login"}

	// Error types
	ErrRequestFailed  = errors.New("request failed")
	ErrConfiguration  = errors.New("configuration error")
	ErrNotFound       = errors.New("not found")
	ErrAuthentication = errors.New("authentication failed")
)

// Client represents the client for interacting with SalesForce
type Client struct {
	Instance       string
	SessionID      string
	Version        string
	identity       string
	key            string
	token          string
	consumerKey    string
	consumerSecret string

	pointer   string
	operation string
}

// New creates a new SalesForce client
func New(identity, key, token string) *Client {
	return &Client{
		identity: identity,
		key:      key,
		token:    token,
		Version:  SFVersion,
	}
}

// WithConsumerCredentials adds OAuth consumer credentials to the client
func (c *Client) WithConsumerCredentials(key, secret string) *Client {
	c.consumerKey = key
	c.consumerSecret = secret
	return c
}

// GetToken fetches the SalesForce token from the client
func (c *Client) GetToken() string {
	return c.token
}

// SalesforceLogin handles authentication with SalesForce
// This implementation focuses solely on the token-based authentication flow
func SalesforceLogin(username, password, securityToken string, session *http.Client) (string, string, error) {
	if session == nil {
		session = &http.Client{}
	}

	// We'll focus only on token login approach for this implementation
	domain := "login" // Default domain

	// Create a login request using the SOAP API
	loginBody := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8" ?>
<env:Envelope
        xmlns:xsd="http://www.w3.org/2001/XMLSchema"
        xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
        xmlns:env="http://schemas.xmlsoap.org/soap/envelope/"
        xmlns:urn="urn:partner.soap.sforce.com">
    <env:Header>
        <urn:CallOptions>
            <urn:client>%s</urn:client>
            <urn:defaultNamespace>sf</urn:defaultNamespace>
        </urn:CallOptions>
    </env:Header>
    <env:Body>
        <n1:login xmlns:n1="urn:partner.soap.sforce.com">
            <n1:username>%s</n1:username>
            <n1:password>%s%s</n1:password>
        </n1:login>
    </env:Body>
</env:Envelope>`, DefaultClientID, username, password, securityToken)

	soapURL := fmt.Sprintf("https://%s.salesforce.com/services/Soap/u/%s", domain, SFVersion)

	// Prepare the request
	req, err := http.NewRequest("POST", soapURL, strings.NewReader(loginBody))
	if err != nil {
		return "", "", fmt.Errorf("%w: %v", ErrRequestFailed, err)
	}

	req.Header.Set("Content-Type", "text/xml")
	req.Header.Set("SOAPAction", "login")

	// Execute the request
	resp, err := session.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("%w: %v", ErrRequestFailed, err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("%w: %v", ErrRequestFailed, err)
	}

	if resp.StatusCode != 200 {
		// Try to extract error from response
		exceptionCode := extractFromXML(string(bodyBytes), "sf:exceptionCode")
		exceptionMsg := extractFromXML(string(bodyBytes), "sf:exceptionMessage")
		if exceptionMsg == "" {
			exceptionMsg = string(bodyBytes)
		}
		return "", "", fmt.Errorf("%w: [%s] %s", ErrAuthentication, exceptionCode, exceptionMsg)
	}

	// Extract session ID and server URL from successful response
	sessionID := extractFromXML(string(bodyBytes), "sessionId")
	serverURL := extractFromXML(string(bodyBytes), "serverUrl")

	if sessionID == "" || serverURL == "" {
		exceptionCode := extractFromXML(string(bodyBytes), "sf:exceptionCode")
		if exceptionCode == "" {
			exceptionCode = "UNKNOWN_EXCEPTION_CODE"
		}
		exceptionMsg := extractFromXML(string(bodyBytes), "sf:exceptionMessage")
		if exceptionMsg == "" {
			exceptionMsg = "UNKNOWN_EXCEPTION_MESSAGE"
		}
		return "", "", fmt.Errorf("%w: [%s] %s", ErrAuthentication, exceptionCode, exceptionMsg)
	}

	// Extract instance from server URL
	sfInstance := serverURL
	sfInstance = strings.Replace(sfInstance, "http://", "", 1)
	sfInstance = strings.Replace(sfInstance, "https://", "", 1)
	sfInstance = strings.Split(sfInstance, "/")[0]
	sfInstance = strings.Replace(sfInstance, "-api", "", 1)

	return sessionID, sfInstance, nil
}

// TokenLogin handles OAuth 2.0 token-based authentication with Salesforce
func TokenLogin(username, password, consumerKey, consumerSecret string, session *http.Client) (string, string, error) {
	if session == nil {
		session = &http.Client{}
	}

	domain := "login" // Default domain
	tokenURL := fmt.Sprintf("https://%s.salesforce.com/services/oauth2/token", domain)

	// Prepare form data
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("client_id", consumerKey)
	data.Set("client_secret", consumerSecret)
	data.Set("username", username)
	data.Set("password", password)

	// Execute the request
	resp, err := session.PostForm(tokenURL, data)
	if err != nil {
		return "", "", fmt.Errorf("%w: %v", ErrRequestFailed, err)
	}
	defer resp.Body.Close()

	var jsonResponse struct {
		AccessToken string `json:"access_token"`
		InstanceURL string `json:"instance_url"`
		Error       string `json:"error"`
		ErrorDesc   string `json:"error_description"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&jsonResponse); err != nil {
		return "", "", fmt.Errorf("%w: %v", ErrRequestFailed, err)
	}

	if resp.StatusCode != 200 {
		return "", "", fmt.Errorf("%w: [%s] %s", ErrAuthentication, jsonResponse.Error, jsonResponse.ErrorDesc)
	}

	if jsonResponse.AccessToken == "" || jsonResponse.InstanceURL == "" {
		return "", "", fmt.Errorf("%w: Missing access_token or instance_url in response", ErrAuthentication)
	}

	// Extract instance from instance URL
	sfInstance := jsonResponse.InstanceURL
	sfInstance = strings.Replace(sfInstance, "http://", "", 1)
	sfInstance = strings.Replace(sfInstance, "https://", "", 1)

	return jsonResponse.AccessToken, sfInstance, nil
}

// extractFromXML is a simple helper to extract values from XML
// Note: For production use, you should use a proper XML parser
func extractFromXML(xml, tag string) string {
	startTag := "<" + tag + ">"
	endTag := "</" + tag + ">"

	startIndex := strings.Index(xml, startTag)
	if startIndex == -1 {
		return ""
	}

	startIndex += len(startTag)
	endIndex := strings.Index(xml[startIndex:], endTag)
	if endIndex == -1 {
		return ""
	}

	return xml[startIndex : startIndex+endIndex]
}

// Authenticate authenticates with Salesforce using the client's credentials
func (c *Client) Authenticate() error {
	var sessionID, instance string
	var err error

	// Use token login if consumer credentials are provided
	if c.consumerKey != "" && c.consumerSecret != "" {
		sessionID, instance, err = TokenLogin(c.identity, c.key, c.consumerKey, c.consumerSecret, nil)
	} else {
		// Fall back to standard login with security token
		sessionID, instance, err = SalesforceLogin(c.identity, c.key, c.token, nil)
	}

	if err != nil {
		return err
	}

	c.SessionID = sessionID
	c.Instance = instance
	return nil
}

// GetPointer retrieves the current pointer value
func (c *Client) GetPointer() (string, error) {
	if c.pointer == "" {
		return "", fmt.Errorf("%w: pointer not found", ErrNotFound)
	}
	return c.pointer, nil
}

// SetPointer updates the pointer value
func (c *Client) SetPointer(pointer string) {
	c.pointer = pointer
}

// WithOperation sets the operation for the connector
func (c *Client) WithOperation(operation string) *Client {
	c.operation = operation
	return c
}

// QueryAll executes a SOQL query against SalesForce
func (c *Client) QueryAll(query string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://%s/services/data/v%s/query/?q=%s",
		c.Instance, c.Version, url.QueryEscape(query))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrRequestFailed, err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.SessionID))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrRequestFailed, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%w: API request failed with status: %d", ErrRequestFailed, resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrRequestFailed, err)
	}

	return result, nil
}

// QueryMore fetches the next page of results
func (c *Client) QueryMore(nextURL string) (map[string]interface{}, error) {
	fullURL := fmt.Sprintf("https://%s%s", c.Instance, nextURL)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrRequestFailed, err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.SessionID))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrRequestFailed, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%w: API request failed with status: %d", ErrRequestFailed, resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrRequestFailed, err)
	}

	return result, nil
}
