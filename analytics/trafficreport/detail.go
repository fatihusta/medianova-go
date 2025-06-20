package trafficreport

import (
	"context"
	"net/http"
	"path"

	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/client/utils"
	"github.com/fatihusta/medianova-go/common"
)

type TrafficReportDetailService struct {
	request *request.RequestConfig
}

func NewTrafficReportDetailService(reqCfg *request.RequestConfig) *TrafficReportDetailService {
	return &TrafficReportDetailService{request: reqCfg}
}

func (s *TrafficReportDetailService) Get(ctx context.Context, reportRequest TrafficReportDetailRequest) *common.Result[TrafficReportDetailResponse] {

	result := common.NewResult[TrafficReportDetailResponse]()
	body, err := utils.ToJSONBodyBuffer(reportRequest)
	if err != nil {
		result.Error = err
		return result
	}

	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "traffic_report", "detail")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		result.Error = err
		return result
	}

	return utils.DoHTTPRequest[TrafficReportDetailResponse](s.request.GetClient(), req)
}
