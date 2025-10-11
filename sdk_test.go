package cufinder

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSDK(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "application/x-www-form-urlencoded", r.Header.Get("Content-Type"))
		assert.Equal(t, "test-api-key", r.Header.Get("x-api-key"))

		// Return mock responses based on endpoint
		switch r.URL.Path {
		case "/cuf":
			response := map[string]interface{}{
				"domain":       "techcorp.com",
				"query":        "TechCorp",
				"credit_count": 1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/lcuf":
			response := map[string]interface{}{
				"linkedin_url": "https://linkedin.com/company/techcorp",
				"query":        "TechCorp",
				"credit_count": 1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/dtc":
			response := map[string]interface{}{
				"company_name": "TechCorp Inc",
				"query":        "techcorp.com",
				"credit_count": 1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/dte":
			response := map[string]interface{}{
				"emails":       []string{"contact@techcorp.com", "info@techcorp.com"},
				"query":        "techcorp.com",
				"credit_count": 1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/ntp":
			response := map[string]interface{}{
				"phones":       []string{"+1-555-0123", "+1-555-0124"},
				"query":        "TechCorp",
				"credit_count": 1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/rel":
			response := map[string]interface{}{
				"person": map[string]interface{}{
					"full_name": "John Doe",
					"email":     "john.doe@techcorp.com",
					"job_title": "Software Engineer",
				},
				"company": map[string]interface{}{
					"name":     "TechCorp",
					"domain":   "techcorp.com",
					"industry": "Technology",
				},
				"query":        "john.doe@techcorp.com",
				"credit_count": 1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/fcl":
			response := map[string]interface{}{
				"lookalikes": []map[string]interface{}{
					{
						"name":     "DataCorp",
						"domain":   "datacorp.com",
						"industry": "Technology",
					},
					{
						"name":     "SoftCorp",
						"domain":   "softcorp.com",
						"industry": "Technology",
					},
				},
				"query":        "TechCorp",
				"credit_count": 1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/elf":
			response := map[string]interface{}{
				"fundraising": map[string]interface{}{
					"total_funding": 1000000,
					"rounds":        3,
					"last_round":    "Series A",
				},
				"query":        "TechCorp",
				"credit_count": 1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/car":
			response := map[string]interface{}{
				"revenue": map[string]interface{}{
					"annual_revenue": 5000000,
					"currency":       "USD",
					"year":           2023,
				},
				"query":        "TechCorp",
				"credit_count": 1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/fcc":
			response := map[string]interface{}{
				"subsidiaries": []map[string]interface{}{
					{
						"name":   "TechCorp Mobile",
						"domain": "mobile.techcorp.com",
					},
					{
						"name":   "TechCorp Cloud",
						"domain": "cloud.techcorp.com",
					},
				},
				"query":        "TechCorp",
				"credit_count": 1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/fts":
			response := map[string]interface{}{
				"tech_stack": map[string]interface{}{
					"programming_languages": []string{"Go", "Python", "JavaScript"},
					"frameworks":            []string{"React", "Node.js", "Django"},
					"databases":             []string{"PostgreSQL", "Redis"},
				},
				"query":        "TechCorp",
				"credit_count": 1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/epp":
			response := map[string]interface{}{
				"person": map[string]interface{}{
					"full_name":    "John Doe",
					"email":        "john.doe@techcorp.com",
					"job_title":    "Software Engineer",
					"linkedin_url": "https://linkedin.com/in/john-doe",
				},
				"company": map[string]interface{}{
					"name":     "TechCorp",
					"domain":   "techcorp.com",
					"industry": "Technology",
				},
				"query":        "https://linkedin.com/in/john-doe",
				"credit_count": 1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/fwe":
			response := map[string]interface{}{
				"email":        "john.doe@techcorp.com",
				"query":        "https://linkedin.com/in/john-doe",
				"credit_count": 1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/tep":
			response := map[string]interface{}{
				"person": map[string]interface{}{
					"full_name": "John Doe",
					"job_title": "Software Engineer",
					"company":   "TechCorp",
				},
				"query":            "John Doe at TechCorp",
				"confidence_level": 88,
				"credit_count":     1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/enc":
			response := map[string]interface{}{
				"company": map[string]interface{}{
					"name":     "TechCorp",
					"domain":   "techcorp.com",
					"industry": "Technology",
					"size":     "51-200",
				},
				"query":        "TechCorp",
				"credit_count": 1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/cec":
			response := map[string]interface{}{
				"countries":     []string{"US", "UK", "CA"},
				"total_results": 3,
				"query":         "TechCorp",
				"credit_count":  1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/clo":
			response := map[string]interface{}{
				"locations": []map[string]interface{}{
					{
						"country": "US",
						"state":   "CA",
						"city":    "San Francisco",
						"address": "123 Tech St",
					},
					{
						"country": "UK",
						"city":    "London",
						"address": "456 Innovation Ave",
					},
				},
				"query":        "TechCorp",
				"credit_count": 1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/cse":
			response := map[string]interface{}{
				"companies": []map[string]interface{}{
					{
						"name":     "TechCorp",
						"domain":   "techcorp.com",
						"industry": "Technology",
					},
					{
						"name":     "DataCorp",
						"domain":   "datacorp.com",
						"industry": "Data Analytics",
					},
				},
				"total_results": 2,
				"page":          1,
				"query":         "technology",
				"credit_count":  1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/pse":
			response := map[string]interface{}{
				"people": []map[string]interface{}{
					{
						"full_name": "John Doe",
						"job_title": "Software Engineer",
						"company":   "TechCorp",
					},
					{
						"full_name": "Jane Smith",
						"job_title": "Product Manager",
						"company":   "TechCorp",
					},
				},
				"total_results": 2,
				"page":          1,
				"query":         "engineer",
				"credit_count":  1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/lbs":
			response := map[string]interface{}{
				"businesses": []map[string]interface{}{
					{
						"name":    "Coffee Shop",
						"address": "123 Main St",
						"city":    "San Francisco",
					},
					{
						"name":    "Restaurant",
						"address": "456 Oak Ave",
						"city":    "San Francisco",
					},
				},
				"total_results": 2,
				"page":          1,
				"query":         "coffee",
				"credit_count":  1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	// Create SDK with custom config to use mock server
	sdk := NewSDKWithConfig(ClientConfig{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
		Timeout: 5 * time.Second,
	})

	t.Run("CUF Service", func(t *testing.T) {
		result, err := sdk.CUF("TechCorp", "US")
		require.NoError(t, err)
		assert.Equal(t, "techcorp.com", result.Domain)
		assert.Equal(t, "TechCorp", result.Query)
		assert.Equal(t, 1, result.CreditCount)
	})

	t.Run("LCUF Service", func(t *testing.T) {
		result, err := sdk.LCUF("TechCorp")
		require.NoError(t, err)
		assert.Equal(t, "https://linkedin.com/company/techcorp", result.LinkedInURL)
		assert.Equal(t, "TechCorp", result.Query)
	})

	t.Run("DTC Service", func(t *testing.T) {
		result, err := sdk.DTC("techcorp.com")
		require.NoError(t, err)
		assert.Equal(t, "TechCorp Inc", result.CompanyName)
		assert.Equal(t, "techcorp.com", result.Query)
	})

	t.Run("DTE Service", func(t *testing.T) {
		result, err := sdk.DTE("techcorp.com")
		require.NoError(t, err)
		assert.Len(t, result.Emails, 2)
		assert.Contains(t, result.Emails, "contact@techcorp.com")
		assert.Contains(t, result.Emails, "info@techcorp.com")
	})

	t.Run("NTP Service", func(t *testing.T) {
		result, err := sdk.NTP("TechCorp")
		require.NoError(t, err)
		assert.Len(t, result.Phones, 2)
		assert.Contains(t, result.Phones, "+1-555-0123")
		assert.Contains(t, result.Phones, "+1-555-0124")
	})

	t.Run("REL Service", func(t *testing.T) {
		result, err := sdk.REL("john.doe@techcorp.com")
		require.NoError(t, err)
		assert.Equal(t, "John Doe", result.Person.FullName)
		assert.Equal(t, "john.doe@techcorp.com", result.Person.Email)
		assert.Equal(t, "Software Engineer", result.Person.JobTitle)
		assert.Equal(t, "TechCorp", result.Company.Name)
		assert.Equal(t, "techcorp.com", result.Company.Domain)
	})

	t.Run("FCL Service", func(t *testing.T) {
		result, err := sdk.FCL("TechCorp")
		require.NoError(t, err)
		assert.Len(t, result.Lookalikes, 2)
		assert.Equal(t, "DataCorp", result.Lookalikes[0].Name)
		assert.Equal(t, "datacorp.com", result.Lookalikes[0].Domain)
		assert.Equal(t, "SoftCorp", result.Lookalikes[1].Name)
	})

	t.Run("ELF Service", func(t *testing.T) {
		result, err := sdk.ELF("TechCorp")
		require.NoError(t, err)
		assert.NotNil(t, result.Fundraising)
	})

	t.Run("CAR Service", func(t *testing.T) {
		result, err := sdk.CAR("TechCorp")
		require.NoError(t, err)
		assert.NotNil(t, result.Revenue)
	})

	t.Run("FCC Service", func(t *testing.T) {
		result, err := sdk.FCC("TechCorp")
		require.NoError(t, err)
		assert.Len(t, result.Subsidiaries, 2)
		assert.Equal(t, "TechCorp Mobile", result.Subsidiaries[0].Name)
		assert.Equal(t, "mobile.techcorp.com", result.Subsidiaries[0].Domain)
	})

	t.Run("FTS Service", func(t *testing.T) {
		result, err := sdk.FTS("TechCorp")
		require.NoError(t, err)
		assert.NotNil(t, result.TechStack)
	})

	t.Run("EPP Service", func(t *testing.T) {
		result, err := sdk.EPP("https://linkedin.com/in/john-doe")
		require.NoError(t, err)
		assert.Equal(t, "John Doe", result.Person.FullName)
		assert.Equal(t, "john.doe@techcorp.com", result.Person.Email)
		assert.Equal(t, "Software Engineer", result.Person.JobTitle)
		assert.Equal(t, "TechCorp", result.Company.Name)
	})

	t.Run("FWE Service", func(t *testing.T) {
		result, err := sdk.FWE("https://linkedin.com/in/john-doe")
		require.NoError(t, err)
		assert.Equal(t, "john.doe@techcorp.com", result.Email)
	})

	t.Run("TEP Service", func(t *testing.T) {
		result, err := sdk.TEP("John Doe", "TechCorp")
		require.NoError(t, err)
		assert.Equal(t, "John Doe", result.Person.FullName)
		assert.Equal(t, "Software Engineer", result.Person.JobTitle)
		assert.Equal(t, "TechCorp", result.Person.Company)
		assert.Equal(t, 88, result.ConfidenceLevel)
	})

	t.Run("ENC Service", func(t *testing.T) {
		result, err := sdk.ENC("TechCorp")
		require.NoError(t, err)
		assert.Equal(t, "TechCorp", result.Company.Name)
		assert.Equal(t, "techcorp.com", result.Company.Domain)
		assert.Equal(t, "Technology", result.Company.Industry)
		assert.Equal(t, "51-200", result.Company.Size)
	})

	t.Run("CEC Service", func(t *testing.T) {
		result, err := sdk.CEC("TechCorp")
		require.NoError(t, err)
		assert.Len(t, result.Countries, 3)
		assert.Contains(t, result.Countries, "US")
		assert.Contains(t, result.Countries, "UK")
		assert.Contains(t, result.Countries, "CA")
		assert.Equal(t, 3, result.TotalResults)
	})

	t.Run("CLO Service", func(t *testing.T) {
		result, err := sdk.CLO("TechCorp")
		require.NoError(t, err)
		assert.Len(t, result.Locations, 2)
		assert.Equal(t, "US", result.Locations[0]["country"])
		assert.Equal(t, "CA", result.Locations[0]["state"])
		assert.Equal(t, "San Francisco", result.Locations[0]["city"])
	})

	t.Run("CSE Service", func(t *testing.T) {
		result, err := sdk.CSE(CseParams{
			Name:     "technology",
			Country:  "US",
			Industry: "Technology",
		})
		require.NoError(t, err)
		assert.Len(t, result.Companies, 2)
		assert.Equal(t, "TechCorp", result.Companies[0].Name)
		assert.Equal(t, "techcorp.com", result.Companies[0].Domain)
		assert.Equal(t, "Technology", result.Companies[0].Industry)
		assert.Equal(t, 2, result.TotalResults)
		assert.Equal(t, 1, result.Page)
	})

	t.Run("PSE Service", func(t *testing.T) {
		result, err := sdk.PSE(PseParams{
			FullName:    "engineer",
			CompanyName: "TechCorp",
		})
		require.NoError(t, err)
		assert.Len(t, result.People, 2)
		assert.Equal(t, "John Doe", result.People[0].FullName)
		assert.Equal(t, "Software Engineer", result.People[0].JobTitle)
		assert.Equal(t, "TechCorp", result.People[0].Company)
		assert.Equal(t, 2, result.TotalResults)
		assert.Equal(t, 1, result.Page)
	})

	t.Run("LBS Service", func(t *testing.T) {
		result, err := sdk.LBS(LbsParams{
			Name: "coffee",
			City: "San Francisco",
		})
		require.NoError(t, err)
		assert.Len(t, result.Businesses, 2)
		assert.Equal(t, "Coffee Shop", result.Businesses[0].Name)
		assert.Equal(t, "123 Main St", result.Businesses[0].Address)
		assert.Equal(t, "San Francisco", result.Businesses[0].City)
		assert.Equal(t, 2, result.TotalResults)
		assert.Equal(t, 1, result.Page)
	})

	t.Run("Error Handling", func(t *testing.T) {
		// Test missing required parameters
		_, err := sdk.CUF("", "US")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "company_name is required")

		_, err = sdk.CUF("TechCorp", "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "country_code is required")

		_, err = sdk.TEP("", "TechCorp")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "full_name is required")

		_, err = sdk.TEP("John Doe", "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "company is required")
	})
}
