package requestreport

import (
	"context"
	"net/http"
	"path"

	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/client/utils"
)

type RequestsReportDetailService struct {
	request *request.RequestConfig
}

func NewRequestsReportDetailService(reqCfg *request.RequestConfig) *RequestsReportDetailService {
	return &RequestsReportDetailService{request: reqCfg}
}

func (s *RequestsReportDetailService) Get(ctx context.Context, reportRequest RequestsReportDetailRequest) (*RequestsReportDetailResponse, error) {

	body, err := utils.ToJSONBodyBuffer(reportRequest)
	if err != nil {
		return &RequestsReportDetailResponse{}, err
	}

	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "requests_report", "detail")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		return &RequestsReportDetailResponse{}, err
	}

	return utils.DoHTTPRequest[*RequestsReportDetailResponse](s.request.GetClient(), req)
}
