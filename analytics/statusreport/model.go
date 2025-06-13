package statusreport

type StatusReportRequest struct {
	OrganizationUUID string   `json:"organization_uuid"`
	From             string   `json:"from"`
	To               string   `json:"to"`
	Resources        []string `json:"resources"`
	Interval         string   `json:"interval,omitempty"`
}

type StatusReportResponse struct {
	Status      bool   `json:"status"`
	OperationID string `json:"operation_id"`
	Payload     any    `json:"data"`
}

// StatusCode enum
type StatusCode string
