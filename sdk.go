package cufinder

import (
	"time"
)

// SDK represents the main CUFinder SDK
type SDK struct {
	client  *Client
	service *Service
}

// NewSDK creates a new CUFinder SDK instance
func NewSDK(apiKey string) *SDK {
	config := ClientConfig{
		APIKey:     apiKey,
		BaseURL:    "https://api.cufinder.io/v2",
		Timeout:    30 * time.Second,
		MaxRetries: 3,
	}

	client := NewClient(config)
	service := NewService(client)

	return &SDK{
		client:  client,
		service: service,
	}
}

// NewSDKWithConfig creates a new CUFinder SDK instance with custom configuration
func NewSDKWithConfig(config ClientConfig) *SDK {
	client := NewClient(config)
	service := NewService(client)

	return &SDK{
		client:  client,
		service: service,
	}
}

// Company Services

// CUF - Get company domain from company name
func (s *SDK) CUF(companyName, countryCode string) (*CufResponse, error) {
	return s.service.GetDomain(CufParams{
		CompanyName: companyName,
		CountryCode: countryCode,
	})
}

// LCUF - Get LinkedIn URL from company name
func (s *SDK) LCUF(companyName string) (*LcufResponse, error) {
	return s.service.GetLinkedInURL(LcufParams{
		CompanyName: companyName,
	})
}

// DTC - Get company name from domain
func (s *SDK) DTC(companyWebsite string) (*DtcResponse, error) {
	return s.service.GetCompanyName(DtcParams{
		CompanyWebsite: companyWebsite,
	})
}

// DTE - Get company emails from domain
func (s *SDK) DTE(companyWebsite string) (*DteResponse, error) {
	return s.service.GetEmails(DteParams{
		CompanyWebsite: companyWebsite,
	})
}

// NTP - Get company phones from company name
func (s *SDK) NTP(companyName string) (*NtpResponse, error) {
	return s.service.GetPhones(NtpParams{
		CompanyName: companyName,
	})
}

// Person Services

// EPP - Enrich LinkedIn profile
func (s *SDK) EPP(linkedInURL string) (*EppResponse, error) {
	return s.service.EnrichProfile(EppParams{
		LinkedInURL: linkedInURL,
	})
}

// REL - Reverse email lookup
func (s *SDK) REL(email string) (*RelResponse, error) {
	return s.service.ReverseEmailLookup(RelParams{
		Email: email,
	})
}

// FWE - Get email from profile
func (s *SDK) FWE(linkedInURL string) (*FweResponse, error) {
	return s.service.GetEmailFromProfile(FweParams{
		LinkedInURL: linkedInURL,
	})
}

// TEP - Enrich person information
func (s *SDK) TEP(fullName, company string) (*TepResponse, error) {
	return s.service.EnrichPerson(TepParams{
		FullName: fullName,
		Company:  company,
	})
}

// Company Intelligence Services

// FCL - Get company lookalikes
func (s *SDK) FCL(query string) (*FclResponse, error) {
	return s.service.GetLookalikes(FclParams{
		Query: query,
	})
}

// ELF - Get company fundraising information
func (s *SDK) ELF(query string) (*ElfResponse, error) {
	return s.service.GetFundraising(ElfParams{
		Query: query,
	})
}

// CAR - Get company revenue
func (s *SDK) CAR(query string) (*CarResponse, error) {
	return s.service.GetRevenue(CarParams{
		Query: query,
	})
}

// FCC - Get company subsidiaries
func (s *SDK) FCC(query string) (*FccResponse, error) {
	return s.service.GetSubsidiaries(FccParams{
		Query: query,
	})
}

// FTS - Get company tech stack
func (s *SDK) FTS(query string) (*FtsResponse, error) {
	return s.service.GetTechStack(FtsParams{
		Query: query,
	})
}

// ENC - Enrich company information
func (s *SDK) ENC(query string) (*EncResponse, error) {
	return s.service.EnrichCompany(EncParams{
		Query: query,
	})
}

// CEC - Get company employee countries
func (s *SDK) CEC(query string) (*CecResponse, error) {
	return s.service.GetEmployeeCountries(CecParams{
		Query: query,
	})
}

// CLO - Get company locations
func (s *SDK) CLO(query string) (*CloResponse, error) {
	return s.service.GetLocations(CloParams{
		Query: query,
	})
}

// Search Services

// CSE - Search companies
func (s *SDK) CSE(params CseParams) (*CseResponse, error) {
	return s.service.SearchCompanies(params)
}

// PSE - Search people
func (s *SDK) PSE(params PseParams) (*PseResponse, error) {
	return s.service.SearchPeople(params)
}

// LBS - Search local businesses
func (s *SDK) LBS(params LbsParams) (*LbsResponse, error) {
	return s.service.SearchLocalBusinesses(params)
}

// BCD - B2B Customers Finder
func (s *SDK) BCD(url string) (*BcdResponse, error) {
	return s.service.ExtractB2BCustomers(BcdParams{
		Url: url,
	})
}

// CCP - Company Career Page Finder
func (s *SDK) CCP(url string) (*CcpResponse, error) {
	return s.service.FindCareersPage(CcpParams{
		Url: url,
	})
}

// GetClient returns the underlying HTTP client for advanced usage
func (s *SDK) GetClient() *Client {
	return s.client
}
