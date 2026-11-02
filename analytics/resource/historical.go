package resource

import (
	"context"
	"fmt"

	"github.com/fatihusta/medianova-go/common"
)

// GetHistorical retrieves historical analytics for hit, traffic, bandwidth and requests metrics.
// For hit_ratio use GetHistoricalHitRatio, for gpu_transaction use GetHistoricalGPUTransaction,
// since the API returns a different payload shape for those types.
func (s *ResourceService) GetHistorical(ctx context.Context, reportRequest ResourceHistoricalRequest) *common.Result[HistoricalResponse] {

	result := common.NewResult[HistoricalResponse]()

	switch reportRequest.Type {
	case Hit, Traffic, Bandwidth, Requests:
	case HitRatio:
		result.Error = fmt.Errorf("type hit_ratio has a different response shape, use GetHistoricalHitRatio")
		return result
	case GPUTransaction:
		result.Error = fmt.Errorf("type gpu_transaction has a different response shape, use GetHistoricalGPUTransaction")
		return result
	default:
		result.Error = fmt.Errorf("type is required, one of: hit, traffic, bandwidth, requests")
		return result
	}

	return doPost[HistoricalResponse](s.request, ctx, reportRequest, "historical")
}

// GetHistoricalHitRatio retrieves historical hit ratio for the given time range.
func (s *ResourceService) GetHistoricalHitRatio(ctx context.Context, reportRequest ResourceRequest) *common.Result[HistoricalHitRatioResponse] {
	return doPost[HistoricalHitRatioResponse](s.request, ctx,
		ResourceHistoricalRequest{ResourceRequest: reportRequest, Type: HitRatio}, "historical")
}

// GetHistoricalGPUTransaction retrieves historical gpu transaction count for the given time range.
// Note: the API ignores the resources filter for this type and returns all resources of the organization.
func (s *ResourceService) GetHistoricalGPUTransaction(ctx context.Context, reportRequest ResourceRequest) *common.Result[HistoricalGPUTransactionResponse] {
	return doPost[HistoricalGPUTransactionResponse](s.request, ctx,
		ResourceHistoricalRequest{ResourceRequest: reportRequest, Type: GPUTransaction}, "historical")
}
