package stealthim

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Client represents the HTTP client for API requests
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	Session    string
}

// NewClient creates a new API client
func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// NewClientWithSession creates a new API client with an existing session
func NewClientWithSession(baseURL, session string) *Client {
	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		Session: session,
	}
}

// NewClientFromEnv creates a new API client using the STEALTHIM_SERVER_URL environment variable
func NewClientFromEnv() *Client {
	serverURL := os.Getenv("STEALTHIM_SERVER_URL")
	return &Client{
		BaseURL: serverURL,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// doRequest performs an HTTP request with proper headers and session management
func (c *Client) doRequest(ctx context.Context, method, endpoint string, body any) (*http.Response, error) {
	url := c.BaseURL + endpoint

	var bodyReader io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set content type
	req.Header.Set("Content-Type", "application/json")

	// Set authorization header if session is available
	if c.Session != "" {
		req.Header.Set("Authorization", "Bearer "+c.Session)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	return resp, nil
}

// parseResponse parses the JSON response into the provided result structure
func (c *Client) parseResponse(resp *http.Response, result any) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, result); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil
}
