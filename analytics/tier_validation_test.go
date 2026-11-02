package analytics_test

import (
	"context"
	"strings"
	"testing"

	mnerrorlogs "github.com/fatihusta/medianova-go/analytics/errorlogs"
	mnerrorreport "github.com/fatihusta/medianova-go/analytics/errorreport"
	mnoverview "github.com/fatihusta/medianova-go/analytics/overview"
	mnrequestreport "github.com/fatihusta/medianova-go/analytics/requestreport"
	mnstatusreport "github.com/fatihusta/medianova-go/analytics/statusreport"
	mntrafficreport "github.com/fatihusta/medianova-go/analytics/trafficreport"
	"github.com/fatihusta/medianova-go/client/request"
)

func TestV1InvalidTierRejectedClientSide(t *testing.T) {
	reqCfg := request.NewRequestConfig()
	ctx := context.Background()

	const (
		org     = "00000000-0000-0000-0000-000000000000"
		from    = "2024-11-12 05:00:00"
		to      = "2024-11-12 05:59:59"
		badTier = "mid"
	)

	check := func(name string, err error) {
		t.Helper()
		if err == nil || !strings.Contains(err.Error(), "tier") {
			t.Errorf("%s: expected tier validation error, got: %v", name, err)
		}
	}

	check("overview.historical.GetHit", mnoverview.NewHistoricalService(reqCfg).GetHit(ctx, mnoverview.HistoricalReportRequest{
		OrganizationUUID: org, Type: mnoverview.Hit, From: from, To: to, Tier: badTier,
	}).Error)
	check("overview.historical.GetHitRatio", mnoverview.NewHistoricalService(reqCfg).GetHitRatio(ctx, mnoverview.HistoricalReportRequest{
		OrganizationUUID: org, Type: mnoverview.HitRatio, From: from, To: to, Tier: badTier,
	}).Error)
	check("overview.historical.GetTraffic", mnoverview.NewHistoricalService(reqCfg).GetTraffic(ctx, mnoverview.HistoricalReportRequest{
		OrganizationUUID: org, Type: mnoverview.Traffic, From: from, To: to, Tier: badTier,
	}).Error)
	check("overview.top_resources", mnoverview.NewTopResourcesService(reqCfg).Get(ctx, mnoverview.TopResourcesRequest{
		OrganizationUUID: org, From: from, To: to, Tier: badTier,
	}).Error)
	check("overview.errors", mnoverview.NewErrorsService(reqCfg).Get(ctx, mnoverview.ErrorsRequest{
		OrganizationUUID: org, From: from, To: to, Tier: badTier,
	}).Error)
	check("overview.visitors_countries", mnoverview.NewVisitorsCountriesService(reqCfg).Get(ctx, mnoverview.VisitorsCountriesRequest{
		OrganizationUUID: org, From: from, To: to, Tier: badTier,
	}).Error)
	check("overview.active_hit_traffic", mnoverview.NewActiveHitTrafficService(reqCfg).Get(ctx, mnoverview.ActiveTypeHit, mnoverview.ActiveHitTrafficRequest{
		OrganizationUUID: org, From: from, To: to, Tier: badTier,
	}).Error)
	check("overview.region", mnoverview.NewRegionService(reqCfg).Get(ctx, mnoverview.RegionRequest{
		OrganizationUUID: org, From: from, To: to, Tier: badTier,
	}).Error)
	check("statusreport", mnstatusreport.NewStatusReportService(reqCfg).Get(ctx, mnstatusreport.Status2xx, mnstatusreport.StatusReportRequest{
		OrganizationUUID: org, From: from, To: to, Tier: badTier,
	}).Error)
	check("trafficreport", mntrafficreport.NewTrafficReportDetailService(reqCfg).Get(ctx, mntrafficreport.TrafficReportDetailRequest{
		OrganizationUUID: org, From: from, To: to, Tier: badTier,
	}).Error)
	check("requestreport", mnrequestreport.NewRequestsReportDetailService(reqCfg).Get(ctx, mnrequestreport.RequestsReportDetailRequest{
		OrganizationUUID: org, From: from, To: to, Tier: badTier,
	}).Error)
	check("errorlogs", mnerrorlogs.NewErrorLogsStatusCodesService(reqCfg).Get(ctx, 1, mnerrorlogs.ErrorLogsStatusCodesRequest{
		OrganizationUUID: org, From: from, To: to, StatusCodes: []int{404}, Tier: badTier,
	}).Error)
	check("errorreport", mnerrorreport.NewErrorReportStatusCodesService(reqCfg).Get(ctx, mnerrorreport.ErrorReportStatusCodesRequest{
		OrganizationUUID: org, From: from, To: to, Tier: badTier,
	}).Error)
}
