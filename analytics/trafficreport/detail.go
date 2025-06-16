package trafficreport

import (
	"context"
	"net/http"
	"path"

	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/client/utils"
)

type TrafficReportDetailService struct {
	request *request.RequestConfig
}

func NewTrafficReportDetailService(reqCfg *request.RequestConfig) *TrafficReportDetailService {
	return &TrafficReportDetailService{request: reqCfg}
}

func (s *TrafficReportDetailService) Get(ctx context.Context, reportRequest TrafficReportDetailRequest) (*TrafficReportDetailResponse, error) {

	body, err := utils.ToJSONBodyBuffer(reportRequest)
	if err != nil {
		return &TrafficReportDetailResponse{}, err
	}

	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "traffic_report", "detail")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		return &TrafficReportDetailResponse{}, err
	}

	return utils.DoHTTPRequest[*TrafficReportDetailResponse](s.request.GetClient(), req)
}
