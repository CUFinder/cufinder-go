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

// ISC - Company Saas Checker
func (s *SDK) ISC(url string) (*IscResponse, error) {
	return s.service.IsSaas(IscParams{
		Url: url,
	})
}

// CBC - Company B2B or B2C Checker
func (s *SDK) CBC(url string) (*CbcResponse, error) {
	return s.service.GetCompanyBusinessType(CbcParams{
		Url: url,
	})
}

// CSC - Company Mission Statement
func (s *SDK) CSC(url string) (*CscResponse, error) {
	return s.service.GetCompanyMissionStatement(CscParams{
		Url: url,
	})
}

// CSN - Company Snapshot
func (s *SDK) CSN(url string) (*CsnResponse, error) {
	return s.service.GetCompanySnapshot(CsnParams{
		Url: url,
	})
}

// NAO - Phone Number Normalizer
func (s *SDK) NAO(phone string) (*NaoResponse, error) {
	return s.service.NormalizePhone(NaoParams{
		Phone: phone,
	})
}

// NAA - Address Normalizer
func (s *SDK) NAA(address string) (*NaaResponse, error) {
	return s.service.NormalizeAddress(NaaParams{
		Address: address,
	})
}

// CEF - Company Employee Finder
func (s *SDK) CEF(query string, page int) (*CefResponse, error) {
	return s.service.FindCompanyEmployees(CefParams{
		Query: query,
		Page:  page,
	})
}

// NAC - Company Name Normalizer
func (s *SDK) NAC(company string) (*NacResponse, error) {
	return s.service.NormalizeCompanyName(NacParams{
		Company: company,
	})
}

// CAA - Company Activity API
func (s *SDK) CAA(query string, page int) (*CaaResponse, error) {
	return s.service.GetCompanyActivities(CaaParams{
		Query: query,
		Page:  page,
	})
}

// CJA - Company Jobs API
func (s *SDK) CJA(params CjaParams) (*CjaResponse, error) {
	return s.service.SearchCompanyJobs(params)
}

// PSA - Contact Signals API
func (s *SDK) PSA(params PsaParams) (*PsaResponse, error) {
	return s.service.GetContactSignals(params)
}

// CSA - Company Signals API
func (s *SDK) CSA(params CsaParams) (*CsaResponse, error) {
	return s.service.GetCompanySignals(params)
}

// JCA - Job Changes API
func (s *SDK) JCA(params JcaParams) (*JcaResponse, error) {
	return s.service.GetJobChanges(params)
}

// CLF - Contact Lookalikes API
func (s *SDK) CLF(query string) (*ClfResponse, error) {
	return s.service.FindContactLookalikes(ClfParams{
		Query: query,
	})
}

// NAP - Normalize Person Name
func (s *SDK) NAP(personName string) (*NapResponse, error) {
	return s.service.NormalizePersonName(NapParams{
		PersonName: personName,
	})
}

// NAU - Normalize URL
func (s *SDK) NAU(url string) (*NauResponse, error) {
	return s.service.NormalizeUrl(NauParams{
		Url: url,
	})
}

// GDC - Company Gives Demo Checker
func (s *SDK) GDC(url string) (*GdcResponse, error) {
	return s.service.GivesDemo(GdcParams{
		Url: url,
	})
}

// COT - Company Offers Free Trial Checker
func (s *SDK) COT(url string) (*CotResponse, error) {
	return s.service.OffersFreeTrial(CotParams{
		Url: url,
	})
}

// GetClient returns the underlying HTTP client for advanced usage
func (s *SDK) GetClient() *Client {
	return s.client
}
