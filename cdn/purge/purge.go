package purge

import (
	"context"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"

	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/client/utils"
)

type PurgeService struct {
	request *request.RequestConfig
}

func NewPurgeService(reqCfg *request.RequestConfig) *PurgeService {
	return &PurgeService{request: reqCfg}
}

func (s *PurgeService) Purge(organizationUUID string, p PurgeRequest) (*PurgeResponse, error) {

	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "cdn", organizationUUID, "job", p.ResourceUUID, "purge")

	body, err := utils.ToJSONBodyBuffer(p)
	if err != nil {
		return &PurgeResponse{}, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), s.request.RequestTimeout*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		return &PurgeResponse{}, err
	}

	return utils.DoHTTPRequest[*PurgeResponse](s.request.GetClient(), req)
}

func (s *PurgeService) Status(r PurgeStatusRequest) (*[]PurgeStatusResponseRequests, error) {

	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "cdn", r.OrganizationUUID, "job", r.ResourceUUID, "purge")
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

	ctx, cancel := context.WithTimeout(context.Background(), s.request.RequestTimeout*time.Second)
	defer cancel()

	resp, err := s.getPurgeStatus(ctx, url)
	if err != nil {
		return &purgeResponse, err
	}

	purgeResponse = append(purgeResponse, resp.Payload.Requests...)

	// Auto Pagination
	if resp.Payload.Total > limit {
		total := resp.Payload.Total
		for done := limit; done <= total; done += limit {
			page += 1
			ctx, cancel := context.WithTimeout(context.Background(), s.request.RequestTimeout*time.Second)
			q := url.Query()
			q.Set("page", strconv.Itoa(page))
			url.RawQuery = q.Encode()
			resp, err := s.getPurgeStatus(ctx, url)
			cancel()
			if err != nil {
				return &[]PurgeStatusResponseRequests{}, err
			}
			purgeResponse = append(purgeResponse, resp.Payload.Requests...)
		}
	}

	return &purgeResponse, nil
}

func (s *PurgeService) getPurgeStatus(ctx context.Context, url url.URL) (*PurgeStatusResponse, error) {

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)

	return utils.DoHTTPRequest[*PurgeStatusResponse](s.request.GetClient(), req)
}
