# CUFinder Go SDK

[![](https://img.shields.io/badge/repo%20status-Active-28a745)](https://github.com/cufinder/cufinder-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-514BEE.svg)](https://opensource.org/licenses/MIT)
[![Go Reference](https://pkg.go.dev/badge/github.com/cufinder/cufinder-go.svg)](https://pkg.go.dev/github.com/cufinder/cufinder-go)

A Go SDK for the CUFinder API that provides access to all company and person enrichment services.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [API Reference](#api-reference)
- [Error Handling](#error-handling)
- [Types](#types)
- [Support](#support)

## Installation

```bash
go get github.com/cufinder/cufinder-go
```

## Usage

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/cufinder/cufinder-go"
)

func main() {
    // Initialize the client
    sdk := cufinder.NewSDK("your-api-key-here")
    
    // Initialize with more options
    sdk := cufinder.NewSDKWithConfig(cufinder.ClientConfig{
        APIKey:     "your-api-key-here",
        BaseURL:    "https://api.cufinder.io/v2",
        Timeout:    60 * time.Second,
        MaxRetries: 3,
    })
}
```

## API Reference

This SDK covers all 40 Cufinder API (v2) endpoints:

- **CUF** - [Company Name to Domain](https://apidoc.cufinder.io/apis/company-name-to-domain)
- **LCUF** - [LinkedIn Company URL Finder](https://apidoc.cufinder.io/apis/company-linkedin-url-finder)
- **DTC** - [Domain to Company Name](https://apidoc.cufinder.io/apis/domain-to-company-name)
- **DTE** - [Company Email Finder](https://apidoc.cufinder.io/apis/company-email-finder)
- **NTP** - [Company Phone Finder](https://apidoc.cufinder.io/apis/company-phone-finder)
- **REL** - [Reverse Email Lookup](https://apidoc.cufinder.io/apis/reverse-email-lookup)
- **FCL** - [Company Lookalikes Finder](https://apidoc.cufinder.io/apis/company-lookalikes-finder)
- **ELF** - [Company Fundraising](https://apidoc.cufinder.io/apis/company-fundraising)
- **CAR** - [Company Revenue Finder](https://apidoc.cufinder.io/apis/company-revenue-finder)
- **FCC** - [Company Subsidiaries Finder](https://apidoc.cufinder.io/apis/company-subsidiaries-finder)
- **FTS** - [Company Tech Stack Finder](https://apidoc.cufinder.io/apis/company-tech-stack-finder)
- **EPP** - [LinkedIn Profile Enrichment](https://apidoc.cufinder.io/apis/linkedin-profile-enrichment)
- **FWE** - [LinkedIn Profile Email Finder](https://apidoc.cufinder.io/apis/linkedin-profile-email-finder)
- **TEP** - [Person Enrichment](https://apidoc.cufinder.io/apis/person-enrichment)
- **ENC** - [Company Enrichment](https://apidoc.cufinder.io/apis/company-enrichment)
- **CEC** - [Company Employee Count](https://apidoc.cufinder.io/apis/company-employee-count)
- **CLO** - [Company Locations](https://apidoc.cufinder.io/apis/company-locations)
- **CSE** - [Company Search](https://apidoc.cufinder.io/apis/company-search)
- **PSE** - [Person Search](https://apidoc.cufinder.io/apis/person-search)
- **LBS** - [Local Business Search (Google Maps Search API)](https://apidoc.cufinder.io/apis/local-business-search-google-maps-search-api)
- **BCD** - [B2B Customers Finder](https://apidoc.cufinder.io/apis/b2b-customers-finder)
- **CCP** - [Company Career Page Finder](https://apidoc.cufinder.io/apis/company-career-page-finder)
- **ISC** - [Company Saas Checker](https://apidoc.cufinder.io/apis/company-saas-checker)
- **CBC** - [Company B2B or B2C Checker](https://apidoc.cufinder.io/apis/company-b2b-or-b2c-checker)
- **CSC** - [Company Mission Statement](https://apidoc.cufinder.io/apis/company-mission-statement)
- **CSN** - [Company Snapshot](https://apidoc.cufinder.io/apis/company-snapshot)
- **NAO** - [Phone Number Normalizer](https://apidoc.cufinder.io/apis/phone-number-normalizer)
- **NAA** - [Address Normalizer](https://apidoc.cufinder.io/apis/address-normalizer)
- **CEF** - [Company Employee Finder](https://apidoc.cufinder.io/apis/company-employee-finder)
- **NAC** - [Company Name Normalizer](https://apidoc.cufinder.io/apis/company-name-normalizer)
- **CAA** - [Company Activity API](https://apidoc.cufinder.io/apis/company-activity-api)
- **CJA** - [Company Jobs API](https://apidoc.cufinder.io/apis/company-jobs-api)
- **PSA** - [Contact Signals API](https://apidoc.cufinder.io/apis/contact-signals-api)
- **CSA** - [Company Signals API](https://apidoc.cufinder.io/apis/company-signals-api)
- **JCA** - [Job Changes API](https://apidoc.cufinder.io/apis/job-changes-api)
- **CLF** - [Contact Lookalikes API](https://apidoc.cufinder.io/apis/contact-lookalikes-api)
- **NAP** - [Person Name Normalizer](https://apidoc.cufinder.io/apis/person-name-normalizer)
- **NAU** - [URL Normalizer](https://apidoc.cufinder.io/apis/url-normalizer)
- **GDC** - [Gives Demo Checker](https://apidoc.cufinder.io/apis/gives-demo-checker)
- **COT** - [Offers Free Trial Checker](https://apidoc.cufinder.io/apis/offers-free-trial-checker)


**CUF - Company Name to Domain API**

Returns the official website URL of a company based on its name.

```go
result, err := sdk.CUF("cufinder", "US")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

**LCUF - Company LinkedIn URL Finder API**

Finds the official LinkedIn company profile URL from a company name.

```go
result, err := sdk.LCUF("cufinder")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

**DTC - Domain to Company Name API**

Retrieves the registered company name associated with a given website domain.

```go
result, err := sdk.DTC("cufinder.io")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

**DTE - Company Email Finder API**

Returns up to five general or role-based business email addresses for a company.

```go
result, err := sdk.DTE("cufinder.io")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

**NTP - Company Phone Finder API**

Returns up to two verified phone numbers for a company.

```go
result, err := sdk.NTP("apple")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

**REL - Reverse Email Lookup API**

Enriches an email address with detailed person and company information.

```go
result, err := sdk.REL("iain.mckenzie@stripe.com")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

**FCL - Company Lookalikes Finder API**

Provides a list of similar companies based on an input company's profile.

```go
result, err := sdk.FCL("apple")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

**ELF - Company Fundraising API**

Returns detailed funding information about a company.

```go
result, err := sdk.ELF("cufinder")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

**CAR - Company Revenue Finder API**

Estimates a company's annual revenue based on name.

```go
result, err := sdk.CAR("apple")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

**FCC - Company Subsidiaries Finder API**

Identifies known subsidiaries of a parent company.

```go
result, err := sdk.FCC("amazon")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

**FTS - Company Tech Stack Finder API**

Detects the technologies a company uses.

```go
result, err := sdk.FTS("cufinder")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

**EPP - LinkedIn Profile Enrichment API**

Takes a LinkedIn profile URL and returns enriched person and company data.

```go
result, err := sdk.EPP("linkedin.com/in/iain-mckenzie")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

**FWE - LinkedIn Profile Email Finder API**

Extracts a verified business email address from a LinkedIn profile URL.

```go
result, err := sdk.FWE("linkedin.com/in/iain-mckenzie")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

**TEP - Person Enrichment API**

Returns enriched person data based on full name and company name.

```go
result, err := sdk.TEP("iain mckenzie", "stripe")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

**ENC - Company Enrichment API**

Provides a complete company profile from a company name.

```go
result, err := sdk.ENC("cufinder")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

**CEC - Company Employee Count API**

Returns an estimated number of employees for a company.

```go
result, err := sdk.CEC("cufinder")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

**CLO - Company Locations API**

Returns the known physical office locations of a company.

```go
result, err := sdk.CLO("apple")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

**CSE - Company Search API**

Search for companies by keyword, partial name, industry, location, or other filters.

```go
result, err := sdk.CSE(cufinder.CseParams{
    Name:    "cufinder",
    Country: "germany",
    State:   "hamburg",
    City:    "hamburg",
})
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

**PSE - Person Search API**

Search for people by name, company, job title, location, or other filters.

```go
result, err := sdk.PSE(cufinder.PseParams{
    FullName:    "iain mckenzie",
    CompanyName: "stripe",
})
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

**LBS - Local Business Search API (Google Maps Search API)**

Search for local businesses by location, industry, or name.

```go
result, err := sdk.LBS(cufinder.LbsParams{
    Country: "united states",
    State:   "california",
    Page:    1,
})
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

**BCD - B2B Customers Finder**

Returns company's careers page

```go
result, err := sdk.BCD("stripe.com")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

**CCP - Company Career Page Finder**

Returns is company SaaS or not

```go
result, err := sdk.CCP("stripe.com")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

**ISC - Company Saas Checker**

Returns is company SaaS or not

```go
result, err := sdk.ISC("stripe.com")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

**CBC - Company B2B or B2C Checker**

Returns company's business type

```go
result, err := sdk.CBC("stripe.com")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

**CSC - Company Mission Statement**

Returns company's mission statement

```go
result, err := sdk.CSC("stripe.com")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

**CSN - Company Snapshot**

Returns company's snapshot information

```go
result, err := sdk.CSN("stripe.com")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

**NAO - Phone Number Normalizer**

Returns normalized phone

```go
result, err := sdk.NAO("+18006676389")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

**NAA - Address Normalizer**

Returns normalized address

```go
result, err := sdk.NAA("1095 avenue of the Americas, 6th Avenue ny 10036")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

**CEF - Company Employee Finder**

Returns a list of employees for a given company

```go
result, err := sdk.CEF("stripe", 1)
if err != nil {
    log.Fatal(err)
}
for _, emp := range result.Employees {
    fmt.Printf("name: %v, title: %v\n", emp.FullName, emp.JobTitle)
}
```

**NAC - Company Name Normalizer**

Returns normalized company name

```go
result, err := sdk.NAC("cufinder inc.")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result.Company)
```

**CAA - Company Activity API**

Returns company LinkedIn activities

```go
result, err := sdk.CAA("cufinder", 1)
if err != nil {
    log.Fatal(err)
}
for _, act := range result.Activities {
    fmt.Printf("id: %v, url: %v\n", act.ActivityID, act.ActivityURL)
}
```

**CJA - Company Jobs API**

Search company job listings with filters

```go
result, err := sdk.CJA(cufinder.CjaParams{
    Name:    "google",
    Country: "united states",
    Page:    1,
})
if err != nil {
    log.Fatal(err)
}
for _, j := range result.Jobs {
    fmt.Printf("company: %v, title: %v, location: %v\n", j.Company.Name, j.Job.Title, j.Job.Location)
}
```

**PSA - Contact Signals API**

Find contacts based on company signals

```go
result, err := sdk.PSA(cufinder.PsaParams{
    SignalName: "employee_growth",
    TimeFrame:  90,
    Bucket:     "high",
    Page:       1,
})
if err != nil {
    log.Fatal(err)
}
for _, c := range result.Contacts {
    fmt.Printf("name: %v, company: %v, job: %v\n", c.FullName, c.Company.Name, c.CurrentJob.Title)
}
```

**CSA - Company Signals API**

Find companies based on signals

```go
result, err := sdk.CSA(cufinder.CsaParams{
    SignalName: "employee_growth",
    TimeFrame:  90,
    Bucket:     "high",
    Page:       1,
})
if err != nil {
    log.Fatal(err)
}
for _, c := range result.Companies {
    fmt.Printf("company: %v, domain: %v, industry: %v\n", c.Name, c.Domain, c.Industry)
}
```

**JCA - Job Changes API**

Find job changes within a date range

```go
result, err := sdk.JCA(cufinder.JcaParams{
    StartDate: "2026-01-01",
    EndDate:   "2026-08-16",
    Type:      "promotion",
})
if err != nil {
    log.Fatal(err)
}
for _, jc := range result.JobChanges {
    fmt.Printf("type: %v, from: %v, to: %v\n", jc.Type, jc.From.Title, jc.To.Title)
}
```

**CLF - Contact Lookalikes API**

Find similar contacts based on a query

```go
result, err := sdk.CLF("linkedin.com/in/mortezaheydari1997")
if err != nil {
    log.Fatal(err)
}
for _, p := range result.Profiles {
    fmt.Printf("name: %v, company: %v, job: %v\n", p.FullName, p.CompanyName, p.JobTitle)
}
```

**NAP - Person Name Normalizer**

Normalize a person name

```go
result, err := sdk.NAP("morteza heydari")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("normalized: %v\n", result.NormalizedName)
```

**NAU - URL Normalizer**

Normalize a URL

```go
result, err := sdk.NAU("https://www.cufinder.io/about-us")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("normalized: %v\n", result.NormalizedUrl)
```

**GDC - Gives Demo Checker**

Check if a company offers demos

```go
result, err := sdk.GDC("https://www.stripe.com")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("offers demo: %v\n", result.OffersDemo)
```

**COT - Offers Free Trial Checker**

Check if a company offers a free trial

```go
result, err := sdk.COT("https://www.stripe.com")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("offers free trial: %v\n", result.OffersFreeTrial)
```

## Error Handling

The SDK returns errors for various scenarios:

```go
result, err := sdk.CUF("cufinder", "US")
if err != nil {
    // Handle different types of errors
    switch {
    case strings.Contains(err.Error(), "401"):
        // Authentication error - Invalid API key
        log.Printf("Authentication failed: %v", err)
    case strings.Contains(err.Error(), "400"):
        // Credit limit error - Not enough credit
        log.Printf("Not enough credit: %v", err)
    case strings.Contains(err.Error(), "404"):
        // Not found error
        log.Printf("Not found result: %v", err)
    case strings.Contains(err.Error(), "422"):
        // Payload error
        log.Printf("Payload error: %v", err)
    case strings.Contains(err.Error(), "429"):
        // Rate limit exceeded
        log.Printf("Rate limit exceeded: %v", err)
    case strings.Contains(err.Error(), "500"):
        // Server error
        log.Printf("Server error: %v", err)
    default:
        log.Printf("Unknown error: %v", err)
    }
    return
}
```

## Types

The SDK exports comprehensive Go types for all API requests and responses:

```go
// Request parameter types
type CseParams struct {
    Name            string
    Domain          string
    Country         string
    State           string
    City            string
    Industry        string
    CompanySize     string
    Revenue         string
    EmployeeCount   string
    Page            int
}

type PseParams struct {
    FullName        string
    FirstName       string
    LastName        string
    CompanyName     string
    CompanyDomain   string
    JobTitle        string
    Country         string
    State           string
    City            string
    Page            int
}

type LbsParams struct {
    Name            string
    Country         string
    State           string
    City            string
    Industry        string
    Page            int
}

type PsaParams struct {
    SignalName  string
    TimeFrame   int
    Bucket      string
    Page        int
}

type CsaParams struct {
    SignalName  string
    TimeFrame   int
    Bucket      string
    Page        int
}

type JcaParams struct {
    StartDate   string
    EndDate     string
    Type        string
    Page        int
}

type ClfParams struct {
    Query       string
}

// Response types
type BaseResponse struct {
    Query       string `json:"query"`
    CreditCount int    `json:"credit_count"`
}

// Model types
type Company struct {
    // The Company struct contains all returned company data.
    Name            string   `json:"name"`
    Domain          string   `json:"domain"`
    Website         string   `json:"website"`
    LinkedinURL     string   `json:"linkedin_url"`
    Country         string   `json:"country"`
    State           string   `json:"state"`
    City            string   `json:"city"`
    Address         string   `json:"address"`
    Industry        string   `json:"industry"`
    CompanySize     string   `json:"company_size"`
    Revenue         string   `json:"revenue"`
    EmployeeCount   int      `json:"employee_count"`
    Subsidiaries    []string `json:"subsidiaries"`
    TechStack       []string `json:"tech_stack"`
    Emails          []string `json:"emails"`
    Phones          []string `json:"phones"`
    Description     string   `json:"description"`
    Locations       []CompanyLocation `json:"locations"`
    FoundedYear     int      `json:"founded_year"`
    LogoURL         string   `json:"logo_url"`
}

type Person struct {
    // The Person struct contains all returned person data.
    FullName        string   `json:"full_name"`
    FirstName       string   `json:"first_name"`
    LastName        string   `json:"last_name"`
    CompanyName     string   `json:"company_name"`
    CompanyDomain   string   `json:"company_domain"`
    JobTitle        string   `json:"job_title"`
    Country         string   `json:"country"`
    State           string   `json:"state"`
    City            string   `json:"city"`
    Email           string   `json:"email"`
    Phone           string   `json:"phone"`
    Description     string   `json:"description"`
    LinkedInURL     string   `json:"linkedin_url"`
}

type LookalikeCompany struct {
    // The LookalikeCompany struct contains all returned lookalike company data.
    Name            string   `json:"name"`
    Domain          string   `json:"domain"`
    Website         string   `json:"website"`
    LinkedinURL     string   `json:"linkedin_url"`
    Country         string   `json:"country"`
    State           string   `json:"state"`
    City            string   `json:"city"`
    Address         string   `json:"address"`
    Industry        string   `json:"industry"`
    CompanySize     string   `json:"company_size"`
    Revenue         string   `json:"revenue"`
    EmployeeCount   int      `json:"employee_count"`
    Subsidiaries    []string `json:"subsidiaries"`
    TechStack       []string `json:"tech_stack"`
    Emails          []string `json:"emails"`
    Phones          []string `json:"phones"`
    Description     string   `json:"description"`
    Locations       []CompanyLocation `json:"locations"`
    FoundedYear     int      `json:"founded_year"`
    LogoURL         string   `json:"logo_url"`
}

type FundraisingInfo struct {
    // Fundraising fields
    FundingLastRoundType         string `json:"funding_last_round_type"`
    FundingAmmountCurrencyCode   string `json:"funding_ammount_currency_code"`
    FundingMoneyRaised           string `json:"funding_money_raised"`
    FundingLastRoundInvestorsUrl string `json:"funding_last_round_investors_url"`
}

type CompanyLocation struct {
    // The CompanyLocation struct contains all returned company location data.
    Country    string `json:"country"`
    State      string `json:"state"`
    City       string `json:"city"`
    PostalCode string `json:"postal_code"`
    Line1      string `json:"line1"`
    Line2      string `json:"line2"`
    Latitude   string `json:"latitude"`
    Longitude  string `json:"longitude"`
}
```

## Support

For support, please open an issue on the [GitHub repository](https://github.com/cufinder/cufinder-go/issues).
