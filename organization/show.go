package organization

import (
	"context"
	"net/http"

	"github.com/fatihusta/medianova-go/client/utils"
	"github.com/fatihusta/medianova-go/common"
)

func (s *OrganizationService) Show(organizationUUID string) *common.Result[OrganizationDetailsResponse] {
	result := common.NewResult[OrganizationDetailsResponse]()

	url, err := utils.JoinPath(s.request.BaseURLV2, "organizations", organizationUUID)
	if err != nil {
		result.Error = err
		return result
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.request.RequestTimeout)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)

	return utils.DoHTTPRequest[OrganizationDetailsResponse](s.request.GetClient(), req)
}
