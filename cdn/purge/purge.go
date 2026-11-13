package purge

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/client/utils"
	"github.com/fatihusta/medianova-go/common"
)

type PurgeService struct {
	request *request.RequestConfig
}

func NewPurgeService(reqCfg *request.RequestConfig) *PurgeService {
	return &PurgeService{request: reqCfg}
}

func (s *PurgeService) Purge(organizationUUID string, p PurgeRequest) *common.Result[PurgeResponse] {

	result := common.NewResult[PurgeResponse]()

	url, err := utils.JoinPath(s.request.BaseURL, "cdn", organizationUUID, "job", p.ResourceUUID, "purge")
	if err != nil {
		result.Error = err
		return result
	}

	body, err := utils.ToJSONBodyBuffer(p)
	if err != nil {
		result.Error = err
		return result
	}
	ctx, cancel := context.WithTimeout(context.Background(), s.request.RequestTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		result.Error = err
		return result
	}

	return utils.DoHTTPRequest[PurgeResponse](s.request.GetClient(), req)
}

func (s *PurgeService) Status(r PurgeStatusRequest) *common.Result[[]PurgeStatusResponseRequests] {

	result := common.NewResult[[]PurgeStatusResponseRequests]()

	url, err := utils.JoinPath(s.request.BaseURL, "cdn", r.OrganizationUUID, "job", r.ResourceUUID, "purge")
	if err != nil {
		result.Error = err
		return result
	}

	page := 1
	limit := 100
	q := url.Query()
	q.Set("page", strconv.Itoa(page))
	q.Set("limit", strconv.Itoa(limit))
	if r.OpID != "" {
		q.Set("opID", r.OpID)
	}
	url.RawQuery = q.Encode()

	var purgeResponse []PurgeStatusResponseRequests

	ctx, cancel := context.WithTimeout(context.Background(), s.request.RequestTimeout)
	defer cancel()

	resp := s.getPurgeStatus(ctx, url)
	if resp.Error != nil {
		result.Error = resp.Error
		return result
	}

	purgeResponse = append(purgeResponse, resp.Body.Payload.Requests...)

	// Auto Pagination
	if resp.Body.Payload.Total > limit {
		total := resp.Body.Payload.Total
		for done := limit; done < total; done += limit {
			page += 1
			ctx, cancel := context.WithTimeout(context.Background(), s.request.RequestTimeout)
			q := url.Query()
			q.Set("page", strconv.Itoa(page))
			url.RawQuery = q.Encode()
			resp := s.getPurgeStatus(ctx, url)
			cancel()
			if resp.Error != nil {
				result.Error = resp.Error
				return result
			}
			purgeResponse = append(purgeResponse, resp.Body.Payload.Requests...)
		}
	}

	result.Status = resp.Status
	result.Headers = resp.Headers
	result.Body = purgeResponse

	return result
}

func (s *PurgeService) getPurgeStatus(ctx context.Context, url url.URL) *common.Result[PurgeStatusResponse] {

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)

	return utils.DoHTTPRequest[PurgeStatusResponse](s.request.GetClient(), req)
}
