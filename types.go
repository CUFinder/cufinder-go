package cufinder

// BaseResponse represents the base response structure
type BaseResponse struct {
	Query           string                 `json:"query,omitempty"`
	CreditCount     int                    `json:"credit_count,omitempty"`
	MetaData        map[string]interface{} `json:"meta_data,omitempty"`
	ConfidenceLevel int                    `json:"confidence_level,omitempty"`
}

// Company represents company information
type Company struct {
	Name           string                 `json:"name,omitempty"`
	Domain         string                 `json:"domain,omitempty"`
	LinkedInURL    string                 `json:"linkedin_url,omitempty"`
	Industry       string                 `json:"industry,omitempty"`
	Size           string                 `json:"size,omitempty"`
	Location       string                 `json:"location,omitempty"`
	Description    string                 `json:"description,omitempty"`
	Founded        int                    `json:"founded,omitempty"`
	Revenue        string                 `json:"revenue,omitempty"`
	Employees      map[string]interface{} `json:"employees,omitempty"`
	Website        string                 `json:"website,omitempty"`
	Phone          string                 `json:"phone,omitempty"`
	Email          string                 `json:"email,omitempty"`
	SocialMedia    map[string]interface{} `json:"social_media,omitempty"`
	Technologies   []string               `json:"technologies,omitempty"`
	Subsidiaries   []string               `json:"subsidiaries,omitempty"`
	Headquarters   string                 `json:"headquarters,omitempty"`
	Country        string                 `json:"country,omitempty"`
	State          string                 `json:"state,omitempty"`
	City           string                 `json:"city,omitempty"`
	ZipCode        string                 `json:"zip_code,omitempty"`
	Address        string                 `json:"address,omitempty"`
}

// Person represents person information
type Person struct {
	FirstName     string                 `json:"first_name,omitempty"`
	LastName      string                 `json:"last_name,omitempty"`
	FullName      string                 `json:"full_name,omitempty"`
	Email         string                 `json:"email,omitempty"`
	Phone         string                 `json:"phone,omitempty"`
	LinkedInURL   string                 `json:"linkedin_url,omitempty"`
	JobTitle      string                 `json:"job_title,omitempty"`
	Company       string                 `json:"company,omitempty"`
	CompanyDomain string                 `json:"company_domain,omitempty"`
	Location      string                 `json:"location,omitempty"`
	Country       string                 `json:"country,omitempty"`
	State         string                 `json:"state,omitempty"`
	City          string                 `json:"city,omitempty"`
	Bio           string                 `json:"bio,omitempty"`
	Experience    []map[string]interface{} `json:"experience,omitempty"`
	Education     []map[string]interface{} `json:"education,omitempty"`
	Skills        []string               `json:"skills,omitempty"`
	Languages     []string               `json:"languages,omitempty"`
	SocialMedia   map[string]interface{} `json:"social_media,omitempty"`
}

// Response types for each service
type CufResponse struct {
	BaseResponse
	Domain string `json:"domain"`
}

type LcufResponse struct {
	BaseResponse
	LinkedInURL string `json:"linkedin_url"`
}

type DtcResponse struct {
	BaseResponse
	CompanyName string `json:"company_name"`
}

type DteResponse struct {
	BaseResponse
	Emails []string `json:"emails"`
}

type NtpResponse struct {
	BaseResponse
	Phones []string `json:"phones"`
}

type RelResponse struct {
	BaseResponse
	Person  Person  `json:"person"`
	Company Company `json:"company"`
}

type FclResponse struct {
	BaseResponse
	Lookalikes []Company `json:"lookalikes"`
}

type ElfResponse struct {
	BaseResponse
	Fundraising map[string]interface{} `json:"fundraising"`
}

type CarResponse struct {
	BaseResponse
	Revenue map[string]interface{} `json:"revenue"`
}

type FccResponse struct {
	BaseResponse
	Subsidiaries []Company `json:"subsidiaries"`
}

type FtsResponse struct {
	BaseResponse
	TechStack map[string]interface{} `json:"tech_stack"`
}

type EppResponse struct {
	BaseResponse
	Person  Person  `json:"person"`
	Company Company `json:"company"`
}

type FweResponse struct {
	BaseResponse
	Email string `json:"email"`
}

type TepResponse struct {
	BaseResponse
	Person Person `json:"person"`
}

type EncResponse struct {
	BaseResponse
	Company Company `json:"company"`
}

type CecResponse struct {
	BaseResponse
	Countries    []string `json:"countries"`
	TotalResults int      `json:"total_results"`
}

