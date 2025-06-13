package prefetch

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"

	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/client/utils"
)

type PrefetchService struct {
	request *request.RequestConfig
}

func NewPrefetchService(reqCfg *request.RequestConfig) *PrefetchService {
	return &PrefetchService{request: reqCfg}
}

func (s *PrefetchService) Prefetch(organizationUUID string, p PrefetchRequest) (*PrefetchResponse, error) {

	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "cdn", organizationUUID, "job", p.ResourceUUID, "prefetch")

	body, err := utils.ToJSONBodyBuffer(p)
	if err != nil {
		return &PrefetchResponse{}, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), s.request.RequestTimeout*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		return &PrefetchResponse{}, err
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json")

	resp, err := s.request.Do(req)
	if err != nil {
		return &PrefetchResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg, err := utils.ToStringBody(resp)
		if err == nil {
			return &PrefetchResponse{}, fmt.Errorf("request not succeeded. status:%d, error:%s", resp.StatusCode, errMsg)
		}
		return &PrefetchResponse{}, fmt.Errorf("request not succeeded. status:%d", resp.StatusCode)
	}

	return utils.FromJSONToStruct[*PrefetchResponse](resp)
}

func (s *PrefetchService) Status(r PrefetchStatusRequest) (*[]PrefetchStatusResponseRequests, error) {

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

	resp, err := s.getPrefetchStatus(ctx, url)
	if err != nil {
		return &prefetchResponse, err
	}

	prefetchResponse = append(prefetchResponse, resp.Payload.Requests...)

	// Auto Pagination
	if resp.Payload.Total > limit {
		total := resp.Payload.Total
		for done := limit; done <= total; done += limit {
			page += 1
			ctx, cancel := context.WithTimeout(context.Background(), s.request.RequestTimeout*time.Second)
			q := url.Query()
			q.Set("page", strconv.Itoa(page))
			url.RawQuery = q.Encode()
			resp, err := s.getPrefetchStatus(ctx, url)
			cancel()
			if err != nil {
				return &[]PrefetchStatusResponseRequests{}, err
			}
			prefetchResponse = append(prefetchResponse, resp.Payload.Requests...)
		}
	}

	return &prefetchResponse, nil
}

func (s *PrefetchService) getPrefetchStatus(ctx context.Context, url url.URL) (*PrefetchStatusResponse, error) {

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)

	req.Header.Set("accept", "application/json")

	resp, err := s.request.Do(req)
	if err != nil {
		return &PrefetchStatusResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg, err := utils.ToStringBody(resp)
		if err == nil {
			return &PrefetchStatusResponse{}, fmt.Errorf("request not succeeded. status:%d, error:%s", resp.StatusCode, errMsg)
		}
		return &PrefetchStatusResponse{}, fmt.Errorf("request not succeeded. status:%d", resp.StatusCode)
	}

	return utils.FromJSONToStruct[*PrefetchStatusResponse](resp)
}
