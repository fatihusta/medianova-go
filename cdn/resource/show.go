package resource

import (
	"context"
	"net/http"

	"github.com/fatihusta/medianova-go/client/utils"
	"github.com/fatihusta/medianova-go/common"
)

func (s *ResourceService) Show(organizationUUID, resourceUUID string) *common.Result[ResourceDetailsResponse] {

	result := common.NewResult[ResourceDetailsResponse]()

	url, err := utils.JoinPath(s.request.BaseURL, "cdn", organizationUUID, "resource", resourceUUID)
	if err != nil {
		result.Error = err
		return result
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.request.RequestTimeout)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)

	return utils.DoHTTPRequest[ResourceDetailsResponse](s.request.GetClient(), req)
}
