package overview

import (
	"context"
	"fmt"
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
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json")

	resp, err := s.request.Do(req)
	if err != nil {
		return &ErrorsResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg, err := utils.ToStringBody(resp)
		if err == nil {
			return &ErrorsResponse{}, fmt.Errorf("request not succeeded. status:%d, error:%s", resp.StatusCode, errMsg)
		}
		return &ErrorsResponse{}, fmt.Errorf("request not succeeded. status:%d", resp.StatusCode)
	}

	return utils.FromJSONToStruct[*ErrorsResponse](resp)
}
