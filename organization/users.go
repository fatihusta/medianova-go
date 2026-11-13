package organization

import (
	"context"
	"net/http"

	"github.com/fatihusta/medianova-go/client/utils"
	"github.com/fatihusta/medianova-go/common"
)

func (s *OrganizationService) Users(organizationUUID string) *common.Result[OrganizationUsersResponse] {
	result := common.NewResult[OrganizationUsersResponse]()

	url, err := utils.JoinPath(s.request.BaseURL, "organizations", organizationUUID, "users")
	if err != nil {
		result.Error = err
		return result
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.request.RequestTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		result.Error = err
		return result
	}

	return utils.DoHTTPRequest[OrganizationUsersResponse](s.request.GetClient(), req)
}
