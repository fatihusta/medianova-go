package prefetch

import (
	"context"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"

	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/client/utils"
	"github.com/fatihusta/medianova-go/common"
)

type PrefetchService struct {
	request *request.RequestConfig
}

func NewPrefetchService(reqCfg *request.RequestConfig) *PrefetchService {
	return &PrefetchService{request: reqCfg}
}

func (s *PrefetchService) Prefetch(organizationUUID string, p PrefetchRequest) *common.Result[PrefetchResponse] {

	result := common.NewResult[PrefetchResponse]()

	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "cdn", organizationUUID, "job", p.ResourceUUID, "prefetch")

	body, err := utils.ToJSONBodyBuffer(p)
	if err != nil {
		result.Error = err
		return result
	}
	ctx, cancel := context.WithTimeout(context.Background(), s.request.RequestTimeout*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		result.Error = err
		return result
	}

	return utils.DoHTTPRequest[PrefetchResponse](s.request.GetClient(), req)
}

func (s *PrefetchService) Status(r PrefetchStatusRequest) *common.Result[[]PrefetchStatusResponseRequests] {

	result := common.NewResult[[]PrefetchStatusResponseRequests]()

	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "cdn", r.OrganizationUUID, "job", r.ResourceUUID, "prefetch")
	page := 1
	limit := 100
	q := url.Query()
	q.Set("page", strconv.Itoa(page))
	q.Set("limit", strconv.Itoa(limit))
	if r.OpID != "" {
		q.Set("opID", r.OpID)
	}
	url.RawQuery = q.Encode()

	var prefetchResponse []PrefetchStatusResponseRequests

	ctx, cancel := context.WithTimeout(context.Background(), s.request.RequestTimeout*time.Second)
	defer cancel()

	resp := s.getPrefetchStatus(ctx, url)
	if resp.Error != nil {
		result.Error = resp.Error
		return result
	}

	prefetchResponse = append(prefetchResponse, resp.Body.Payload.Requests...)

	// Auto Pagination
	if resp.Body.Payload.Total > limit {
		total := resp.Body.Payload.Total
		for done := limit; done <= total; done += limit {
			page += 1
			ctx, cancel := context.WithTimeout(context.Background(), s.request.RequestTimeout*time.Second)
			q := url.Query()
			q.Set("page", strconv.Itoa(page))
			url.RawQuery = q.Encode()
			resp := s.getPrefetchStatus(ctx, url)
			cancel()
			if resp.Error != nil {
				result.Error = resp.Error
				return result
			}
			prefetchResponse = append(prefetchResponse, resp.Body.Payload.Requests...)
		}
	}

	result.Status = resp.Status
	result.Headers = resp.Headers
	result.Body = prefetchResponse

	return result
}

func (s *PrefetchService) getPrefetchStatus(ctx context.Context, url url.URL) *common.Result[PrefetchStatusResponse] {

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)

	return utils.DoHTTPRequest[PrefetchStatusResponse](s.request.GetClient(), req)
}
