package cufinder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client represents the CUFinder API client
type Client struct {
	apiKey    string
	baseURL   string
	httpClient *http.Client
}

// ClientConfig holds configuration for the client
type ClientConfig struct {
	APIKey     string
	BaseURL    string
	Timeout    time.Duration
	MaxRetries int
}

// NewClient creates a new CUFinder client
func NewClient(config ClientConfig) *Client {
	if config.BaseURL == "" {
		config.BaseURL = "https://api.cufinder.io/v2"
	}
	if config.Timeout == 0 {
		config.Timeout = 30 * time.Second
	}
	if config.MaxRetries == 0 {
		config.MaxRetries = 3
	}

	return &Client{
		apiKey:  config.APIKey,
		baseURL: config.BaseURL,
		httpClient: &http.Client{
			Timeout: config.Timeout,
		},
	}
}

// Post sends a POST request to the API
func (c *Client) Post(endpoint string, data interface{}) (map[string]interface{}, error) {
	url := c.baseURL + endpoint
	
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request data: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API error: status %d, body: %s", resp.StatusCode, string(body))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return result, nil
}
