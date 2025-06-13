package trafficreport

import (
	"context"
	"fmt"
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
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json")

	resp, err := s.request.Do(req)
	if err != nil {
		return &TrafficReportDetailResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg, err := utils.ToStringBody(resp)
		if err == nil {
			return &TrafficReportDetailResponse{}, fmt.Errorf("request not succeeded. status:%d, error:%s", resp.StatusCode, errMsg)
		}
		return &TrafficReportDetailResponse{}, fmt.Errorf("request not succeeded. status:%d", resp.StatusCode)
	}

	return utils.FromJSONToStruct[*TrafficReportDetailResponse](resp)
}
