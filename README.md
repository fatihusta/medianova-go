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
	organizationUUID := "MEDIANOVA_ORG_UUID"

	/*
		logLevel := &slog.LevelVar{}
		logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
		logLevel.Set(slog.LevelError)
		slog.SetDefault(logger)
	*/

	reqConfig := request.NewRequest()
	middlewares := client.WithMiddlewares(
		client.RetryMiddleware(3, 2*time.Second),
		client.AuthWithOrganizationToken(organizationToken),
		client.LoggingMiddleware(),
	)
	mn := client.NewClient(reqConfig, middlewares)

	resources, err := mn.CDN.Resource.List(organizationUUID)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	for _, r := range resources {
		fmt.Printf(" Resouce UUID:%s, Resource CDN URL:%s\n", r.ResourceUUID, r.CdnURL)
	}
}
```

# Supported API endpoints
- CDN > Resource > list
- CDN > Resource > show
- CDN > Purge > purge
- CDN > Purge > status
- CDN > Prefetch > prefetch
- CDN > Prefetch > status
- Analytics > Overview > historical
- Analytics > Overview > top resources
- Analytics > Overview > errors
- Analytics > Overview > visitor countries
- Analytics > Overview > active hit traffic
- Analytics > Overview > region
- Analytics > ErrorLogs > StatusCodes
- Analytics > ErrorReport > StatusCodes
- Analytics > RequestReport > detail
- Analytics > StatusReport > status report
- Analytics > TrafficReport > detail report
