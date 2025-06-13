package overview

type TopResourcesRequest struct {
	OrganizationUUID string   `json:"organization_uuid"`
	Resources        []string `json:"resources"`
	From             string   `json:"from"`
	To               string   `json:"to"`
	Limit            int      `json:"limit,omitempty"`
	Page             int      `json:"page,omitempty"`
}

type TopResourcesResponse struct {
	Status      bool                        `json:"status"`
	OperationID string                      `json:"operation_id"`
	Payload     TopResourcesResponsePayload `json:"data"`
}

type TopResourcesResponsePayload struct {
	CurrentPage int                        `json:"current_page"`
	LastPage    int                        `json:"last_page"`
	Total       int                        `json:"total"`
	Data        []TopResourcesResponseData `json:"data"`
}

type TopResourcesResponseData struct {
	Resource         string `json:"resource"`
	Type             string `json:"type"`
	Traffic          int64  `json:"traffic"`
	TrafficFormatted string `json:"traffic_formatted"`
	Hit              int    `json:"hit"`
	HitFormatted     string `json:"hit_formatted"`
	Bandwidth        string `json:"bandwidth"`
}
