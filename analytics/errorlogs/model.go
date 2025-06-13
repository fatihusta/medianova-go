package errorlogs

type ErrorLogsStatusCodesRequest struct {
	PageSize         int      `json:"page_size,omitempty"`
	OrganizationUUID string   `json:"organization_uuid"`
	From             string   `json:"from"`
	To               string   `json:"to"`
	StatusCodes      []int    `json:"statusCodes"`
	Resources        []string `json:"resources,omitempty"`
}

type ErrorLogsStatusCodesResponse struct {
	Status      bool   `json:"status"`
	OperationID string `json:"operation_id"`
	Payload     any    `json:"data"`
}

type ErrorLogsStatusCodesErrorList struct {
	Hit      int    `json:"hit"`
	URL      string `json:"url"`
	Protocol string `json:"protocol"`
	Method   string `json:"method"`
}
type ErrorLogsStatusCodesData struct {
	CurrentPage     int                             `json:"current_page"`
	PageSize        int                             `json:"page_size"`
	Total           int                             `json:"total"`
	ErrorStatusCode string                          `json:"error_status_code"`
	ErrorList       []ErrorLogsStatusCodesErrorList `json:"error_list"`
}
