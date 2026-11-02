package resource_test

import (
	"context"
	"net/http"
	"os"
	"sync"
	"testing"
	"time"

	mnresource "github.com/fatihusta/medianova-go/analytics/resource"
	"github.com/fatihusta/medianova-go/client"
	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/common"
)

const dateLayout = "2006-01-02 15:04:05"

// defaultWindow returns yesterday's window in "YYYY-MM-DD HH:MM:SS" format,
// the only format accepted by the API. Granular graph data is not retained
// for older dates, so recent dates are used.
func defaultWindow() (string, string) {
	from := time.Now().Add(-26 * time.Hour).Truncate(time.Hour)
	return from.Format(dateLayout), from.Add(time.Hour).Format(dateLayout)
}

func newHTTPClient(token string) *client.Client {
	reqConfig := request.NewRequestConfig()
	reqConfig.RequestTimeout = 30 * time.Second

	middlewares := client.WithMiddlewares(
		client.RetryMiddleware(2, 5*time.Second),
		client.AuthWithOrganizationToken(token),
	)

	return client.NewClient(reqConfig, middlewares)
}

func newTestClient(t *testing.T, token string) *client.Client {
	t.Helper()
	return newHTTPClient(token)
}

// testEnv holds the organization uuid and resource hostname used by the live
// tests. Values come from the environment when set, otherwise they are
// resolved from the API: MEDIANOVA_ORG_TOKEN is the only required variable,
// MEDIANOVA_ORG_UUID falls back to the first organization of the token,
// MEDIANOVA_RESOURCE_CDN_URL falls back to the first resource of the organization.
type testEnv struct {
	mn      *client.Client
	orgUUID string
	cdnURL  string
	from    string
	to      string
}

var (
	resolveOnce sync.Once
	resolvedEnv testEnv
	resolveErr  error
)

func newTestEnv(t *testing.T) testEnv {
	t.Helper()

	token := os.Getenv("MEDIANOVA_ORG_TOKEN")
	if token == "" {
		t.Skip("MEDIANOVA_ORG_TOKEN is not set, skipping live API test")
	}

	env := newTestEnvOnce(token)
	env.mn = newTestClient(t, token)

	if from, to := os.Getenv("MEDIANOVA_ANALYTICS_FROM"), os.Getenv("MEDIANOVA_ANALYTICS_TO"); from != "" && to != "" {
		env.from, env.to = from, to
	}
	return env
}

func newTestEnvOnce(token string) testEnv {
	resolveOnce.Do(func() {
		mn := newHTTPClient(token)

		env := testEnv{}
		env.from, env.to = defaultWindow()

		env.orgUUID = os.Getenv("MEDIANOVA_ORG_UUID")
		if env.orgUUID == "" {
			organizations := mn.Organization.List()
			if organizations.Error != nil {
				resolveErr = organizations.Error
				return
			}
			if len(organizations.Body) == 0 {
				resolveErr = errNotFound("no organization found for the given token")
				return
			}
			env.orgUUID = organizations.Body[0].UUID
		}

		env.cdnURL = os.Getenv("MEDIANOVA_RESOURCE_CDN_URL")
		if env.cdnURL == "" {
			resources := mn.CDN.Resource.List(env.orgUUID)
			if resources.Error != nil {
				resolveErr = resources.Error
				return
			}
			for _, r := range resources.Body {
				if r.CdnURL != "" {
					env.cdnURL = r.CdnURL
					break
				}
			}
			if env.cdnURL == "" {
				resolveErr = errNotFound("no resource with a cdn url found in organization " + env.orgUUID)
				return
			}
		}

		resolvedEnv = env
	})

	if resolveErr != nil {
		panic(resolveErr)
	}
	return resolvedEnv
}

type errNotFound string

func (e errNotFound) Error() string { return string(e) }

func (e testEnv) request() mnresource.ResourceRequest {
	return mnresource.ResourceRequest{
		OrganizationUUID: e.orgUUID,
		Resources:        []string{e.cdnURL},
		From:             e.from,
		To:               e.to,
	}
}

func (e testEnv) histogramRequest() mnresource.ResourceHistogramRequest {
	return mnresource.ResourceHistogramRequest{
		ResourceRequest: e.request(),
		Interval:        "5m",
	}
}

