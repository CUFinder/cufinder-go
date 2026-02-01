package cufinder

// BaseResponse represents the base response structure
type BaseResponse struct {
	Query           interface{}            `json:"query,omitempty"`
	CreditCount     int                    `json:"credit_count,omitempty"`
	MetaData        map[string]interface{} `json:"meta_data,omitempty"`
	ConfidenceLevel int                    `json:"confidence_level,omitempty"`
}

type MainLocation struct {
	Geo        string `json:"geo,omitempty"`
	Country    string `json:"country,omitempty"`
	State      string `json:"state,omitempty"`
	City       string `json:"city,omitempty"`
	Address    string `json:"address,omitempty"`
	Continent  string `json:"continent,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
}

type CompanySocial struct {
	Facebook  string `json:"facebook,omitempty"`
	LinkedIn  string `json:"linkedin,omitempty"`
	Twitter   string `json:"twitter,omitempty"`
	Youtube   string `json:"youtube,omitempty"`
	Instagram string `json:"instagram,omitempty"`
}

type CompanyEmployees struct {
	Range string `json:"range,omitempty"`
	Count int    `json:"count,omitempty"`
}

// Company represents company information
type Company struct {
	Name         string           `json:"name,omitempty"`
	Domain       string           `json:"domain,omitempty"`
	LinkedInURL  string           `json:"linkedin_url,omitempty"`
	Industry     string           `json:"industry,omitempty"`
	Overview     string           `json:"overview,omitempty"`
	Type         string           `json:"type,omitempty"`
	Size         string           `json:"size,omitempty"`
	MainLocation MainLocation     `json:"main_location,omitempty"`
	Location     string           `json:"location,omitempty"`
	Description  string           `json:"description,omitempty"`
	Founded      int              `json:"founded,omitempty"`
	Revenue      string           `json:"revenue,omitempty"`
	Employees    CompanyEmployees `json:"employees,omitempty"`
	Website      string           `json:"website,omitempty"`
	Phone        string           `json:"phone,omitempty"`
	Email        string           `json:"email,omitempty"`
	Social       CompanySocial    `json:"social,omitempty"`
	Technologies []string         `json:"technologies,omitempty"`
	Subsidiaries []string         `json:"subsidiaries,omitempty"`
	Headquarters string           `json:"headquarters,omitempty"`
	Country      string           `json:"country,omitempty"`
	State        string           `json:"state,omitempty"`
	City         string           `json:"city,omitempty"`
	ZipCode      string           `json:"zip_code,omitempty"`
	Address      string           `json:"address,omitempty"`
}

type JobTitleCategory struct {
	Category      string `json:"category,omitempty"`
	SuperCategory string `json:"super_category,omitempty"`
}

type PeopleCurrentJob struct {
	Title      string             `json:"title,omitempty"`
	Role       string             `json:"role,omitempty"`
	Level      string             `json:"level,omitempty"`
	Categories []JobTitleCategory `json:"categories,omitempty"`
}

type PeopleConnections struct {
	HasWorkEmail     bool   `json:"has_work_email,omitempty"`
	HasPersonalEmail bool   `json:"has_personal_email,omitempty"`
	HasPhone         bool   `json:"has_phone,omitempty"`
	WorkEmail        string `json:"work_email,omitempty"`
	PersonalEmail    string `json:"personal_email,omitempty"`
	Phone            string `json:"phone,omitempty"`
	IsAcceptAll      bool   `json:"is_accept_all,omitempty"`
	IsAcceptEmail    bool   `json:"is_accept_email,omitempty"`
}

type PeopleEducationSchoolLocation struct {
	Name      string `json:"name,omitempty"`
	Locality  string `json:"locality,omitempty"`
	Region    string `json:"region,omitempty"`
	Country   string `json:"country,omitempty"`
	Continent string `json:"continent,omitempty"`
}

type PeopleEducationSchool struct {
	Name              string                        `json:"name,omitempty"`
	Type              string                        `json:"type,omitempty"`
	Id                string                        `json:"id,omitempty"`
	Location          PeopleEducationSchoolLocation `json:"location,omitempty"`
	LinkedinURL       string                        `json:"linkedin_url,omitempty"`
	FacebookURL       string                        `json:"facebook_url,omitempty"`
	TwitterURL        string                        `json:"twitter_url,omitempty"`
	LinkedinID        string                        `json:"linkedin_id,omitempty"`
	Website           string                        `json:"website,omitempty"`
	Domain            string                        `json:"domain,omitempty"`
	JobCompanyIDMongo string                        `json:"job_company_id_mongo,omitempty"`
	UniversityIDMongo string                        `json:"university_id_mongo,omitempty"`
}

type PeopleEducation struct {
	School    PeopleEducationSchool `json:"school,omitempty"`
	EndDate   string                `json:"end_date,omitempty"`
	StartDate string                `json:"start_date,omitempty"`
	Gpa       string                `json:"gpa,omitempty"`
	Degrees   []string              `json:"degrees,omitempty"`
	Majors    []string              `json:"majors,omitempty"`
	Minors    []string              `json:"minors,omitempty"`
	Summary   string                `json:"summary,omitempty"`
}

type PeopleExperienceCompany struct {
	Name        string `json:"name,omitempty"`
	LinkedInURL string `json:"linkedin_url,omitempty"`
	Domain      string `json:"domain,omitempty"`
	Size        string `json:"size,omitempty"`
	Industry    string `json:"industry,omitempty"`
}

type PeopleExperienceTitle struct {
	Name    string   `json:"name,omitempty"`
	Role    string   `json:"role,omitempty"`
	SubRole string   `json:"sub_role,omitempty"`
	Levels  []string `json:"levels,omitempty"`
}

type PeopleExperience struct {
	Company       PeopleExperienceCompany
	LocationNames []string
	EndDate       string
	StartDate     string
	Title         PeopleExperienceTitle
	IsPrimary     bool
	Summary       string
}

type PeopleCertification struct {
	Certification  string `json:"certification,omitempty"`
	Issuer         string `json:"issuer,omitempty"`
	IssueDate      string `json:"issue_date,omitempty"`
	ExpirationDate string `json:"expiration_date,omitempty"`
}

type PeopleLocation struct {
	Country string `json:"country,omitempty"`
	State   string `json:"state,omitempty"`
	City    string `json:"city,omitempty"`
}

type PeopleSocial struct {
	LinkedinUsername    string `json:"linkedin_username,omitempty"`
	LinkedinConnections int    `json:"linkedin_connections,omitempty"`
	LinkedIn            string `json:"linkedin,omitempty"`
	Twitter             string `json:"twitter,omitempty"`
	Facebook            string `json:"facebook,omitempty"`
	Github              string `json:"github,omitempty"`
}

// Person represents person information
type Person struct {
	FirstName      string                `json:"first_name,omitempty"`
	LastName       string                `json:"last_name,omitempty"`
	FullName       string                `json:"full_name,omitempty"`
	Logo           string                `json:"logo,omitempty"`
	Overview       string                `json:"overview,omitempty"`
	Experience     any                   `json:"experience,omitempty"` // FIXME: should be fixed from the source
	Connections    PeopleConnections     `json:"connections,omitempty"`
	Interests      []string              `json:"interests,omitempty"`
	Skills         []string              `json:"skills,omitempty"`
	Educations     []PeopleEducation     `json:"educations,omitempty"`
	Experiences    []PeopleExperience    `json:"experiences,omitempty"`
	Certifications []PeopleCertification `json:"certifications,omitempty"`
	Company        Company               `json:"company,omitempty"`
	Location       PeopleLocation        `json:"location,omitempty"`
	CurrentJob     PeopleCurrentJob      `json:"current_job,omitempty"`
	Social         PeopleSocial          `json:"social,omitempty"`
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

type RelPerson struct {
	FullName           string   `json:"full_name,omitempty"`
	FirstName          string   `json:"first_name,omitempty"`
	LastName           string   `json:"last_name,omitempty"`
	LinkedInURL        string   `json:"linkedin_url,omitempty"`
	Summary            string   `json:"summary,omitempty"`
	LinkedInFollowers  string   `json:"linkedin_followers,omitempty"`
	Facebook           string   `json:"facebook,omitempty"`
	Twitter            string   `json:"twitter,omitempty"`
	Avatar             string   `json:"avatar,omitempty"`
	Country            string   `json:"country,omitempty"`
	State              string   `json:"state,omitempty"`
	City               string   `json:"city,omitempty"`
	JobTitle           string   `json:"job_title,omitempty"`
	JobTitleCategories []string `json:"job_title_categories,omitempty"`
	CompanyName        string   `json:"company_name,omitempty"`
	CompanyLinkedIn    string   `json:"company_linkedin,omitempty"`
	CompanyWebsite     string   `json:"company_website,omitempty"`
	CompanySize        string   `json:"company_size,omitempty"`
	CompanyIndustry    string   `json:"company_industry,omitempty"`
	CompanyFaceBook    string   `json:"company_facebook,omitempty"`
	CompanyTwitter     string   `json:"company_twitter,omitempty"`
	CompanyCountry     string   `json:"company_country,omitempty"`
	CompanyState       string   `json:"company_state,omitempty"`
	CompanyCity        string   `json:"company_city,omitempty"`
}

type RelResponse struct {
	BaseResponse
	Person RelPerson `json:"person"`
}

type FclCompany struct {
	Name           string `json:"name,omitempty"`
	Website        string `json:"website,omitempty"`
	EmployeeCount  int    `json:"employee_count,omitempty"`
	Size           string `json:"size,omitempty"`
	Industry       string `json:"industry,omitempty"`
	Description    string `json:"description,omitempty"`
	LinkedinUrl    string `json:"linkedin_url,omitempty"`
	Domain         string `json:"domain,omitempty"`
	Country        string `json:"country,omitempty"`
	State          string `json:"state,omitempty"`
	City           string `json:"city,omitempty"`
	Address        string `json:"address,omitempty"`
	FoundedYear    string `json:"founded_year,omitempty"`
	LogoUrl        string `json:"logo_url,omitempty"`
	FollowersCount int    `json:"followers_count,omitempty"`
}

type FclResponse struct {
	BaseResponse
	Companies []FclCompany `json:"companies"`
}

type ElfFundraising struct {
	FundingLastRoundType         string `json:"funding_last_round_type,omitempty"`
	FundingAmmountCurrencyCode   string `json:"funding_ammount_currency_code,omitempty"`
	FundingMoneyRaised           string `json:"funding_money_raised,omitempty"`
	FundingLastRoundInvestorsUrl string `json:"funding_last_round_investors_url,omitempty"`
}

type ElfResponse struct {
	BaseResponse
	Fundraising ElfFundraising `json:"fundraising_info"`
}

type CarResponse struct {
	BaseResponse
	Revenue string `json:"annual_revenue"`
}

type FccResponse struct {
	BaseResponse
	Subsidiaries []string `json:"subsidiaries"`
}

type FtsResponse struct {
	BaseResponse
	Technologies []string `json:"technologies"`
}

type EppPerson struct {
	FullName           string   `json:"full_name,omitempty"`
	FirstName          string   `json:"first_name,omitempty"`
	LastName           string   `json:"last_name,omitempty"`
	LinkedInURL        string   `json:"linkedin_url,omitempty"`
	Summary            string   `json:"summary,omitempty"`
	LinkedInFollowers  int      `json:"linkedin_followers,omitempty"`
	Facebook           string   `json:"facebook,omitempty"`
	Twitter            string   `json:"twitter,omitempty"`
	Avatar             string   `json:"avatar,omitempty"`
	Country            string   `json:"country,omitempty"`
	State              string   `json:"state,omitempty"`
	City               string   `json:"city,omitempty"`
	JobTitle           string   `json:"job_title,omitempty"`
	JobTitleCategories []string `json:"job_title_categories,omitempty"`
	CompanyName        string   `json:"company_name,omitempty"`
	CompanyLinkedIn    string   `json:"company_linkedin,omitempty"`
	CompanyWebsite     string   `json:"company_website,omitempty"`
	CompanySize        string   `json:"company_size,omitempty"`
	CompanyIndustry    string   `json:"company_industry,omitempty"`
	CompanyFacebook    string   `json:"company_facebook,omitempty"`
	CompanyTwitter     string   `json:"company_twitter,omitempty"`
	CompanyCountry     string   `json:"company_country,omitempty"`
	CompanyState       string   `json:"company_state,omitempty"`
	CompanyCity        string   `json:"company_city,omitempty"`
}

type EppResponse struct {
	BaseResponse
	Person EppPerson `json:"person"`
}

type FweResponse struct {
	BaseResponse
	WorkEmail string `json:"work_email"`
}

type TepPerson struct {
	FullName           string   `json:"full_name,omitempty"`
	FirstName          string   `json:"first_name,omitempty"`
	LastName           string   `json:"last_name,omitempty"`
	LinkedInURL        string   `json:"linkedin_url,omitempty"`
	Summary            string   `json:"summary,omitempty"`
	LinkedInFollowers  int      `json:"linkedin_followers,omitempty"`
	Facebook           string   `json:"facebook,omitempty"`
	Twitter            string   `json:"twitter,omitempty"`
	Avatar             string   `json:"avatar,omitempty"`
	Country            string   `json:"country,omitempty"`
	State              string   `json:"state,omitempty"`
	City               string   `json:"city,omitempty"`
	JobTitle           string   `json:"job_title,omitempty"`
	JobTitleCategories []string `json:"job_title_categories,omitempty"`
	CompanyName        string   `json:"company_name,omitempty"`
	CompanyLinkedIn    string   `json:"company_linkedin,omitempty"`
	CompanyWebsite     string   `json:"company_website,omitempty"`
	CompanySize        string   `json:"company_size,omitempty"`
	CompanyIndustry    string   `json:"company_industry,omitempty"`
	CompanyFacebook    string   `json:"company_facebook,omitempty"`
	CompanyTwitter     string   `json:"company_twitter,omitempty"`
	CompanyCountry     string   `json:"company_country,omitempty"`
	CompanyState       string   `json:"company_state,omitempty"`
	CompanyCity        string   `json:"company_city,omitempty"`
	Email              string   `json:"email,omitempty"`
	Phone              string   `json:"phone,omitempty"`
}

type TepResponse struct {
	BaseResponse
	Person TepPerson `json:"person"`
}

type EncCompany struct {
	Name           string `json:"name,omitempty"`
	Website        string `json:"website,omitempty"`
	EmployeeCount  int    `json:"employee_count,omitempty"`
	Industry       string `json:"size,omitempty"`
	Size           string `json:"industry,omitempty"`
	Description    string `json:"description,omitempty"`
	LinkedInURL    string `json:"linkedin_url,omitempty"`
	Type           string `json:"type,omitempty"`
	Domain         string `json:"domain,omitempty"`
	Country        string `json:"country,omitempty"`
	State          string `json:"state,omitempty"`
	City           string `json:"city,omitempty"`
	Address        string `json:"address,omitempty"`
	FoundedYear    string `json:"founded_year,omitempty"`
	LogoUrl        string `json:"logo_url,omitempty"`
	FollowersCount int    `json:"followers_count,omitempty"`
}

type EncResponse struct {
	BaseResponse
	Company EncCompany `json:"company"`
}

type CecResponse struct {
	BaseResponse
	Countries interface{} `json:"countries"`
}

type CloLocation struct {
	Country    string `json:"country,omitempty"`
	State      string `json:"state,omitempty"`
	City       string `json:"city,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
	Line1      string `json:"line1,omitempty"`
	Line2      string `json:"line2,omitempty"`
	Latitude   string `json:"latitude,omitempty"`
	Longitude  string `json:"longitude,omitempty"`
}

