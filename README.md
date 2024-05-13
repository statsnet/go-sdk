# Golang SDK for statsnet.co
## Installation

```sh
go get github.com/statsnet/gosdk
```

## Usage
Import statsnet client and set API keys:

```go
package main
import (
    "github.com/statsnet/gosdk"
)

func main() {
	const token = "123" // or set via STATSNET_API_KEY environment variable
}
```

```go
client, err := NewClient(token)
```

```go
// Get info about current user (/api/v2/user/me)
me, err := client.Me(context.Background())
```

```go
// Find company by query and jurisdiction (/api/v2/business/search)
jurisdiction := "kz"
limit := 5
companies, err := client.Search(context.Background(), "казпочта", &jurisdiction, &limit)
// Or without specifiying jurisdiction
companies, err := client.Search(context.Background(), "казпочта", nil, &limit)
```

```go
// Get company by id and jurisdiction (/api/v2/business/{jurisdiction}/{id}/paid
company, err := client.GetCompany(context.Background(), "kz", 1)
```

```go
// Get company meta by id (/api/v2/business/{id}/data/view/meta)
meta, err := client.GetCompanyMeta(context.Background(), 1)
```

```go
// Get company court cases (/business/{id}/court_cases)
limit := 5
courtCases, err := client.GetCompanyCourtCases(context.Background(), 1, &limit)
```

```go
// Get company departments (/business/{id}/department)
limit := 5
departments, err := client.GetCompanyDepartments(context.Background(), 1, &limit)
```

```go
// Get company gov contracts (/business/{id}/gov_contracts)
limit := 5
govContracts, err := client.GetCompanyGovContracts(context.Background(), 1, &limit)
```

```go
// Get company events (/business/{id}/events)
limit := 5
companyEvents, err := client.GetCompanyEvents(context.Background(), 1, &limit)
```

```go
// Get company relations (/business/{id}/relations/table)
limit := 5
companyRelations, err := client.GetCompanyRelations(context.Background(), 1, &limit)
```

```go
// Get company by identifier
company, err := client.GetCompanyByIdentifier(context.Background(), "1")
```

##### Limit must be between 1 and 500, 501 is already invalid value