package cufinder

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// Client represents the CUFinder API client
type Client struct {
	apiKey     string
	baseURL    string
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

	// Convert data to form-encoded format
	formData, err := c.StructToFormData(data)
	if err != nil {
		return nil, fmt.Errorf("failed to convert data to form format: %w", err)
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(formData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("x-api-key", c.apiKey)
	req.Header.Set("User-Agent", "cufinder-go/1.0.0")

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

// StructToFormData converts a struct to form-encoded data
func (c *Client) StructToFormData(data interface{}) (string, error) {
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return "", fmt.Errorf("data must be a struct")
	}

	t := v.Type()
	values := url.Values{}

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		// Skip unexported fields
		if !fieldValue.CanInterface() {
			continue
		}

		// Get JSON tag name
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		// Remove omitempty and other options
		jsonTag = strings.Split(jsonTag, ",")[0]

		// Handle different field types
		if fieldValue.Kind() == reflect.Slice {
			// Handle slices (like []string)
			if fieldValue.Len() > 0 {
				var sliceValues []string
				for i := 0; i < fieldValue.Len(); i++ {
					item := fieldValue.Index(i)
					if item.Kind() == reflect.String {
						sliceValues = append(sliceValues, item.String())
					} else {
						sliceValues = append(sliceValues, fmt.Sprintf("%v", item.Interface()))
					}
				}
				// Join slice values with commas
				values.Set(jsonTag, strings.Join(sliceValues, ","))
			}
			continue
		}

		// Convert field value to string
		var strValue string
		switch fieldValue.Kind() {
		case reflect.String:
			strValue = fieldValue.String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			strValue = strconv.FormatInt(fieldValue.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			strValue = strconv.FormatUint(fieldValue.Uint(), 10)
		case reflect.Float32, reflect.Float64:
			strValue = strconv.FormatFloat(fieldValue.Float(), 'f', -1, 64)
		case reflect.Bool:
			strValue = strconv.FormatBool(fieldValue.Bool())
		default:
			// For other types, try to convert to string
			strValue = fmt.Sprintf("%v", fieldValue.Interface())
		}

		// Add non-empty values, but handle different types appropriately
		if fieldValue.Kind() == reflect.String {
			// For strings, only add if not empty
			if strValue != "" {
				values.Set(jsonTag, strValue)
			}
		} else if fieldValue.Kind() == reflect.Bool {
			// For booleans, always add the value (true/false)
			if strValue != "false" {
				values.Set(jsonTag, strconv.FormatBool(fieldValue.Bool()))
			}
		} else if fieldValue.Kind() >= reflect.Int && fieldValue.Kind() <= reflect.Int64 {
			// For integers, only add if not zero
			if fieldValue.Int() != 0 {
				values.Set(jsonTag, strValue)
			}
		} else if fieldValue.Kind() >= reflect.Uint && fieldValue.Kind() <= reflect.Uint64 {
			// For unsigned integers, only add if not zero
			if fieldValue.Uint() != 0 {
				values.Set(jsonTag, strValue)
			}
		} else if fieldValue.Kind() >= reflect.Float32 && fieldValue.Kind() <= reflect.Float64 {
			// For floats, only add if not zero
			if fieldValue.Float() != 0 {
				values.Set(jsonTag, strValue)
			}
		} else {
			// For other types, only add if not empty string representation
			if strValue != "" {
				values.Set(jsonTag, strValue)
			}
		}
	}

	return values.Encode(), nil
}
