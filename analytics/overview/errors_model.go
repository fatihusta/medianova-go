package overview

type ErrorsRequest struct {
	Limit            int      `json:"limit,omitempty"`
	Page             int      `json:"page,omitempty"`
	OrganizationUUID string   `json:"organization_uuid"`
	From             string   `json:"from"`
	To               string   `json:"to"`
	Resources        []string `json:"resources"`
}

type ErrorsResponse struct {
	Status      bool                 `json:"status"`
	OperationID string               `json:"operation_id"`
	Payload     []ErrorsResponseData `json:"data"`
}

type ErrorsResponseData struct {
	Total      int    `json:"total"`
	StatusCode string `json:"status_code"`
}