type CloResponse struct {
	BaseResponse
	Locations []CloLocation `json:"locations"`
}

type CseResponse struct {
	BaseResponse
	Companies []Company `json:"companies"`
}

type PseResponse struct {
	BaseResponse
	Peoples []Person `json:"peoples"`
}

type LbsResponse struct {
	BaseResponse
	Companies []Company `json:"companies"`
}

type BcdResponse struct {
	BaseResponse
	Customers []string `json:"customers"`
}

type CcpResponse struct {
	BaseResponse
	CareersPageUrl string `json:"careers_page_url"`
}

type IscResponse struct {
	BaseResponse
	IsSaas string `json:"is_saas"`
}

type CbcResponse struct {
	BaseResponse
	BusinessType string `json:"business_type"`
}

type CscResponse struct {
	BaseResponse
	MissionStatement string `json:"mission_statement"`
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
	Name              string   `json:"name,omitempty"`
	Country           string   `json:"country,omitempty"`
	State             string   `json:"state,omitempty"`
	City              string   `json:"city,omitempty"`
	FollowersCountMin int      `json:"followers_count_min,omitempty"`
	FollowersCountMax int      `json:"followers_count_max,omitempty"`
	Industry          string   `json:"industry,omitempty"`
	EmployeeSize      string   `json:"employee_size,omitempty"`
	FoundedAfterYear  int      `json:"founded_after_year,omitempty"`
	FoundedBeforeYear int      `json:"founded_before_year,omitempty"`
	FundingAmountMax  int      `json:"funding_amount_max,omitempty"`
	FundingAmountMin  int      `json:"funding_amount_min,omitempty"`
	ProductsServices  []string `json:"products_services,omitempty"`
	IsSchool          bool     `json:"is_school,omitempty"`
	AnnualRevenueMin  int      `json:"annual_revenue_min,omitempty"`
	AnnualRevenueMax  int      `json:"annual_revenue_max,omitempty"`
	Page              int      `json:"page,omitempty"`
}