func (e testEnv) limitRequest() mnresource.ResourceLimitRequest {
	limit := 10
	page := 1
	return mnresource.ResourceLimitRequest{
		ResourceRequest: e.request(),
		Limit:           &limit,
		Page:            &page,
	}
}

func requireOK[T any](t *testing.T, result *common.Result[T]) {
	t.Helper()

	if result.Error != nil {
		t.Fatalf("request failed: %s", result.Error)
	}
	if result.Status != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, result.Status)
	}
}

func withCtx(t *testing.T) context.Context {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	t.Cleanup(cancel)
	return ctx
}

func TestHistorical(t *testing.T) {
	env := newTestEnv(t)

	tests := []struct {
		name  string
		typ   mnresource.StatType
		value func(period mnresource.HistoricalMetricPeriod) mnresource.Number
	}{
		{"hit", mnresource.Hit, func(p mnresource.HistoricalMetricPeriod) mnresource.Number { return p.Hit }},
		{"traffic", mnresource.Traffic, func(p mnresource.HistoricalMetricPeriod) mnresource.Number { return p.Traffic }},
		{"bandwidth", mnresource.Bandwidth, func(p mnresource.HistoricalMetricPeriod) mnresource.Number { return p.Bandwidth }},
		{"requests", mnresource.Requests, func(p mnresource.HistoricalMetricPeriod) mnresource.Number { return p.Count }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := mnresource.ResourceHistoricalRequest{
				ResourceRequest: env.request(),
				Type:            tt.typ,
			}

			result := env.mn.Analytics.ResourceAnalytics.GetHistorical(withCtx(t), req)
			requireOK(t, result)

			if !result.Body.Status {
				t.Fatalf("expected body status true, got: %s", result.JSON().String())
			}
			if len(result.Body.Payload.TimePeriods) == 0 {
				t.Fatalf("expected non-empty time_periods, got: %s", result.JSON().String())
			}
			if result.Body.Payload.Overall.Payload.SelectedTime.From != req.From {
				t.Fatalf("expected overall selected_time.from %q, got: %s", req.From, result.JSON().String())
			}

			// the API fills only the field related to the requested type
			overall := tt.value(result.Body.Payload.Overall.Payload.SelectedTime)
			if _, err := overall.Float64(); err != nil {
				t.Fatalf("expected overall %s %q to parse as float: %s", tt.name, overall, err)
			}
			perResource, ok := result.Body.Payload.PerResources[env.cdnURL]
			if !ok {
				t.Fatalf("expected per_resources entry for %s, got: %s", env.cdnURL, result.JSON().String())
			}
			if _, err := tt.value(perResource.SelectedTime).Float64(); err != nil {
				t.Fatalf("expected per-resource %s to parse as float: %s", tt.name, err)
			}

			t.Logf("historical[%s] overall: %+v", tt.name, result.Body.Payload.Overall.Payload.SelectedTime)
		})
	}
}

func TestHistoricalHitRatio(t *testing.T) {
	env := newTestEnv(t)

	result := env.mn.Analytics.ResourceAnalytics.GetHistoricalHitRatio(withCtx(t), env.request())
	requireOK(t, result)

	if !result.Body.Status {
		t.Fatalf("expected body status true, got: %s", result.JSON().String())
	}
	if _, err := result.Body.Payload.Overall.HitRatio.Float64(); err != nil {
		t.Fatalf("expected overall hit_ratio %q to parse as float: %s", result.Body.Payload.Overall.HitRatio, err)
	}
	if _, ok := result.Body.Payload.PerResources[env.cdnURL]; !ok {
		t.Fatalf("expected per_resources entry for %s, got: %s", env.cdnURL, result.JSON().String())
	}
	t.Logf("historical hit_ratio overall: %s", result.Body.Payload.Overall.HitRatio)
}

func TestHistoricalGPUTransaction(t *testing.T) {
	env := newTestEnv(t)

	result := env.mn.Analytics.ResourceAnalytics.GetHistoricalGPUTransaction(withCtx(t), env.request())
	requireOK(t, result)

	if !result.Body.Status {
		t.Fatalf("expected body status true, got: %s", result.JSON().String())
	}
	// the API returns all organization resources for this type
	if len(result.Body.Payload.PerResources) == 0 {
		t.Fatalf("expected non-empty per_resources, got: %s", result.JSON().String())
	}
	if _, err := result.Body.Payload.Overall.GPUTransaction.Float64(); err != nil {
		t.Fatalf("expected overall gpu_transaction %q to parse as float: %s", result.Body.Payload.Overall.GPUTransaction, err)
	}
	t.Logf("historical gpu_transaction overall: %s", result.Body.Payload.Overall.GPUTransaction)
}

