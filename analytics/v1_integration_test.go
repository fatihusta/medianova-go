package analytics_test

import (
	"context"
	"os"
	"testing"
	"time"

	mnerrorlogs "github.com/fatihusta/medianova-go/analytics/errorlogs"
	mnerrorreport "github.com/fatihusta/medianova-go/analytics/errorreport"
	mnoverview "github.com/fatihusta/medianova-go/analytics/overview"
	mnrequestreport "github.com/fatihusta/medianova-go/analytics/requestreport"
	mnstatusreport "github.com/fatihusta/medianova-go/analytics/statusreport"
	mntier "github.com/fatihusta/medianova-go/analytics/tier"
	mntrafficreport "github.com/fatihusta/medianova-go/analytics/trafficreport"
	"github.com/fatihusta/medianova-go/client"
	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/common"
)

type v1Env struct {
	mn      *client.Client
	ctx     context.Context
	orgUUID string
	cdnURL  string
	from    string
	to      string
}

func newV1Env(t *testing.T) v1Env {
	t.Helper()

	token := os.Getenv("MEDIANOVA_ORG_TOKEN")
	if token == "" {
		t.Skip("MEDIANOVA_ORG_TOKEN is not set, skipping live API test")
	}

	mn := client.NewClient(request.NewRequestConfig(),
		client.WithMiddlewares(client.AuthWithOrganizationToken(token)))

	env := v1Env{mn: mn, ctx: context.Background()}

	env.orgUUID = os.Getenv("MEDIANOVA_ORG_UUID")
	if env.orgUUID == "" {
		list := mn.Organization.List()
		if list.Error != nil || len(list.Body) == 0 {
			t.Fatalf("could not resolve organization: %v", list.Error)
		}
		env.orgUUID = list.Body[0].UUID
	}

	env.cdnURL = os.Getenv("MEDIANOVA_RESOURCE_CDN_URL")
	if env.cdnURL == "" {
		resources := mn.CDN.Resource.List(env.orgUUID)
		if resources.Error != nil || len(resources.Body) == 0 {
			t.Fatalf("could not resolve cdn resource: %v", resources.Error)
		}
		env.cdnURL = resources.Body[0].CdnURL
	}

	// one full hour within the last 26h, v1 format "YYYY-MM-DD HH:MM:SS"
	from := time.Now().Add(-26 * time.Hour).Truncate(time.Hour)
	env.from = from.Format("2006-01-02 15:04:05")
	env.to = from.Add(time.Hour).Format("2006-01-02 15:04:05")

	return env
}

func requireOK[T any](t *testing.T, res *common.Result[T]) T {
	t.Helper()
	if res.Error != nil {
		t.Fatalf("unexpected error: %v", res.Error)
	}
	if res.Status != 200 {
		t.Fatalf("status=%d, body=%s", res.Status, res.JSON())
	}
	return res.Body
}

func TestLiveHistorical(t *testing.T) {
	env := newV1Env(t)
	svc := env.mn.Analytics.Historical

	hit := requireOK(t, svc.GetHit(env.ctx, mnoverview.HistoricalReportRequest{
		OrganizationUUID: env.orgUUID, Type: mnoverview.Hit, From: env.from, To: env.to,
		Resources: []string{env.cdnURL},
	}))
	t.Logf("hit selected_time=%d (%s)", hit.Payload.SelectedTime.Hit, hit.Payload.SelectedTime.HitFormatted)

	bandwidth := requireOK(t, svc.GetBandwidth(env.ctx, mnoverview.HistoricalReportRequest{
		OrganizationUUID: env.orgUUID, Type: mnoverview.Bandwidth, From: env.from, To: env.to,
		Resources: []string{env.cdnURL},
	}))
	t.Logf("bandwidth=%s", bandwidth.Payload.SelectedTime.BandwidthFormatted)

	requests := requireOK(t, svc.GetRequests(env.ctx, mnoverview.HistoricalReportRequest{
		OrganizationUUID: env.orgUUID, Type: mnoverview.Requests, From: env.from, To: env.to,
		Resources: []string{env.cdnURL}, CalcType: mnoverview.CalcTypeRequestNumber,
	}))
	t.Logf("requests=%d (%s)", requests.Payload.SelectedTime.Count, requests.Payload.SelectedTime.CountFormatted)

	edge := requireOK(t, svc.GetHit(env.ctx, mnoverview.HistoricalReportRequest{
		OrganizationUUID: env.orgUUID, Type: mnoverview.Hit, From: env.from, To: env.to,
		Resources: []string{env.cdnURL}, Tier: mntier.TierEdge,
	}))
	t.Logf("hit(tier=edge) selected_time=%d", edge.Payload.SelectedTime.Hit)
}

