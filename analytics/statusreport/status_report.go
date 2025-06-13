package statusreport

import (
	"context"
	"fmt"
	"net/http"
	"path"

	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/client/utils"
)

type StatusReportService struct {
	request *request.RequestConfig
}

func NewStatusReportService(reqCfg *request.RequestConfig) *StatusReportService {
	return &StatusReportService{request: reqCfg}
}

// StatusCode enum values
const (
	Status2xx StatusCode = "2xx"
	Status4xx StatusCode = "4xx"
	Status5xx StatusCode = "5xx"
)

func (s *StatusReportService) Get(ctx context.Context, status_code StatusCode, reportRequest StatusReportRequest) (*StatusReportResponse, error) {

	body, err := utils.ToJSONBodyBuffer(reportRequest)
	if err != nil {
		return &StatusReportResponse{}, err
	}

	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "status_report", "graph", string(status_code))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		return &StatusReportResponse{}, err
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json")

	resp, err := s.request.Do(req)
	if err != nil {
		return &StatusReportResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg, err := utils.ToStringBody(resp)
		if err == nil {
			return &StatusReportResponse{}, fmt.Errorf("request not succeeded. status:%d, error:%s", resp.StatusCode, errMsg)
		}
		return &StatusReportResponse{}, fmt.Errorf("request not succeeded. status:%d", resp.StatusCode)
	}

	return utils.FromJSONToStruct[*StatusReportResponse](resp)
}