func TestHistoricalWithoutType(t *testing.T) {
	env := newTestEnv(t)

	result := env.mn.Analytics.ResourceAnalytics.GetHistorical(withCtx(t), mnresource.ResourceHistoricalRequest{
		ResourceRequest: env.request(),
	})

	if result.Error == nil {
		t.Fatal("expected validation error for empty type")
	}
}

func TestHistoricalHitRatioOnGetHistorical(t *testing.T) {
	env := newTestEnv(t)

	result := env.mn.Analytics.ResourceAnalytics.GetHistorical(withCtx(t), mnresource.ResourceHistoricalRequest{
		ResourceRequest: env.request(),
		Type:            mnresource.HitRatio,
	})

	if result.Error == nil {
		t.Fatal("expected error directing to GetHistoricalHitRatio")
	}
}

func TestTrafficReport(t *testing.T) {
	env := newTestEnv(t)

	result := env.mn.Analytics.ResourceAnalytics.GetTrafficReport(withCtx(t), env.histogramRequest())
	requireOK(t, result)

	if !result.Body.Status {
		t.Fatalf("expected body status true, got: %s", result.JSON().String())
	}
	if len(result.Body.Payload.Overall.Progressive) == 0 && len(result.Body.Payload.Overall.Streaming) == 0 {
		t.Fatalf("expected non-empty overall histogram, got: %s", result.JSON().String())
	}
	if _, ok := result.Body.Payload.PerResources[env.cdnURL]; !ok {
		t.Fatalf("expected per_resources entry for %s, got: %s", env.cdnURL, result.JSON().String())
	}

	points := result.Body.Payload.Overall.Progressive
	if len(points) == 0 {
		points = result.Body.Payload.Overall.Streaming
	}
	if _, err := points[0].Traffic.Float64(); err != nil {
		t.Fatalf("expected traffic to parse as float: %s", err)
	}
	t.Logf("traffic_report point: %+v", points[0])
}

func TestRequestReport(t *testing.T) {
	env := newTestEnv(t)

	result := env.mn.Analytics.ResourceAnalytics.GetRequestReport(withCtx(t), env.histogramRequest())
	requireOK(t, result)

	if !result.Body.Status {
		t.Fatalf("expected body status true, got: %s", result.JSON().String())
	}
	if len(result.Body.Payload.Overall.Progressive) == 0 && len(result.Body.Payload.Overall.Streaming) == 0 {
		t.Fatalf("expected non-empty overall histogram, got: %s", result.JSON().String())
	}

	points := result.Body.Payload.Overall.Progressive
	if len(points) == 0 {
		points = result.Body.Payload.Overall.Streaming
	}
	if _, err := points[0].Total.Float64(); err != nil {
		t.Fatalf("expected total to parse as float: %s", err)
	}
	t.Logf("request_report point: %+v", points[0])
}

func TestBandwidthReport(t *testing.T) {
	env := newTestEnv(t)

	result := env.mn.Analytics.ResourceAnalytics.GetBandwidthReport(withCtx(t), env.histogramRequest())
	requireOK(t, result)

	if !result.Body.Status {
		t.Fatalf("expected body status true, got: %s", result.JSON().String())
	}
	if len(result.Body.Payload.Overall.Progressive) == 0 && len(result.Body.Payload.Overall.Streaming) == 0 {
		t.Fatalf("expected non-empty overall histogram, got: %s", result.JSON().String())
	}

	points := result.Body.Payload.Overall.Progressive
	if len(points) == 0 {
		points = result.Body.Payload.Overall.Streaming
	}
	if _, err := points[0].Bandwidth.Float64(); err != nil {
		t.Fatalf("expected bandwidth to parse as float: %s", err)
	}
	t.Logf("bandwidth_report point: %+v", points[0])
}

