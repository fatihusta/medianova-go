package statusreport

import (
	"context"
	"net/http"
	"path"

	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/client/utils"
	"github.com/fatihusta/medianova-go/common"
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
	Status3xx StatusCode = "3xx"
	Status4xx StatusCode = "4xx"
	Status5xx StatusCode = "5xx"
)

func (s *StatusReportService) Get(ctx context.Context, status_code StatusCode, reportRequest StatusReportRequest) *common.Result[StatusReportResponse] {

	result := common.NewResult[StatusReportResponse]()

	body, err := utils.ToJSONBodyBuffer(reportRequest)
	if err != nil {
		result.Error = err
		return result
	}

	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "status_report", "graph", string(status_code))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		result.Error = err
		return result
	}

	return utils.DoHTTPRequest[StatusReportResponse](s.request.GetClient(), req)
}
