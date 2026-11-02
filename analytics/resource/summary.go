package resource

import (
	"context"

	"github.com/fatihusta/medianova-go/common"
)

func (s *ResourceService) GetErrors(ctx context.Context, reportRequest ResourceLimitRequest) *common.Result[ErrorsResponse] {
	return doPost[ErrorsResponse](s.request, ctx, reportRequest, "errors")
}

func (s *ResourceService) GetVisitorsCountries(ctx context.Context, reportRequest ResourceLimitRequest) *common.Result[VisitorsCountriesResponse] {
	return doPost[VisitorsCountriesResponse](s.request, ctx, reportRequest, "visitors_countries")
}

func (s *ResourceService) GetStatus(ctx context.Context, reportRequest ResourceRequest) *common.Result[StatusResponse] {
	return doPost[StatusResponse](s.request, ctx, reportRequest, "status")
}

func (s *ResourceService) GetRegion(ctx context.Context, reportRequest ResourceRequest) *common.Result[RegionResponse] {
	return doPost[RegionResponse](s.request, ctx, reportRequest, "region")
}
