package stook_test

import (
	"context"
	"os"
	"sync"
	"testing"
	"time"

	mnstook "github.com/fatihusta/medianova-go/analytics/stook"
	"github.com/fatihusta/medianova-go/client"
	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/common"
)

const dateLayout = "2006-01-02 15:04:05"

func newHTTPClient(auth client.Middleware) *client.Client {
	reqConfig := request.NewRequestConfig()
	reqConfig.RequestTimeout = 30 * time.Second

	middlewares := client.WithMiddlewares(
		// the stook analytics backend is intermittently slow to resolve
		// organization buckets (HTTP 500); retry a few times
		client.RetryMiddleware(4, 3*time.Second),
		auth,
	)

	return client.NewClient(reqConfig, middlewares)
}

// stookAuth prefers the dashboard session token: the stook analytics
// endpoints currently answer HTTP 500 "Failed to fetch organization buckets"
// for organization API tokens and HTTP 200 for session JWTs.
func stookAuth() (client.Middleware, bool) {
	if session := os.Getenv("MEDIANOVA_SESSION_TOKEN"); session != "" {
		return client.AuthWithBearer(session), true
	}
	if org := os.Getenv("MEDIANOVA_ORG_TOKEN"); org != "" {
		return client.AuthWithOrganizationToken(org), true
	}
	return nil, false
}

type testEnv struct {
	stook     *mnstook.StookService
	orgUUID   string
	from      string
	to        string
	available bool
	unavailAt int
}

var (
	resolveOnce sync.Once
	resolvedEnv testEnv
)

// newTestEnv resolves the organization and probes once whether the stook
// analytics routes answer for the configured token. Tests are skipped when
// the probe keeps failing (organization API tokens currently get HTTP 500
// "Failed to fetch organization buckets"; dashboard session JWTs work).
func newTestEnv(t *testing.T) testEnv {
	t.Helper()

	auth, ok := stookAuth()
	if !ok {
		t.Skip("MEDIANOVA_SESSION_TOKEN or MEDIANOVA_ORG_TOKEN is not set, skipping live API test")
	}

	resolveOnce.Do(func() {
		mn := newHTTPClient(auth)

		env := testEnv{}
		from := time.Now().Add(-26 * time.Hour).Truncate(time.Hour)
		env.from = from.Format(dateLayout)
		env.to = from.Add(time.Hour).Format(dateLayout)

		env.orgUUID = os.Getenv("MEDIANOVA_ORG_UUID")
		if env.orgUUID == "" {
			organizations := mn.Organization.List()
			if organizations.Error != nil || len(organizations.Body) == 0 {
				return
			}
			env.orgUUID = organizations.Body[0].UUID
		}

		env.stook = mn.Analytics.Stook
		probe := newHTTPClient(auth).Analytics.Stook.GetFilters(context.Background(), env.request())
		env.available = probe.Error == nil && probe.Status == 200
		env.unavailAt = probe.Status
		resolvedEnv = env
	})

	env := resolvedEnv
	env.stook = newHTTPClient(auth).Analytics.Stook
	if !env.available {
		t.Skipf("stook analytics is not available for this organization (probe HTTP %d), skipping live API test", env.unavailAt)
	}
	return env
}

func (e testEnv) request() mnstook.StookRequest {
	return mnstook.StookRequest{
		OrganizationUUID: e.orgUUID,
		From:             e.from,
		To:               e.to,
	}
}

func (e testEnv) histogramRequest() mnstook.StookHistogramRequest {
	return mnstook.StookHistogramRequest{StookRequest: e.request(), Interval: mnstook.Interval5m}
}

func requireOK[T any](t *testing.T, result *common.Result[T]) {
	t.Helper()
	if result.Error != nil {
		t.Fatalf("unexpected error: %v", result.Error)
	}
	if result.Status != 200 {
		t.Fatalf("expected status 200, got %d: %s", result.Status, result.JSON().String())
	}
}

func TestFilters(t *testing.T) {
	env := newTestEnv(t)

	result := env.stook.GetFilters(context.Background(), env.request())
	requireOK(t, result)

	if !result.Body.Status {
		t.Fatalf("expected body status true, got: %s", result.JSON().String())
	}
	t.Logf("filters: %d fields", len(result.Body.Payload))
	for _, f := range result.Body.Payload {
		t.Logf("  field=%s type=%s operators=%v", f.Field, f.Type, f.Operators)
	}
}