func TestCacheReports(t *testing.T) {
	env := newTestEnv(t)

	t.Run("cache", func(t *testing.T) {
		result := env.mn.Analytics.ResourceAnalytics.GetCacheReport(withCtx(t), env.histogramRequest())
		requireOK(t, result)

		if !result.Body.Status {
			t.Fatalf("expected body status true, got: %s", result.JSON().String())
		}
		if len(result.Body.Payload.Overall) == 0 {
			t.Fatalf("expected non-empty overall series map, got: %s", result.JSON().String())
		}
		t.Logf("cache_report overall series: %d", len(result.Body.Payload.Overall))
	})

	t.Run("non_cache", func(t *testing.T) {
		result := env.mn.Analytics.ResourceAnalytics.GetNonCacheReport(withCtx(t), env.histogramRequest())
		requireOK(t, result)

		if !result.Body.Status {
			t.Fatalf("expected body status true, got: %s", result.JSON().String())
		}
		t.Logf("non_cache_report overall series: %d", len(result.Body.Payload.Overall))
	})
}

func TestCacheTrafficReports(t *testing.T) {
	env := newTestEnv(t)

	t.Run("cache_traffic", func(t *testing.T) {
		result := env.mn.Analytics.ResourceAnalytics.GetCacheTrafficReport(withCtx(t), env.histogramRequest())
		requireOK(t, result)

		if !result.Body.Status {
			t.Fatalf("expected body status true, got: %s", result.JSON().String())
		}
		if len(result.Body.Payload.Overall) == 0 {
			t.Fatalf("expected non-empty overall series map, got: %s", result.JSON().String())
		}
		t.Logf("cache_traffic_report overall series: %d", len(result.Body.Payload.Overall))
	})

	t.Run("non_cache_traffic", func(t *testing.T) {
		result := env.mn.Analytics.ResourceAnalytics.GetNonCacheTrafficReport(withCtx(t), env.histogramRequest())
		requireOK(t, result)

		if !result.Body.Status {
			t.Fatalf("expected body status true, got: %s", result.JSON().String())
		}
		t.Logf("non_cache_traffic_report overall series: %d", len(result.Body.Payload.Overall))
	})
}

func TestStatusReports(t *testing.T) {
	env := newTestEnv(t)

	classes := []mnresource.StatusClass{mnresource.Status2xx, mnresource.Status3xx, mnresource.Status4xx, mnresource.Status5xx}

	for _, class := range classes {
		t.Run(string(class), func(t *testing.T) {
			result := env.mn.Analytics.ResourceAnalytics.GetStatusReport(withCtx(t), class, env.histogramRequest())
			requireOK(t, result)

			if !result.Body.Status {
				t.Fatalf("expected body status true, got: %s", result.JSON().String())
			}
			t.Logf("status_report[%s] overall series: %d", class, len(result.Body.Payload.Overall))
		})
	}
}

func TestStatusReportInvalidClass(t *testing.T) {
	env := newTestEnv(t)

	result := env.mn.Analytics.ResourceAnalytics.GetStatusReport(withCtx(t), "6xx", env.histogramRequest())
	if result.Error == nil {
		t.Fatal("expected validation error for invalid status class")
	}
}

func TestErrors(t *testing.T) {
	env := newTestEnv(t)

	result := env.mn.Analytics.ResourceAnalytics.GetErrors(withCtx(t), env.limitRequest())
	requireOK(t, result)

	if !result.Body.Status {
		t.Fatalf("expected body status true, got: %s", result.JSON().String())
	}
	t.Logf("errors overall: %+v", result.Body.Payload.Overall)
}

func TestVisitorsCountries(t *testing.T) {
	env := newTestEnv(t)

	result := env.mn.Analytics.ResourceAnalytics.GetVisitorsCountries(withCtx(t), env.limitRequest())
	requireOK(t, result)

	if !result.Body.Status {
		t.Fatalf("expected body status true, got: %s", result.JSON().String())
	}
	if len(result.Body.Payload.Overall) == 0 {
		t.Fatalf("expected non-empty overall countries, got: %s", result.JSON().String())
	}

	item := result.Body.Payload.Overall[0]
	if item.Country == "" {
		t.Fatalf("expected country name, got: %+v", item)
	}
	// traffic_gb comes back as a pre-formatted string from the API (e.g. "3,038.96")
	if item.TrafficGb.String() == "" {
		t.Fatalf("expected non-empty traffic_gb, got: %+v", item)
	}
	t.Logf("visitors_countries top: %+v", item)
}

