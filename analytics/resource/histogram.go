package resource

import (
	"context"
	"fmt"

	"github.com/fatihusta/medianova-go/common"
)

func (s *ResourceService) GetTrafficReport(ctx context.Context, reportRequest ResourceHistogramRequest) *common.Result[TrafficHistogramResponse] {
	return doPost[TrafficHistogramResponse](s.request, ctx, reportRequest, "traffic_report", "graph")
}

func (s *ResourceService) GetRequestReport(ctx context.Context, reportRequest ResourceHistogramRequest) *common.Result[RequestHistogramResponse] {
	return doPost[RequestHistogramResponse](s.request, ctx, reportRequest, "request_report", "graph")
}

func (s *ResourceService) GetBandwidthReport(ctx context.Context, reportRequest ResourceHistogramRequest) *common.Result[BandwidthHistogramResponse] {
	return doPost[BandwidthHistogramResponse](s.request, ctx, reportRequest, "bandwidth_report", "graph")
}

func (s *ResourceService) GetCacheReport(ctx context.Context, reportRequest ResourceHistogramRequest) *common.Result[CacheHitHistogramResponse] {
	return doPost[CacheHitHistogramResponse](s.request, ctx, reportRequest, "cache_report", "graph")
}

func (s *ResourceService) GetNonCacheReport(ctx context.Context, reportRequest ResourceHistogramRequest) *common.Result[CacheHitHistogramResponse] {
	return doPost[CacheHitHistogramResponse](s.request, ctx, reportRequest, "non_cache_report", "graph")
}

func (s *ResourceService) GetCacheTrafficReport(ctx context.Context, reportRequest ResourceHistogramRequest) *common.Result[CacheTrafficHistogramResponse] {
	return doPost[CacheTrafficHistogramResponse](s.request, ctx, reportRequest, "cache_traffic_report", "graph")
}

func (s *ResourceService) GetNonCacheTrafficReport(ctx context.Context, reportRequest ResourceHistogramRequest) *common.Result[CacheTrafficHistogramResponse] {
	return doPost[CacheTrafficHistogramResponse](s.request, ctx, reportRequest, "non_cache_traffic_report", "graph")
}

func (s *ResourceService) GetStatusReport(ctx context.Context, statusClass StatusClass, reportRequest ResourceHistogramRequest) *common.Result[StatusHistogramResponse] {

	result := common.NewResult[StatusHistogramResponse]()

	switch statusClass {
	case Status2xx, Status3xx, Status4xx, Status5xx:
	default:
		result.Error = fmt.Errorf("invalid status class: %s, should be one of: 2xx, 3xx, 4xx, 5xx", statusClass)
		return result
	}

	return doPost[StatusHistogramResponse](s.request, ctx, reportRequest, "status_report", "graph", string(statusClass))
}
