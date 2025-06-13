package requestreport

import (
	"context"
	"fmt"
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
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json")

	resp, err := s.request.Do(req)
	if err != nil {
		return &RequestsReportDetailResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg, err := utils.ToStringBody(resp)
		if err == nil {
			return &RequestsReportDetailResponse{}, fmt.Errorf("request not succeeded. status:%d, error:%s", resp.StatusCode, errMsg)
		}
		return &RequestsReportDetailResponse{}, fmt.Errorf("request not succeeded. status:%d", resp.StatusCode)
	}

	return utils.FromJSONToStruct[*RequestsReportDetailResponse](resp)
}