func TestOverview(t *testing.T) {
	env := newTestEnv(t)

	result := env.stook.GetOverview(context.Background(), mnstook.StookOverviewRequest{StookRequest: env.request()})
	requireOK(t, result)

	if !result.Body.Status {
		t.Fatalf("expected body status true, got: %s", result.JSON().String())
	}
	t.Logf("overview requests=%s (%s) success=%s stored=%s download=%s (%s %s)",
		result.Body.Payload.TotalRequests.String(), result.Body.Payload.TotalRequestsFormatted,
		result.Body.Payload.SuccessRateFormatted, result.Body.Payload.TotalStoredDataFormatted,
		result.Body.Payload.Download.Value.String(), result.Body.Payload.Download.Formatted, result.Body.Payload.Download.Unit)
}

func TestRequestHistogram(t *testing.T) {
	env := newTestEnv(t)

	result := env.stook.GetRequestHistogram(context.Background(), env.histogramRequest())
	requireOK(t, result)

	if !result.Body.Status {
		t.Fatalf("expected body status true, got: %s", result.JSON().String())
	}
	t.Logf("request_histogram interval=%s total=%d points=%d", result.Body.Payload.Interval,
		result.Body.Payload.TotalRequests, len(result.Body.Payload.Data))
}

func TestErrorRatioHistogram(t *testing.T) {
	env := newTestEnv(t)

	result := env.stook.GetErrorRatioHistogram(context.Background(), env.histogramRequest())
	requireOK(t, result)

	if !result.Body.Status {
		t.Fatalf("expected body status true, got: %s", result.JSON().String())
	}
	t.Logf("error_ratio_histogram interval=%s points=%d", result.Body.Payload.Interval, len(result.Body.Payload.Data))
}

func TestCharts(t *testing.T) {
	env := newTestEnv(t)

	cases := map[string]mnstook.StookChartsRequest{
		"metric": {
			StookRequest: env.request(),
			ChartType:    mnstook.ChartTypeMetric,
			Metric:       mnstook.MetricRequests,
			Interval:     mnstook.Interval1h,
		},
		"per_bucket": {
			StookRequest: env.request(),
			ChartType:    mnstook.ChartTypePerBucket,
			Metric:       mnstook.MetricBytes,
			Interval:     mnstook.Interval1h,
		},
	}

	for name, req := range cases {
		t.Run(name, func(t *testing.T) {
			result := env.stook.GetCharts(context.Background(), req)
			requireOK(t, result)

			if !result.Body.Status {
				t.Fatalf("expected body status true, got: %s", result.JSON().String())
			}
			t.Logf("charts[%s] chart_type=%s metric=%s series=%d", name,
				result.Body.Payload.ChartType, result.Body.Payload.Metric, len(result.Body.Payload.Series))
		})
	}
}

func TestBucketDetails(t *testing.T) {
	env := newTestEnv(t)

	result := env.stook.GetBucketDetails(context.Background(), env.request())
	requireOK(t, result)

	if !result.Body.Status {
		t.Fatalf("expected body status true, got: %s", result.JSON().String())
	}
	t.Logf("bucket_details buckets=%d", len(result.Body.Payload.Buckets))
}

func TestFailureRatePerBucket(t *testing.T) {
	env := newTestEnv(t)

	result := env.stook.GetFailureRatePerBucket(context.Background(), env.histogramRequest())
	requireOK(t, result)

	if !result.Body.Status {
		t.Fatalf("expected body status true, got: %s", result.JSON().String())
	}
	t.Logf("failure_rate_per_bucket interval=%s series=%d", result.Body.Payload.Interval, len(result.Body.Payload.Series))
}

func TestRawLogs(t *testing.T) {
	env := newTestEnv(t)

	limit := 5
	page := 1
	result := env.stook.GetRawLogs(context.Background(), mnstook.StookRawLogsRequest{
		StookRequest: env.request(),
		Page:         &page,
		Limit:        &limit,
		SortField:    "ts",
		SortOrder:    mnstook.SortOrderDesc,
	})
	requireOK(t, result)

	if !result.Body.Status {
		t.Fatalf("expected body status true, got: %s", result.JSON().String())
	}
	t.Logf("raw_logs total=%d page=%d/%d logs=%d", result.Body.Payload.TotalCount,
		result.Body.Payload.CurrentPage, result.Body.Payload.TotalPage, len(result.Body.Payload.Logs))
}