type CloResponse struct {
	BaseResponse
	Locations []map[string]interface{} `json:"locations"`
}

type CseResponse struct {
	BaseResponse
	Companies    []Company `json:"companies"`
	TotalResults int       `json:"total_results"`
	Page         int       `json:"page"`
}

type PseResponse struct {
	BaseResponse
	People       []Person `json:"people"`
	TotalResults int      `json:"total_results"`
	Page         int      `json:"page"`
}

type LbsResponse struct {
	BaseResponse
	Businesses   []Company `json:"businesses"`
	TotalResults int       `json:"total_results"`
	Page         int       `json:"page"`
}

// Parameter types for each service
type CufParams struct {
	CompanyName string `json:"company_name"`
	CountryCode string `json:"country_code"`
}

type LcufParams struct {
	CompanyName string `json:"company_name"`
}

type DtcParams struct {
	CompanyWebsite string `json:"company_website"`
}

type DteParams struct {
	CompanyWebsite string `json:"company_website"`
}

type NtpParams struct {
	CompanyName string `json:"company_name"`
}

type RelParams struct {
	Email string `json:"email"`
}

type FclParams struct {
	Query string `json:"query"`
}

type ElfParams struct {
	Query string `json:"query"`
}

type CarParams struct {
	Query string `json:"query"`
}

type FccParams struct {
	Query string `json:"query"`
}

type FtsParams struct {
	Query string `json:"query"`
}

type EppParams struct {
	LinkedInURL string `json:"linkedin_url"`
}

type FweParams struct {
	LinkedInURL string `json:"linkedin_url"`
}

type TepParams struct {
	FullName string `json:"full_name"`
	Company  string `json:"company"`
}

type EncParams struct {
	Query string `json:"query"`
}

type CecParams struct {
	Query string `json:"query"`
}

type CloParams struct {
	Query string `json:"query"`
}

type CseParams struct {
	Name                   string   `json:"name,omitempty"`
	Country                string   `json:"country,omitempty"`
	State                  string   `json:"state,omitempty"`
	City                   string   `json:"city,omitempty"`
	FollowersCountMin      int      `json:"followers_count_min,omitempty"`
	FollowersCountMax      int      `json:"followers_count_max,omitempty"`
	Industry               string   `json:"industry,omitempty"`
	EmployeeSize           string   `json:"employee_size,omitempty"`
	FoundedAfterYear       int      `json:"founded_after_year,omitempty"`
	FoundedBeforeYear      int      `json:"founded_before_year,omitempty"`
	FundingAmountMax       int      `json:"funding_amount_max,omitempty"`
	FundingAmountMin       int      `json:"funding_amount_min,omitempty"`
	ProductsServices       []string `json:"products_services,omitempty"`
	IsSchool               bool     `json:"is_school,omitempty"`
	AnnualRevenueMin       int      `json:"annual_revenue_min,omitempty"`
	AnnualRevenueMax       int      `json:"annual_revenue_max,omitempty"`
	Page                   int      `json:"page,omitempty"`
}

type PseParams struct {
	FullName                    string   `json:"full_name,omitempty"`
	Country                     string   `json:"country,omitempty"`
	State                       string   `json:"state,omitempty"`
	City                        string   `json:"city,omitempty"`
	JobTitleRole                string   `json:"job_title_role,omitempty"`
	JobTitleLevel               string   `json:"job_title_level,omitempty"`
	CompanyCountry              string   `json:"company_country,omitempty"`
	CompanyState                string   `json:"company_state,omitempty"`
	CompanyCity                 string   `json:"company_city,omitempty"`
	CompanyName                 string   `json:"company_name,omitempty"`
	CompanyLinkedInURL          string   `json:"company_linkedin_url,omitempty"`
	CompanyIndustry             string   `json:"company_industry,omitempty"`
	CompanyEmployeeSize         string   `json:"company_employee_size,omitempty"`
	CompanyProductsServices     []string `json:"company_products_services,omitempty"`
	CompanyAnnualRevenueMin     int      `json:"company_annual_revenue_min,omitempty"`
	CompanyAnnualRevenueMax     int      `json:"company_annual_revenue_max,omitempty"`
	Page                        int      `json:"page,omitempty"`
}

type LbsParams struct {
	Name      string `json:"name,omitempty"`
	Country   string `json:"country,omitempty"`
	State     string `json:"state,omitempty"`
	City      string `json:"city,omitempty"`
	Industry  string `json:"industry,omitempty"`
	Page      int    `json:"page,omitempty"`
}