type PseParams struct {
	FullName                string   `json:"full_name,omitempty"`
	Country                 string   `json:"country,omitempty"`
	State                   string   `json:"state,omitempty"`
	City                    string   `json:"city,omitempty"`
	JobTitleRole            string   `json:"job_title_role,omitempty"`
	JobTitleLevel           string   `json:"job_title_level,omitempty"`
	CompanyCountry          string   `json:"company_country,omitempty"`
	CompanyState            string   `json:"company_state,omitempty"`
	CompanyCity             string   `json:"company_city,omitempty"`
	CompanyName             string   `json:"company_name,omitempty"`
	CompanyLinkedInURL      string   `json:"company_linkedin_url,omitempty"`
	CompanyIndustry         string   `json:"company_industry,omitempty"`
	CompanyEmployeeSize     string   `json:"company_employee_size,omitempty"`
	CompanyProductsServices []string `json:"company_products_services,omitempty"`
	CompanyAnnualRevenueMin int      `json:"company_annual_revenue_min,omitempty"`
	CompanyAnnualRevenueMax int      `json:"company_annual_revenue_max,omitempty"`
	Page                    int      `json:"page,omitempty"`
}

type LbsParams struct {
	Name     string `json:"name,omitempty"`
	Country  string `json:"country,omitempty"`
	State    string `json:"state,omitempty"`
	City     string `json:"city,omitempty"`
	Industry string `json:"industry,omitempty"`
	Page     int    `json:"page,omitempty"`
}

type BcdParams struct {
	Url string `json:"url"`
}

type CcpParams struct {
	Url string `json:"url"`
}

type IscParams struct {
	Url string `json:"url"`
}

type CbcParams struct {
	Url string `json:"url"`
}

type CscParams struct {
	Url string `json:"url"`
}
