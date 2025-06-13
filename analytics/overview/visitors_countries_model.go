package overview

type VisitorsCountriesRequest struct {
	OrganizationUUID string   `json:"organization_uuid"`
	From             string   `json:"from"`
	To               string   `json:"to"`
	Limit            int      `json:"limit,omitempty"`
	Page             int      `json:"page,omitempty"`
	Resources        []string `json:"resources"`
}

type VisitorsCountriesResponse struct {
	Status      bool                             `json:"status"`
	OperationID string                           `json:"operation_id"`
	Payload     VisitorsCountriesResponsePayload `json:"data"`
}

type VisitorsCountriesResponsePayload struct {
	CurrentPage int                             `json:"current_page"`
	LastPage    int                             `json:"last_page"`
	Total       int                             `json:"total"`
	Data        []VisitorsCountriesResponseData `json:"data"`
}

type VisitorsCountriesResponseData struct {
	Hit              int    `json:"hit"`
	Traffic          int64  `json:"traffic"`
	Country          string `json:"country"`
	TrafficGb        string `json:"traffic_gb"`
	TrafficFormatted string `json:"traffic_formatted"`
	HitFormatted     string `json:"hit_formatted"`
}
