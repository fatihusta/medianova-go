package stook

import (
	"context"
	"net/http"
	"net/url"

	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/client/utils"
	"github.com/fatihusta/medianova-go/common"
)

// ListBuckets returns all Stook buckets of the organization.
func (s *StookService) ListBuckets(organizationUUID string) *common.Result[[]Bucket] {
	result := common.NewResult[[]Bucket]()

	url, err := utils.JoinPath(s.request.BaseURLV2, "stook", organizationUUID, "bucket")
	if err != nil {
		result.Error = err
		return result
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.request.RequestTimeout)
	defer cancel()

	resp := doGet[bucketListResponse](ctx, s.request, url)
	if resp.Error != nil {
		result.Error = resp.Error
		return result
	}

	result.Status = resp.Status
	result.Headers = resp.Headers
	result.Body = resp.Body.Payload

	return result
}

// GetBucket returns the details of a single Stook bucket, including its
// credentials and FTP configuration.
func (s *StookService) GetBucket(organizationUUID, bucketUUID string) *common.Result[BucketDetail] {
	result := common.NewResult[BucketDetail]()

	url, err := utils.JoinPath(s.request.BaseURLV2, "stook", organizationUUID, "bucket", bucketUUID)
	if err != nil {
		result.Error = err
		return result
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.request.RequestTimeout)
	defer cancel()

	resp := doGet[bucketDetailResponse](ctx, s.request, url)
	if resp.Error != nil {
		result.Error = resp.Error
		return result
	}

	result.Status = resp.Status
	result.Headers = resp.Headers
	result.Body = resp.Body.Payload

	return result
}

func doGet[T any](ctx context.Context, reqCfg *request.RequestConfig, url url.URL) *common.Result[T] {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		result := common.NewResult[T]()
		result.Error = err
		return result
	}

	return utils.DoHTTPRequest[T](reqCfg.GetClient(), req)
}
