# Medianova Cloud API Go Client Library

This is an unofficial library.

> I have only implemented the parts I needed. \
> I may develop the rest of the API as time allows. \
> Feel free to submit pull requests.


[Medianova Official Documentation](https://clients.medianova.com/api-documentation)

# Example usage
```go
package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/fatihusta/medianova-go/client"
	"github.com/fatihusta/medianova-go/client/request"
)

func main() {
	organizationToken := "MEDIANOVA_ORG_TOKEN"

	/*
		logLevel := &slog.LevelVar{}
		logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
		logLevel.Set(slog.LevelError)
		slog.SetDefault(logger)
	*/

	reqConfig := request.NewRequestConfig()
	middlewares := client.WithMiddlewares(
		client.RetryMiddleware(3, 2*time.Second),
		client.AuthWithOrganizationToken(organizationToken),
	)
	mn := client.NewClient(reqConfig, middlewares)

	organizations := mn.Organization.List()
	if organizations.Error != nil {
		slog.Error(organizations.Error.Error())
		os.Exit(1)
	}

	resources := mn.CDN.Resource.List(organizations.Body[0].UUID)
	if resources.Error != nil {
		slog.Error(resources.Error.Error())
		os.Exit(1)
	}

	for _, r := range resources.Body {
		fmt.Printf("Resouce UUID:%s, Resource CDN URL:%s\n", r.ResourceUUID, r.CdnURL)
	}
}
```

> Please contact to your Medianova Solution Engineer to obtain your Organization Token.

# Live integration tests

The integration tests send real requests to the Medianova Cloud API and verify the responses.
They live next to the packages they cover and read their configuration from environment variables.
Tests that require credentials are skipped automatically when `MEDIANOVA_ORG_TOKEN` is not set,
so plain `go test ./...` stays safe to run without a Medianova account.

Copy `.env.example` to `.env`, replace the dummy values with your own and load it:

```bash
cp .env.example .env
# edit .env with your real values, then:
source .env

go test ./... -v          # run all tests
go test ./analytics/resource/ -v   # run a single package
```

| Variable | Required | Used by | Fallback |
|---|---|---|---|
| `MEDIANOVA_ORG_TOKEN` | yes | all live tests | none, tests are skipped when it is not set |
| `MEDIANOVA_ORG_UUID` | no | all live tests | first organization of the token |
| `MEDIANOVA_RESOURCE_CDN_URL` | no | tests that need a CDN resource | first resource (`cdn_url`) of the organization |
| `MEDIANOVA_ANALYTICS_FROM` | no | analytics tests | one hour window from yesterday |
| `MEDIANOVA_ANALYTICS_TO` | no | analytics tests | one hour window from yesterday |

Notes:
- The API only accepts `"YYYY-MM-DD HH:MM:SS"` date format.
- Granular graph data is not retained for older dates, keep the default recent window unless you need a specific range.
- Never commit `.env`, it is already in `.gitignore`.

# Tier filtering

All analytics request structs (v1 overview/reports and v2 resource) accept an optional `Tier` field
(`tier.TierEdge` or `tier.TierAll`, defined in `analytics/tier`):

- `TierEdge`: counts only end-user (Edge) requests. It must be set explicitly, the API does not default to edge.
- `TierAll`: counts requests across the Edge + Mid + Origin layers, the same request may be counted more than once.
  It is the API default; leaving the field empty sends no tier and behaves as `all`.

```go
result := mn.Analytics.ResourceAnalytics.GetRequestReport(ctx, resource.ResourceHistogramRequest{
	ResourceRequest: resource.ResourceRequest{
		OrganizationUUID: organizationUUID,
		From:             "2024-11-12 05:00:00",
		To:               "2024-11-12 05:59:59",
		Tier:             tier.TierAll,
	},
	Interval: "5m",
})
```

Invalid values are rejected client side before any HTTP request is sent.

# Supported API endpoints
- Analytics > Overview > historical (hit, hit_ratio, traffic, gpu_transaction, bandwidth, requests)
- Analytics > Overview > top resources
- Analytics > Overview > errors
- Analytics > Overview > visitor countries
- Analytics > Overview > active hit traffic
- Analytics > Overview > region
- Analytics > Resource > historical (hit, hit_ratio, traffic, gpu_transaction, bandwidth, requests)
- Analytics > Resource > traffic / request / bandwidth report graphs
- Analytics > Resource > cache / non-cache report graphs
- Analytics > Resource > cache / non-cache traffic report graphs
- Analytics > Resource > status report graphs (2xx, 3xx, 4xx, 5xx)
- Analytics > Resource > errors
- Analytics > Resource > visitors countries
- Analytics > Resource > status
- Analytics > Resource > region
- Analytics > Resource > active traffic / active hit
- Analytics > Resource > traffic detail
- Analytics > Resource > request detail
- Analytics > Stook > overview
- Analytics > Stook > filters
- Analytics > Stook > request histogram
- Analytics > Stook > error ratio histogram
- Analytics > Stook > charts (metric, per_bucket, per_endpoint, per_method, per_country)
- Analytics > Stook > bucket details
- Analytics > Stook > failure rate per bucket
- Analytics > Stook > raw logs
- Analytics > ErrorLogs > StatusCodes
- Analytics > ErrorReport > StatusCodes
- Analytics > RequestReport > detail
- Analytics > StatusReport > status report
- Analytics > TrafficReport > detail report
- CDN > Resource > list
- CDN > Resource > show
- CDN > Purge > purge
- CDN > Purge > status
- CDN > Prefetch > prefetch
- CDN > Prefetch > status
- CDN > SSL > list
- Organization > list
- Organization > show
- Organization > users
