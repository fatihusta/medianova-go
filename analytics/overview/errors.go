package overview

import (
	"context"
	"net/http"
	"path"

	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/client/utils"
)

type ErrorsService struct {
	request *request.RequestConfig
}

func NewErrorsService(reqCfg *request.RequestConfig) *ErrorsService {
	return &ErrorsService{request: reqCfg}
}

func (s *ErrorsService) Get(ctx context.Context, reportRequest ErrorsRequest) (*ErrorsResponse, error) {

	body, err := utils.ToJSONBodyBuffer(reportRequest)
	if err != nil {
		return &ErrorsResponse{}, err
	}

	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "overview_report", "errors")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		return &ErrorsResponse{}, err
	}

	return utils.DoHTTPRequest[*ErrorsResponse](s.request.GetClient(), req)
}
