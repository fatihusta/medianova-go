package analytics

import (
	"github.com/fatihusta/medianova-go/analytics/errorlogs"
	"github.com/fatihusta/medianova-go/analytics/errorreport"
	"github.com/fatihusta/medianova-go/analytics/overview"
	"github.com/fatihusta/medianova-go/analytics/requestreport"
	"github.com/fatihusta/medianova-go/analytics/statusreport"
	"github.com/fatihusta/medianova-go/analytics/trafficreport"
	"github.com/fatihusta/medianova-go/client/request"
)

type Analytics struct {
	Historical             *overview.HistoricalService
	TopResources           *overview.TopResourcesService
	Errors                 *overview.ErrorsService
	VistorCountries        *overview.VisitorsCountriesService
	ActiveHitTraffic       *overview.ActiveHitTrafficService
	Region                 *overview.RegionService
	StatusReport           *statusreport.StatusReportService
	TrafficReportDetail    *trafficreport.TrafficReportDetailService
	RequestsReportDetail   *requestreport.RequestsReportDetailService
	ErrorLogsStatusCodes   *errorlogs.ErrorLogsStatusCodesService
	ErrorReportStatusCodes *errorreport.ErrorReportStatusCodesService
}

func NewAnalytics(reqCfg *request.RequestConfig) *Analytics {
	return &Analytics{
		Historical:             overview.NewHistoricalService(reqCfg),
		TopResources:           overview.NewTopResourcesService(reqCfg),
		Errors:                 overview.NewErrorsService(reqCfg),
		VistorCountries:        overview.NewVisitorsCountriesService(reqCfg),
		ActiveHitTraffic:       overview.NewActiveHitTrafficService(reqCfg),
		Region:                 overview.NewRegionService(reqCfg),
		StatusReport:           statusreport.NewStatusReportService(reqCfg),
		TrafficReportDetail:    trafficreport.NewTrafficReportDetailService(reqCfg),
		RequestsReportDetail:   requestreport.NewRequestsReportDetailService(reqCfg),
		ErrorLogsStatusCodes:   errorlogs.NewErrorLogsStatusCodesService(reqCfg),
		ErrorReportStatusCodes: errorreport.NewErrorReportStatusCodesService(reqCfg),
	}
}