func TestLiveOverviewReports(t *testing.T) {
	env := newV1Env(t)

	errs := requireOK(t, env.mn.Analytics.Errors.Get(env.ctx, mnoverview.ErrorsRequest{
		OrganizationUUID: env.orgUUID, From: env.from, To: env.to, Resources: []string{env.cdnURL},
	}))
	t.Logf("errors=%d status codes", len(errs.Payload))

	region := requireOK(t, env.mn.Analytics.Region.Get(env.ctx, mnoverview.RegionRequest{
		OrganizationUUID: env.orgUUID, From: env.from, To: env.to, Resources: []string{env.cdnURL},
	}))
	t.Logf("region EU hit=%d MEAA hit=%d", region.Payload.Eu.TotalHit, region.Payload.Meaa.TotalHit)

	top := requireOK(t, env.mn.Analytics.TopResources.Get(env.ctx, mnoverview.TopResourcesRequest{
		OrganizationUUID: env.orgUUID, From: env.from, To: env.to, Resources: []string{env.cdnURL},
	}))
	_ = top

	visitors := requireOK(t, env.mn.Analytics.VistorCountries.Get(env.ctx, mnoverview.VisitorsCountriesRequest{
		OrganizationUUID: env.orgUUID, From: env.from, To: env.to, Resources: []string{env.cdnURL},
	}))
	t.Logf("visitors countries total=%d", visitors.Payload.Total)

	active := requireOK(t, env.mn.Analytics.ActiveHitTraffic.Get(env.ctx, mnoverview.ActiveTypeHit,
		mnoverview.ActiveHitTrafficRequest{
			OrganizationUUID: env.orgUUID, From: env.from, To: env.to, Resources: []string{env.cdnURL},
		}))
	t.Logf("active hit current=%d max=%d %s", active.Payload.Current, active.Payload.Max, active.Payload.Unit)
}

func TestLiveReports(t *testing.T) {
	env := newV1Env(t)

	status := requireOK(t, env.mn.Analytics.StatusReport.Get(env.ctx, mnstatusreport.Status2xx,
		mnstatusreport.StatusReportRequest{
			OrganizationUUID: env.orgUUID, From: env.from, To: env.to, Resources: []string{env.cdnURL},
			Interval: "1h",
		}))
	if !status.Status {
		t.Errorf("status report not ok: %+v", status)
	}

	errReport := requireOK(t, env.mn.Analytics.ErrorReportStatusCodes.Get(env.ctx,
		mnerrorreport.ErrorReportStatusCodesRequest{
			OrganizationUUID: env.orgUUID, From: env.from, To: env.to, Resources: []string{env.cdnURL},
		}))
	t.Logf("error report total=%d client_errors=%d server_errors=%d",
		errReport.Payload.TotalRequest, errReport.Payload.ClientErrorResponses.TotalClientErrorResponses,
		errReport.Payload.ServerErrorResponses.TotalServerErrorResponses)

	traffic := requireOK(t, env.mn.Analytics.TrafficReportDetail.Get(env.ctx,
		mntrafficreport.TrafficReportDetailRequest{
			OrganizationUUID: env.orgUUID, From: env.from, To: env.to, Resources: []string{env.cdnURL},
		}))
	t.Logf("cached traffic=%s", traffic.Payload.CachedData.TotalCachedTraffic.TrafficFormatted)

	reqs := requireOK(t, env.mn.Analytics.RequestsReportDetail.Get(env.ctx,
		mnrequestreport.RequestsReportDetailRequest{
			OrganizationUUID: env.orgUUID, From: env.from, To: env.to, Resources: []string{env.cdnURL},
		}))
	_ = reqs
}

func TestLiveErrorLogs(t *testing.T) {
	env := newV1Env(t)

	logs := requireOK(t, env.mn.Analytics.ErrorLogsStatusCodes.Get(env.ctx, 1,
		mnerrorlogs.ErrorLogsStatusCodesRequest{
			OrganizationUUID: env.orgUUID, From: env.from, To: env.to,
			Resources:   []string{env.cdnURL},
			StatusCodes: []int{404},
			PageSize:    5,
		}))
	if !logs.Status {
		t.Errorf("error logs not ok: %+v", logs)
	}
}
