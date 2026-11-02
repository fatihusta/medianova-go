package stook

import (
	"context"
	"net/http"
	"path"

	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/client/utils"
	"github.com/fatihusta/medianova-go/common"
)

type StookService struct {
	request *request.RequestConfig
}

func NewStookService(reqCfg *request.RequestConfig) *StookService {
	return &StookService{request: reqCfg}
}

func doPost[T any](reqCfg *request.RequestConfig, ctx context.Context, reportRequest any, elem ...string) *common.Result[T] {
	result := common.NewResult[T]()

	if v, ok := reportRequest.(interface{ validate() error }); ok {
		if err := v.validate(); err != nil {
			result.Error = err
			return result
		}
	}

	body, err := utils.ToJSONBodyBuffer(reportRequest)
	if err != nil {
		result.Error = err
		return result
	}

	url := *reqCfg.BaseURLV2
	url.Path = path.Join(append([]string{url.Path, "analytics", "stook"}, elem...)...)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		result.Error = err
		return result
	}

	return utils.DoHTTPRequest[T](reqCfg.GetClient(), req)
}

func (s *StookService) GetOverview(ctx context.Context, reportRequest StookOverviewRequest) *common.Result[OverviewResponse] {
	return doPost[OverviewResponse](s.request, ctx, reportRequest, "overview")
}

func (s *StookService) GetFilters(ctx context.Context, reportRequest StookRequest) *common.Result[FiltersResponse] {
	return doPost[FiltersResponse](s.request, ctx, reportRequest, "filters")
}

func (s *StookService) GetRequestHistogram(ctx context.Context, reportRequest StookHistogramRequest) *common.Result[RequestHistogramResponse] {
	return doPost[RequestHistogramResponse](s.request, ctx, reportRequest, "request_histogram")
}

func (s *StookService) GetErrorRatioHistogram(ctx context.Context, reportRequest StookHistogramRequest) *common.Result[ErrorRatioHistogramResponse] {
	return doPost[ErrorRatioHistogramResponse](s.request, ctx, reportRequest, "error_ratio_histogram")
}

func (s *StookService) GetCharts(ctx context.Context, reportRequest StookChartsRequest) *common.Result[ChartsResponse] {
	return doPost[ChartsResponse](s.request, ctx, reportRequest, "charts")
}

func (s *StookService) GetBucketDetails(ctx context.Context, reportRequest StookRequest) *common.Result[BucketDetailsResponse] {
	return doPost[BucketDetailsResponse](s.request, ctx, reportRequest, "bucket_details")
}

func (s *StookService) GetFailureRatePerBucket(ctx context.Context, reportRequest StookHistogramRequest) *common.Result[FailureRatePerBucketResponse] {
	return doPost[FailureRatePerBucketResponse](s.request, ctx, reportRequest, "failure_rate_per_bucket")
}

func (s *StookService) GetRawLogs(ctx context.Context, reportRequest StookRawLogsRequest) *common.Result[RawLogsResponse] {
	return doPost[RawLogsResponse](s.request, ctx, reportRequest, "raw_logs")
}
