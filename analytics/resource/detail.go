package resource

import (
	"context"

	"github.com/fatihusta/medianova-go/common"
)

func (s *ResourceService) GetTrafficDetail(ctx context.Context, reportRequest ResourceRequest) *common.Result[TrafficDetailResponse] {
	return doPost[TrafficDetailResponse](s.request, ctx, reportRequest, "traffic_detail")
}

func (s *ResourceService) GetRequestDetail(ctx context.Context, reportRequest ResourceRequest) *common.Result[RequestDetailResponse] {
	return doPost[RequestDetailResponse](s.request, ctx, reportRequest, "request_detail")
}
