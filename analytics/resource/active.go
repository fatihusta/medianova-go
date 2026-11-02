package resource

import (
	"context"

	"github.com/fatihusta/medianova-go/common"
)

func (s *ResourceService) GetActiveTraffic(ctx context.Context, reportRequest ResourceHistogramRequest) *common.Result[ActiveDataResponse] {
	return doPost[ActiveDataResponse](s.request, ctx, reportRequest, "active_traffic")
}

func (s *ResourceService) GetActiveHit(ctx context.Context, reportRequest ResourceHistogramRequest) *common.Result[ActiveDataResponse] {
	return doPost[ActiveDataResponse](s.request, ctx, reportRequest, "active_hit")
}
