package stook

import (
	"context"
	"net/http"
	"net/url"
	"path"

	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/client/utils"
	"github.com/fatihusta/medianova-go/common"
)

// ListBuckets returns all Stook buckets of the organization.
func (s *StookService) ListBuckets(organizationUUID string) *common.Result[[]Bucket] {
	result := common.NewResult[[]Bucket]()

	url := *s.request.BaseURLV2
	url.Path = path.Join(url.Path, "stook", organizationUUID, "bucket")

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

	url := *s.request.BaseURLV2
	url.Path = path.Join(url.Path, "stook", organizationUUID, "bucket", bucketUUID)

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
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)

	return utils.DoHTTPRequest[T](reqCfg.GetClient(), req)
}
