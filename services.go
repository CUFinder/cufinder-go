package cufinder

import (
	"encoding/json"
	"fmt"
)

// Service represents a base service
type Service struct {
	client *Client
}

// NewService creates a new service instance
func NewService(client *Client) *Service {
	return &Service{client: client}
}

// CUF Service - Company URL Finder
func (s *Service) GetDomain(params CufParams) (*CufResponse, error) {
	if params.CompanyName == "" {
		return nil, fmt.Errorf("company_name is required")
	}
	if params.CountryCode == "" {
		return nil, fmt.Errorf("country_code is required")
	}

	response, err := s.client.Post("/cuf", params)
	if err != nil {
		return nil, fmt.Errorf("CUF service error: %w", err)
	}

	var result CufResponse
	if err := mapToStruct(response, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// LCUF Service - LinkedIn Company URL Finder
func (s *Service) GetLinkedInURL(params LcufParams) (*LcufResponse, error) {
	if params.CompanyName == "" {
		return nil, fmt.Errorf("company_name is required")
	}

	response, err := s.client.Post("/lcuf", params)
	if err != nil {
		return nil, fmt.Errorf("LCUF service error: %w", err)
	}

	var result LcufResponse
	if err := mapToStruct(response, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// DTC Service - Domain to Company
func (s *Service) GetCompanyName(params DtcParams) (*DtcResponse, error) {
	if params.CompanyWebsite == "" {
		return nil, fmt.Errorf("company_website is required")
	}

	response, err := s.client.Post("/dtc", params)
	if err != nil {
		return nil, fmt.Errorf("DTC service error: %w", err)
	}

	var result DtcResponse
	if err := mapToStruct(response, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// DTE Service - Domain to Emails
func (s *Service) GetEmails(params DteParams) (*DteResponse, error) {
	if params.CompanyWebsite == "" {
		return nil, fmt.Errorf("company_website is required")
	}

	response, err := s.client.Post("/dte", params)
	if err != nil {
		return nil, fmt.Errorf("DTE service error: %w", err)
	}

	var result DteResponse
	if err := mapToStruct(response, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// NTP Service - Name to Phones
func (s *Service) GetPhones(params NtpParams) (*NtpResponse, error) {
	if params.CompanyName == "" {
		return nil, fmt.Errorf("company_name is required")
	}

	response, err := s.client.Post("/ntp", params)
	if err != nil {
		return nil, fmt.Errorf("NTP service error: %w", err)
	}

	var result NtpResponse
	if err := mapToStruct(response, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// REL Service - Reverse Email Lookup
func (s *Service) ReverseEmailLookup(params RelParams) (*RelResponse, error) {
	if params.Email == "" {
		return nil, fmt.Errorf("email is required")
	}

	response, err := s.client.Post("/rel", params)
	if err != nil {
		return nil, fmt.Errorf("REL service error: %w", err)
	}

	var result RelResponse
	if err := mapToStruct(response, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// FCL Service - Find Company Lookalikes
func (s *Service) GetLookalikes(params FclParams) (*FclResponse, error) {
	if params.Query == "" {
		return nil, fmt.Errorf("query is required")
	}

	response, err := s.client.Post("/fcl", params)
	if err != nil {
		return nil, fmt.Errorf("FCL service error: %w", err)
	}

	var result FclResponse
	if err := mapToStruct(response, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// ELF Service - Enrich LinkedIn Fundraising
func (s *Service) GetFundraising(params ElfParams) (*ElfResponse, error) {
	if params.Query == "" {
		return nil, fmt.Errorf("query is required")
	}

	response, err := s.client.Post("/elf", params)
	if err != nil {
		return nil, fmt.Errorf("ELF service error: %w", err)
	}

	var result ElfResponse
	if err := mapToStruct(response, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// CAR Service - Company Annual Revenue
func (s *Service) GetRevenue(params CarParams) (*CarResponse, error) {
	if params.Query == "" {
		return nil, fmt.Errorf("query is required")
	}

	response, err := s.client.Post("/car", params)
	if err != nil {
		return nil, fmt.Errorf("CAR service error: %w", err)
	}

	var result CarResponse
	if err := mapToStruct(response, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// FCC Service - Find Company Children
func (s *Service) GetSubsidiaries(params FccParams) (*FccResponse, error) {
	if params.Query == "" {
		return nil, fmt.Errorf("query is required")
	}

	response, err := s.client.Post("/fcc", params)
	if err != nil {
		return nil, fmt.Errorf("FCC service error: %w", err)
	}

	var result FccResponse
	if err := mapToStruct(response, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// FTS Service - Find Tech Stack
func (s *Service) GetTechStack(params FtsParams) (*FtsResponse, error) {
	if params.Query == "" {
		return nil, fmt.Errorf("query is required")
	}

	response, err := s.client.Post("/fts", params)
	if err != nil {
		return nil, fmt.Errorf("FTS service error: %w", err)
	}

	var result FtsResponse
	if err := mapToStruct(response, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// EPP Service - Enrich Profile
func (s *Service) EnrichProfile(params EppParams) (*EppResponse, error) {
	if params.LinkedInURL == "" {
		return nil, fmt.Errorf("linkedin_url is required")
	}

	response, err := s.client.Post("/epp", params)
	if err != nil {
		return nil, fmt.Errorf("EPP service error: %w", err)
	}

	var result EppResponse
	if err := mapToStruct(response, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// FWE Service - Find Work Email
func (s *Service) GetEmailFromProfile(params FweParams) (*FweResponse, error) {
	if params.LinkedInURL == "" {
		return nil, fmt.Errorf("linkedin_url is required")
	}

	response, err := s.client.Post("/fwe", params)
	if err != nil {
		return nil, fmt.Errorf("FWE service error: %w", err)
	}

	var result FweResponse
	if err := mapToStruct(response, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// TEP Service - Person Enrichment
func (s *Service) EnrichPerson(params TepParams) (*TepResponse, error) {
	if params.FullName == "" {
		return nil, fmt.Errorf("full_name is required")
	}
	if params.Company == "" {
		return nil, fmt.Errorf("company is required")
	}

	response, err := s.client.Post("/tep", params)
	if err != nil {
		return nil, fmt.Errorf("TEP service error: %w", err)
	}

	var result TepResponse
	if err := mapToStruct(response, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// ENC Service - Company Enrichment
func (s *Service) EnrichCompany(params EncParams) (*EncResponse, error) {
	if params.Query == "" {
		return nil, fmt.Errorf("query is required")
	}

	response, err := s.client.Post("/enc", params)
	if err != nil {
		return nil, fmt.Errorf("ENC service error: %w", err)
	}

	var result EncResponse
	if err := mapToStruct(response, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// CEC Service - Company Employee Countries
func (s *Service) GetEmployeeCountries(params CecParams) (*CecResponse, error) {
	if params.Query == "" {
		return nil, fmt.Errorf("query is required")
	}

	response, err := s.client.Post("/cec", params)
	if err != nil {
		return nil, fmt.Errorf("CEC service error: %w", err)
	}

	var result CecResponse
	if err := mapToStruct(response, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// CLO Service - Company Locations
func (s *Service) GetLocations(params CloParams) (*CloResponse, error) {
	if params.Query == "" {
		return nil, fmt.Errorf("query is required")
	}

	response, err := s.client.Post("/clo", params)
	if err != nil {
		return nil, fmt.Errorf("CLO service error: %w", err)
	}

	var result CloResponse
	if err := mapToStruct(response, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// CSE Service - Company Search
func (s *Service) SearchCompanies(params CseParams) (*CseResponse, error) {
	response, err := s.client.Post("/cse", params)
	if err != nil {
		return nil, fmt.Errorf("CSE service error: %w", err)
	}

	var result CseResponse
	if err := mapToStruct(response, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// PSE Service - Person Search
func (s *Service) SearchPeople(params PseParams) (*PseResponse, error) {
	response, err := s.client.Post("/pse", params)
	if err != nil {
		return nil, fmt.Errorf("PSE service error: %w", err)
	}

	var result PseResponse
	if err := mapToStruct(response, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// LBS Service - Local Business Search
func (s *Service) SearchLocalBusinesses(params LbsParams) (*LbsResponse, error) {
	response, err := s.client.Post("/lbs", params)
	if err != nil {
		return nil, fmt.Errorf("LBS service error: %w", err)
	}

	var result LbsResponse
	if err := mapToStruct(response, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// BCD Service - B2B Customers Finder
func (s *Service) ExtractB2BCustomers(params BcdParams) (*BcdResponse, error) {
	if params.Url == "" {
		return nil, fmt.Errorf("url is required")
	}

	response, err := s.client.Post("/bcd", params)
	if err != nil {
		return nil, fmt.Errorf("BCD service error: %w", err)
	}

	var result BcdResponse
	if err := mapToStruct(response, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// CCP Service - Company Career Page Finder
func (s *Service) FindCareersPage(params CcpParams) (*CcpResponse, error) {
	if params.Url == "" {
		return nil, fmt.Errorf("url is required")
	}

	response, err := s.client.Post("/ccp", params)
	if err != nil {
		return nil, fmt.Errorf("CCP service error: %w", err)
	}

	var result CcpResponse
	if err := mapToStruct(response, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// Helper function to convert map to struct
func mapToStruct(data map[string]interface{}, result interface{}) error {
	// Check if the response has a "data" wrapper (like Python SDK)
	if dataWrapper, exists := data["data"]; exists {
		// Extract the actual data from the wrapper
		if dataMap, ok := dataWrapper.(map[string]interface{}); ok {
			// Add meta_data if it exists in the outer response
			if metaData, metaExists := data["meta_data"]; metaExists {
				dataMap["meta_data"] = metaData
			}
			data = dataMap
		}
	}

	// Convert map to JSON bytes
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Unmarshal JSON bytes to struct
	return json.Unmarshal(jsonData, result)
}
