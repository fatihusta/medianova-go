package errorlogs

import (
	"context"
	"net/http"
	"path"
	"strconv"

	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/client/utils"
)

type ErrorLogsStatusCodesService struct {
	request *request.RequestConfig
}

func NewErrorLogsStatusCodesService(reqCfg *request.RequestConfig) *ErrorLogsStatusCodesService {
	return &ErrorLogsStatusCodesService{request: reqCfg}
}

var errorLogsStatusCodesDefault = []int{400, 401, 403, 429, 500, 502, 503, 504}

func (s *ErrorLogsStatusCodesService) Get(ctx context.Context, page int, reportRequest ErrorLogsStatusCodesRequest) (*ErrorLogsStatusCodesResponse, error) {

	if len(reportRequest.StatusCodes) < 1 {
		reportRequest.StatusCodes = errorLogsStatusCodesDefault
	}

	body, err := utils.ToJSONBodyBuffer(reportRequest)
	if err != nil {
		return &ErrorLogsStatusCodesResponse{}, err
	}

	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "error_logs", "status_codes")
	q := url.Query()
	q.Set("page", strconv.Itoa(page))
	url.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		return &ErrorLogsStatusCodesResponse{}, err
	}

	return utils.DoHTTPRequest[*ErrorLogsStatusCodesResponse](s.request.GetClient(), req)
}
