# CUFinder Go SDK

A Go SDK for the CUFinder API that provides access to all company and person enrichment services.

## Installation

```bash
go get github.com/cufinder/cufinder-go
```

## Usage

### Basic Usage

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/cufinder/cufinder-go"
)

func main() {
    // Initialize the SDK
    sdk := cufinder.NewSDK("your-api-key")
    
    // Get company domain
    result, err := sdk.CUF("TechCorp", "US")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Domain: %s\n", result.Domain)
}
```

### Advanced Configuration

```go
import "time"

sdk := cufinder.NewSDKWithConfig(cufinder.ClientConfig{
    APIKey:     "your-api-key",
    BaseURL:    "https://api.cufinder.io/v2",
    Timeout:    30 * time.Second,
    MaxRetries: 3,
})
```

## Services

### Company Services

#### CUF - Company URL Finder
```go
result, err := sdk.CUF("TechCorp", "US")
// Returns: domain, query, credit_count
```

#### LCUF - LinkedIn Company URL Finder
```go
result, err := sdk.LCUF("TechCorp")
// Returns: linkedin_url, query, credit_count
```

#### DTC - Domain to Company
```go
result, err := sdk.DTC("techcorp.com")
// Returns: company_name, query, credit_count
```

#### DTE - Domain to Emails
```go
result, err := sdk.DTE("techcorp.com")
// Returns: emails[], query, credit_count
```

#### NTP - Name to Phones
```go
result, err := sdk.NTP("TechCorp")
// Returns: phones[], query, credit_count
```

### Person Services

#### EPP - Enrich Profile
```go
result, err := sdk.EPP("https://linkedin.com/in/john-doe")
// Returns: person, company, query, credit_count
```

#### REL - Reverse Email Lookup
```go
result, err := sdk.REL("john.doe@techcorp.com")
// Returns: person, company, query, credit_count
```

#### FWE - Find Work Email
```go
result, err := sdk.FWE("https://linkedin.com/in/john-doe")
// Returns: email, query, credit_count
```

#### TEP - Person Enrichment
```go
result, err := sdk.TEP("John Doe", "TechCorp")
// Returns: person, query, confidence_level, credit_count
```

### Company Intelligence Services

#### FCL - Find Company Lookalikes
```go
result, err := sdk.FCL("TechCorp")
// Returns: lookalikes[], query, credit_count
```

#### ELF - Enrich LinkedIn Fundraising
```go
result, err := sdk.ELF("TechCorp")
// Returns: fundraising, query, credit_count
```

#### CAR - Company Annual Revenue
```go
result, err := sdk.CAR("TechCorp")
// Returns: revenue, query, credit_count
```

#### FCC - Find Company Children
```go
result, err := sdk.FCC("TechCorp")
// Returns: subsidiaries[], query, credit_count
```

#### FTS - Find Tech Stack
```go
result, err := sdk.FTS("TechCorp")
// Returns: tech_stack, query, credit_count
```

#### ENC - Company Enrichment
```go
result, err := sdk.ENC("TechCorp")
// Returns: company, query, credit_count
```

#### CEC - Company Employee Countries
```go
result, err := sdk.CEC("TechCorp")
// Returns: countries[], total_results, query, credit_count
```

#### CLO - Company Locations
```go
result, err := sdk.CLO("TechCorp")
// Returns: locations[], query, credit_count
```

### Search Services

#### CSE - Company Search
```go
result, err := sdk.CSE(cufinder.CseParams{
    Name:     "technology",
    Country:  "US",
    Industry: "Technology",
    Page:     1,
})
// Returns: companies[], total_results, page, query, credit_count
```

#### PSE - Person Search
```go
result, err := sdk.PSE(cufinder.PseParams{
    FullName:    "engineer",
    CompanyName: "TechCorp",
    Page:        1,
})
// Returns: people[], total_results, page, query, credit_count
```

#### LBS - Local Business Search
```go
result, err := sdk.LBS(cufinder.LbsParams{
    Name:  "coffee",
    City:  "San Francisco",
    Page:  1,
})
// Returns: businesses[], total_results, page, query, credit_count
```

## Error Handling

The SDK returns errors for various scenarios:

```go
result, err := sdk.CUF("TechCorp", "US")
if err != nil {
    // Handle error
    log.Printf("Error: %v", err)
    return
}
```

## Testing

Run the tests:

```bash
go test -v
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