func TestStatus(t *testing.T) {
	env := newTestEnv(t)

	result := env.mn.Analytics.ResourceAnalytics.GetStatus(withCtx(t), env.request())
	requireOK(t, result)

	if !result.Body.Status {
		t.Fatalf("expected body status true, got: %s", result.JSON().String())
	}
	if result.Body.Payload.Overall.SuccessfulResponses.TotalSuccessfulResponse == 0 {
		t.Fatalf("expected non-zero successful responses, got: %s", result.JSON().String())
	}
	t.Logf("status overall: %+v", result.Body.Payload.Overall)
}

func TestRegion(t *testing.T) {
	env := newTestEnv(t)

	result := env.mn.Analytics.ResourceAnalytics.GetRegion(withCtx(t), env.request())
	requireOK(t, result)

	if !result.Body.Status {
		t.Fatalf("expected body status true, got: %s", result.JSON().String())
	}
	if len(result.Body.Payload.Overall) == 0 {
		t.Fatalf("expected non-empty overall regions, got: %s", result.JSON().String())
	}
	t.Logf("region overall regions: %d", len(result.Body.Payload.Overall))
}

func TestActiveTraffic(t *testing.T) {
	env := newTestEnv(t)

	result := env.mn.Analytics.ResourceAnalytics.GetActiveTraffic(withCtx(t), env.histogramRequest())
	requireOK(t, result)

	if !result.Body.Status {
		t.Fatalf("expected body status true, got: %s", result.JSON().String())
	}
	if result.Body.Payload.Overall.Unit == "" {
		t.Fatalf("expected unit in overall active traffic, got: %s", result.JSON().String())
	}
	if _, err := result.Body.Payload.Overall.Current.Float64(); err != nil {
		t.Fatalf("expected current to parse as float: %s", err)
	}
	t.Logf("active_traffic overall: %+v", result.Body.Payload.Overall)
}

func TestActiveHit(t *testing.T) {
	env := newTestEnv(t)

	result := env.mn.Analytics.ResourceAnalytics.GetActiveHit(withCtx(t), env.histogramRequest())
	requireOK(t, result)

	if !result.Body.Status {
		t.Fatalf("expected body status true, got: %s", result.JSON().String())
	}
	if result.Body.Payload.Overall.Unit == "" {
		t.Fatalf("expected unit in overall active hit, got: %s", result.JSON().String())
	}
	t.Logf("active_hit overall: %+v", result.Body.Payload.Overall)
}

func TestTrafficDetail(t *testing.T) {
	env := newTestEnv(t)

	result := env.mn.Analytics.ResourceAnalytics.GetTrafficDetail(withCtx(t), env.request())
	requireOK(t, result)

	if !result.Body.Status {
		t.Fatalf("expected body status true, got: %s", result.JSON().String())
	}
	if len(result.Body.Payload.Overall.Traffic) == 0 {
		t.Fatalf("expected non-empty overall traffic periods, got: %s", result.JSON().String())
	}
	if _, err := result.Body.Payload.Overall.CachedData.TotalCachedTraffic.Traffic.Float64(); err != nil {
		t.Fatalf("expected total_cached_traffic to parse as float: %s", err)
	}
	t.Logf("traffic_detail overall: %+v", result.Body.Payload.Overall.CachedData)
}

func TestRequestDetail(t *testing.T) {
	env := newTestEnv(t)

	result := env.mn.Analytics.ResourceAnalytics.GetRequestDetail(withCtx(t), env.request())
	requireOK(t, result)

	if !result.Body.Status {
		t.Fatalf("expected body status true, got: %s", result.JSON().String())
	}
	if result.Body.Payload.Overall.TotalRequest == 0 {
		t.Fatalf("expected non-zero overall total_request, got: %s", result.JSON().String())
	}
	if result.Body.Payload.Overall.CachedData.TotalHits+result.Body.Payload.Overall.NonCachedData.TotalMisses == 0 {
		t.Fatalf("expected non-zero hits or misses, got: %s", result.JSON().String())
	}
	t.Logf("request_detail overall: %+v", result.Body.Payload.Overall)
}
