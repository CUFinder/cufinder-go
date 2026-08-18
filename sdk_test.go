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
				"companies": []map[string]interface{}{
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
				"subsidiaries": []string{"TechCorp Mobile", "TechCorp Cloud"},
				"query":        "TechCorp",
				"credit_count": 1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/fts":
			response := map[string]interface{}{
				"technologies": []string{"Go", "Python", "JavaScript"},
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
				"work_email":   "john.doe@techcorp.com",
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
				"peoples": []map[string]interface{}{
					{
						"full_name": "John Doe",
						"current_job": map[string]interface{}{
							"title": "Software Engineer",
						},
						"company": map[string]interface{}{
							"name": "TechCorp",
						},
					},
					{
						"full_name": "Jane Smith",
						"current_job": map[string]interface{}{
							"title": "Product Manager",
						},
						"company": map[string]interface{}{
							"name": "TechCorp",
						},
					},
				},
				"query":        "engineer",
				"credit_count": 1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/lbs":
			response := map[string]interface{}{
				"companies": []map[string]interface{}{
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
				"query":        "coffee",
				"credit_count": 1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/naa":
			response := map[string]interface{}{
				"address":      "123 Main St, San Francisco, CA 94105, US",
				"query":        "1095 avenue of the Americas, 6th Avenue ny 10036",
				"credit_count": 1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/cef":
			response := map[string]interface{}{
				"employees": []map[string]interface{}{
					{
						"full_name":    "John Doe",
						"first_name":   "John",
						"last_name":    "Doe",
						"linkedin_url": "https://linkedin.com/in/john-doe",
						"job_title":    "Software Engineer",
						"company_name": "TechCorp",
						"country":      "US",
						"state":        "CA",
						"city":         "San Francisco",
					},
					{
						"full_name":    "Jane Smith",
						"first_name":   "Jane",
						"last_name":    "Smith",
						"linkedin_url": "https://linkedin.com/in/jane-smith",
						"job_title":    "Product Manager",
						"company_name": "TechCorp",
						"country":      "US",
						"state":        "CA",
						"city":         "San Francisco",
					},
				},
				"query":        "TechCorp",
				"credit_count": 1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/nac":
			response := map[string]interface{}{
				"company":      "Cufinder Inc.",
				"query":        "cufinder inc.",
				"credit_count": 1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/caa":
			response := map[string]interface{}{
				"activities": []map[string]interface{}{
					{
						"activity_url":             "https://www.linkedin.com/posts/cufinder_anthropic-alphabet-ai-activity-7462132400869888002-bCRi",
						"activity_id":              "7462132400869888002",
						"activity_comments_count":  59,
						"activity_hashtags":        []string{"#AI", "#Anthropic", "#CUFinder"},
						"activity_is_video":        true,
						"activity_posted_at":       "2026-05-18T13:30:04.063Z",
						"activity_reactions_count": 3,
						"activity_text":            "Anthropic is projected to grow 222x by 2030",
						"activity_top_comments":    []string{},
						"activity_images":          []string{"https://media.licdn.com/image.jpg"},
						"activity_videos":          []string{"https://dms.licdn.com/video.mp4"},
					},
				},
				"query":        "TechCorp",
				"credit_count": 1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/cja":
			response := map[string]interface{}{
				"jobs": []map[string]interface{}{
					{
						"company": map[string]interface{}{
							"name":            "google",
							"industry":        "software development",
							"website":         "https://google.com",
							"linkedin":        "linkedin.com/company/google",
							"followers_count": 41911172,
							"employees": map[string]interface{}{
								"range": "10001+",
							},
							"founded_date":   nil,
							"annual_revenue": "$100-1000B",
							"funding_amount": "25000000.0",
							"main_location": map[string]interface{}{
								"country": "united states",
								"state":   "california",
								"city":    "mountain view",
							},
						},
						"job": map[string]interface{}{
							"job_id":         "4430052243",
							"title":          "AI Driven Defense UTL",
							"url":            "https://nz.linkedin.com/jobs/view/ai-driven-defense-utl-at-google-4430052243",
							"location":       "New Zealand",
							"posted_at":      "2026-06-20T02:08:22+00:00",
							"posted_at_text": "7 hours ago",
						},
					},
				},
				"query":        map[string]interface{}{"name": "google"},
				"credit_count": 1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/psa":
			response := map[string]interface{}{
				"contacts": []map[string]interface{}{
					{
						"full_name": "John Doe",
						"current_job": map[string]interface{}{
							"title": "Software Engineer",
						},
						"company": map[string]interface{}{
							"name":     "TechCorp",
							"linkedin": "linkedin.com/company/techcorp",
							"website":  "https://techcorp.com",
							"industry": "software development",
							"main_location": map[string]interface{}{
								"country": "united states",
								"state":   "california",
								"city":    "san francisco",
							},
						},
						"location": map[string]interface{}{
							"country": "united states",
							"state":   "california",
							"city":    "san francisco",
						},
						"signal": map[string]interface{}{
							"name":       "employee_growth",
							"time_frame": 90,
							"bucket":     "high",
						},
					},
				},
				"query": map[string]interface{}{
					"signal_name": "employee_growth",
					"time_frame":  90,
					"bucket":      "high",
					"page":        1,
				},
				"credit_count": 1,
				"meta_data":    map[string]interface{}{"total_results": 1},
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/csa":
			response := map[string]interface{}{
				"companies": []map[string]interface{}{
					{
						"name":     "TechCorp",
						"website":  "https://techcorp.com",
						"domain":   "techcorp.com",
						"industry": "software development",
						"overview": "Enterprise software company",
						"type":     "private",
						"employees": map[string]interface{}{
							"range": "1001-5000",
						},
						"main_location": map[string]interface{}{
							"country": "united states",
							"state":   "california",
							"city":    "san francisco",
							"address": "123 Tech St",
						},
						"signal": map[string]interface{}{
							"name":       "employee_growth",
							"time_frame": 90,
							"bucket":     "high",
						},
					},
				},
				"query": map[string]interface{}{
					"signal_name": "employee_growth",
					"time_frame":  90,
					"bucket":      "high",
					"page":        1,
				},
				"credit_count": 1,
				"meta_data":    map[string]interface{}{"total_results": 1},
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/jca":
			response := map[string]interface{}{
				"job_changes": []map[string]interface{}{
					{
						"type":         "promotion",
						"linkedin_url": "https://linkedin.com/in/john-doe",
						"detected_at":  "2026-08-01T12:00:00Z",
						"from": map[string]interface{}{
							"company_linkedin_url": "https://linkedin.com/company/techcorp",
							"company_linkedin_id":  "12345",
							"company_name":         "TechCorp",
							"title":                "Software Engineer",
						},
						"to": map[string]interface{}{
							"company_linkedin_url": "https://linkedin.com/company/techcorp",
							"company_linkedin_id":  "12345",
							"company_name":         "TechCorp",
							"title":                "Senior Software Engineer",
						},
					},
				},
				"query": map[string]interface{}{
					"start_date": "2026-01-01",
					"end_date":   "2026-08-16",
					"type":       "promotion",
				},
				"credit_count": 1,
				"meta_data":    map[string]interface{}{"total_results": 1},
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/clf":
			response := map[string]interface{}{
				"profiles": []map[string]interface{}{
					{
						"full_name":    "Morteza Heydari",
						"linkedin_url": "https://linkedin.com/in/mortezaheydari1997",
						"job_title":    "Founder & CEO",
						"company_name": "CUFinder",
						"country":      "united states",
						"state":        "new york",
						"city":         "new york",
					},
				},
				"query":        "linkedin.com/in/mortezaheydari1997",
				"credit_count": 1,
				"meta_data":    map[string]interface{}{"total_results": 1},
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/nap":
			response := map[string]interface{}{
				"normalized_name": "Morteza Heydari",
				"query":           "morteza heydari",
				"credit_count":    1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/nau":
			response := map[string]interface{}{
				"normalized_url": "https://www.cufinder.io/about-us",
				"query":          "https://www.cufinder.io/about-us",
				"credit_count":   1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/gdc":
			response := map[string]interface{}{
				"offers_demo":  "yes",
				"query":        "https://www.stripe.com",
				"credit_count": 1,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		case "/cot":
			response := map[string]interface{}{
				"offers_free_trial": "yes",
				"query":             "https://www.stripe.com",
				"credit_count":      1,
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
		assert.Equal(t, "Software Engineer", result.Person.JobTitle)
	})

	t.Run("FCL Service", func(t *testing.T) {
		result, err := sdk.FCL("TechCorp")
		require.NoError(t, err)
		assert.Len(t, result.Companies, 2)
		assert.Equal(t, "DataCorp", result.Companies[0].Name)
		assert.Equal(t, "datacorp.com", result.Companies[0].Domain)
		assert.Equal(t, "SoftCorp", result.Companies[1].Name)
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
		assert.ElementsMatch(t, []string{"TechCorp Mobile", "TechCorp Cloud"}, result.Subsidiaries)
	})

	t.Run("FTS Service", func(t *testing.T) {
		result, err := sdk.FTS("TechCorp")
		require.NoError(t, err)
		assert.Equal(t, []string{"Go", "Python", "JavaScript"}, result.Technologies)
	})

	t.Run("EPP Service", func(t *testing.T) {
		result, err := sdk.EPP("https://linkedin.com/in/john-doe")
		require.NoError(t, err)
		assert.Equal(t, "John Doe", result.Person.FullName)
		assert.Equal(t, "Software Engineer", result.Person.JobTitle)
	})

	t.Run("FWE Service", func(t *testing.T) {
		result, err := sdk.FWE("https://linkedin.com/in/john-doe")
		require.NoError(t, err)
		assert.Equal(t, "john.doe@techcorp.com", result.WorkEmail)
	})

	t.Run("TEP Service", func(t *testing.T) {
		result, err := sdk.TEP("John Doe", "TechCorp")
		require.NoError(t, err)
		assert.Equal(t, "John Doe", result.Person.FullName)
		assert.Equal(t, "Software Engineer", result.Person.JobTitle)
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
	})

	t.Run("CLO Service", func(t *testing.T) {
		result, err := sdk.CLO("TechCorp")
		require.NoError(t, err)
		assert.Len(t, result.Locations, 2)
		assert.Equal(t, "US", result.Locations[0].Country)
		assert.Equal(t, "CA", result.Locations[0].State)
		assert.Equal(t, "San Francisco", result.Locations[0].City)
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
	})

	t.Run("PSE Service", func(t *testing.T) {
		result, err := sdk.PSE(PseParams{
			FullName:    "engineer",
			CompanyName: "TechCorp",
		})
		require.NoError(t, err)
		assert.Len(t, result.Peoples, 2)
		assert.Equal(t, "John Doe", result.Peoples[0].FullName)
		assert.Equal(t, "Software Engineer", result.Peoples[0].CurrentJob.Title)
		assert.Equal(t, "TechCorp", result.Peoples[0].Company.Name)
	})

	t.Run("LBS Service", func(t *testing.T) {
		result, err := sdk.LBS(LbsParams{
			Name: "coffee",
			City: "San Francisco",
		})
		require.NoError(t, err)
		assert.Len(t, result.Companies, 2)
		assert.Equal(t, "Coffee Shop", result.Companies[0].Name)
		assert.Equal(t, "123 Main St", result.Companies[0].Address)
		assert.Equal(t, "San Francisco", result.Companies[0].City)
	})

	t.Run("CEF Service", func(t *testing.T) {
		result, err := sdk.CEF("TechCorp", 1)
		require.NoError(t, err)
		assert.Len(t, result.Employees, 2)
		assert.Equal(t, "John Doe", result.Employees[0].FullName)
		assert.Equal(t, "John", result.Employees[0].FirstName)
		assert.Equal(t, "Doe", result.Employees[0].LastName)
		assert.Equal(t, "https://linkedin.com/in/john-doe", result.Employees[0].LinkedInURL)
		assert.Equal(t, "Software Engineer", result.Employees[0].JobTitle)
		assert.Equal(t, "TechCorp", result.Employees[0].CompanyName)
		assert.Equal(t, "Jane Smith", result.Employees[1].FullName)
		assert.Equal(t, "Product Manager", result.Employees[1].JobTitle)
		assert.Equal(t, "TechCorp", result.Query)
		assert.Equal(t, 1, result.CreditCount)
	})

	t.Run("NAC Service", func(t *testing.T) {
		result, err := sdk.NAC("cufinder inc.")
		require.NoError(t, err)
		assert.Equal(t, "Cufinder Inc.", result.Company)
		assert.Equal(t, "cufinder inc.", result.Query)
		assert.Equal(t, 1, result.CreditCount)
	})

	t.Run("CAA Service", func(t *testing.T) {
		result, err := sdk.CAA("TechCorp", 1)
		require.NoError(t, err)
		assert.Len(t, result.Activities, 1)
		assert.Equal(t, "https://www.linkedin.com/posts/cufinder_anthropic-alphabet-ai-activity-7462132400869888002-bCRi", result.Activities[0].ActivityURL)
		assert.Equal(t, "7462132400869888002", result.Activities[0].ActivityID)
		assert.Equal(t, 59, result.Activities[0].ActivityCommentsCount)
		assert.Equal(t, []string{"#AI", "#Anthropic", "#CUFinder"}, result.Activities[0].ActivityHashtags)
		assert.Equal(t, true, result.Activities[0].ActivityIsVideo)
		assert.Equal(t, "2026-05-18T13:30:04.063Z", result.Activities[0].ActivityPostedAt)
		assert.Equal(t, 3, result.Activities[0].ActivityReactionsCount)
		assert.Equal(t, "TechCorp", result.Query)
		assert.Equal(t, 1, result.CreditCount)
	})

	t.Run("CJA Service", func(t *testing.T) {
		result, err := sdk.CJA(CjaParams{
			Name: "google",
		})
		require.NoError(t, err)
		assert.Len(t, result.Jobs, 1)
		assert.Equal(t, "google", result.Jobs[0].Company.Name)
		assert.Equal(t, "software development", result.Jobs[0].Company.Industry)
		assert.Equal(t, "https://google.com", result.Jobs[0].Company.Website)
		assert.Equal(t, "linkedin.com/company/google", result.Jobs[0].Company.Linkedin)
		assert.Equal(t, 41911172, result.Jobs[0].Company.FollowersCount)
		assert.Equal(t, "10001+", result.Jobs[0].Company.Employees.Range)
		assert.Equal(t, "$100-1000B", result.Jobs[0].Company.AnnualRevenue)
		assert.Equal(t, "25000000.0", result.Jobs[0].Company.FundingAmount)
		assert.Equal(t, "united states", result.Jobs[0].Company.MainLocation.Country)
		assert.Equal(t, "california", result.Jobs[0].Company.MainLocation.State)
		assert.Equal(t, "mountain view", result.Jobs[0].Company.MainLocation.City)
		assert.Equal(t, "4430052243", result.Jobs[0].Job.JobID)
		assert.Equal(t, "AI Driven Defense UTL", result.Jobs[0].Job.Title)
		assert.Equal(t, "https://nz.linkedin.com/jobs/view/ai-driven-defense-utl-at-google-4430052243", result.Jobs[0].Job.URL)
		assert.Equal(t, "New Zealand", result.Jobs[0].Job.Location)
		assert.Equal(t, "2026-06-20T02:08:22+00:00", result.Jobs[0].Job.PostedAt)
		assert.Equal(t, "7 hours ago", result.Jobs[0].Job.PostedAtText)
		assert.Equal(t, 1, result.CreditCount)
	})

	t.Run("PSA Service", func(t *testing.T) {
		result, err := sdk.PSA(PsaParams{
			SignalName: "employee_growth",
			TimeFrame:  90,
			Bucket:     "high",
			Page:       1,
		})
		require.NoError(t, err)
		assert.Len(t, result.Contacts, 1)
		assert.Equal(t, "John Doe", result.Contacts[0].FullName)
		assert.Equal(t, "Software Engineer", result.Contacts[0].CurrentJob.Title)
		assert.Equal(t, "TechCorp", result.Contacts[0].Company.Name)
		assert.Equal(t, "software development", result.Contacts[0].Company.Industry)
		assert.Equal(t, "united states", result.Contacts[0].Company.MainLocation.Country)
		assert.Equal(t, "employee_growth", result.Contacts[0].Signal.Name)
		assert.Equal(t, 90, result.Contacts[0].Signal.TimeFrame)
		assert.Equal(t, "high", result.Contacts[0].Signal.Bucket)
		assert.Equal(t, 1, result.CreditCount)
	})

	t.Run("CSA Service", func(t *testing.T) {
		result, err := sdk.CSA(CsaParams{
			SignalName: "employee_growth",
			TimeFrame:  90,
			Bucket:     "high",
			Page:       1,
		})
		require.NoError(t, err)
		assert.Len(t, result.Companies, 1)
		assert.Equal(t, "TechCorp", result.Companies[0].Name)
		assert.Equal(t, "techcorp.com", result.Companies[0].Domain)
		assert.Equal(t, "software development", result.Companies[0].Industry)
		assert.Equal(t, "1001-5000", result.Companies[0].Employees.Range)
		assert.Equal(t, "san francisco", result.Companies[0].MainLocation.City)
		assert.Equal(t, "employee_growth", result.Companies[0].Signal.Name)
		assert.Equal(t, 1, result.CreditCount)
	})

	t.Run("JCA Service", func(t *testing.T) {
		result, err := sdk.JCA(JcaParams{
			StartDate: "2026-01-01",
			EndDate:   "2026-08-16",
			Type:      "promotion",
		})
		require.NoError(t, err)
		assert.Len(t, result.JobChanges, 1)
		assert.Equal(t, "promotion", result.JobChanges[0].Type)
		assert.Equal(t, "https://linkedin.com/in/john-doe", result.JobChanges[0].LinkedinURL)
		assert.Equal(t, "TechCorp", result.JobChanges[0].From.CompanyName)
		assert.Equal(t, "Software Engineer", result.JobChanges[0].From.Title)
		assert.Equal(t, "Senior Software Engineer", result.JobChanges[0].To.Title)
		assert.Equal(t, 1, result.CreditCount)
	})

	t.Run("CLF Service", func(t *testing.T) {
		result, err := sdk.CLF("linkedin.com/in/mortezaheydari1997")
		require.NoError(t, err)
		assert.Len(t, result.Profiles, 1)
		assert.Equal(t, "Morteza Heydari", result.Profiles[0].FullName)
		assert.Equal(t, "https://linkedin.com/in/mortezaheydari1997", result.Profiles[0].LinkedinURL)
		assert.Equal(t, "Founder & CEO", result.Profiles[0].JobTitle)
		assert.Equal(t, "CUFinder", result.Profiles[0].CompanyName)
		assert.Equal(t, "new york", result.Profiles[0].City)
		assert.Equal(t, 1, result.CreditCount)
	})

	t.Run("NAP Service", func(t *testing.T) {
		result, err := sdk.NAP("morteza heydari")
		require.NoError(t, err)
		assert.Equal(t, "Morteza Heydari", result.NormalizedName)
		assert.Equal(t, 1, result.CreditCount)
	})

	t.Run("NAU Service", func(t *testing.T) {
		result, err := sdk.NAU("https://www.cufinder.io/about-us")
		require.NoError(t, err)
		assert.Equal(t, "https://www.cufinder.io/about-us", result.NormalizedUrl)
		assert.Equal(t, 1, result.CreditCount)
	})

	t.Run("GDC Service", func(t *testing.T) {
		result, err := sdk.GDC("https://www.stripe.com")
		require.NoError(t, err)
		assert.Equal(t, "yes", result.OffersDemo)
		assert.Equal(t, 1, result.CreditCount)
	})

	t.Run("COT Service", func(t *testing.T) {
		result, err := sdk.COT("https://www.stripe.com")
		require.NoError(t, err)
		assert.Equal(t, "yes", result.OffersFreeTrial)
		assert.Equal(t, 1, result.CreditCount)
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

		_, err = sdk.PSA(PsaParams{TimeFrame: 90, Bucket: "high"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "signal_name is required")

		_, err = sdk.PSA(PsaParams{SignalName: "employee_growth", TimeFrame: 90})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "bucket is required")

		_, err = sdk.CSA(CsaParams{TimeFrame: 90, Bucket: "high"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "signal_name is required")

		_, err = sdk.CSA(CsaParams{SignalName: "employee_growth"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "bucket is required")

		_, err = sdk.JCA(JcaParams{EndDate: "2026-08-16"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "start_date is required")

		_, err = sdk.JCA(JcaParams{StartDate: "2026-01-01"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "end_date is required")

		_, err = sdk.CLF("")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "query is required")

		_, err = sdk.NAP("")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "person_name is required")

		_, err = sdk.NAU("")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "url is required")

		_, err = sdk.GDC("")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "url is required")

		_, err = sdk.COT("")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "url is required")
	})
}
