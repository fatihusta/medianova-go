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

type StatusReporResponePayload2xx struct {
	Status200 []StatusReportHit `json:"200"`
	Status2xx []StatusReportHit `json:"2xx"`
}

type StatusReporResponePayload3xx struct {
	Status301 []StatusReportHit `json:"301"`
	Status302 []StatusReportHit `json:"302"`
	Status3xx []StatusReportHit `json:"3xx"`
}

type StatusReporResponePayload4xx struct {
	Status403 []StatusReportHit `json:"403"`
	Status404 []StatusReportHit `json:"404"`
	Status429 []StatusReportHit `json:"429"`
	Status4xx []StatusReportHit `json:"4xx"`
}

type StatusReporResponePayload5xx struct {
	Status500 []StatusReportHit `json:"500"`
	Status502 []StatusReportHit `json:"502"`
	Status504 []StatusReportHit `json:"504"`
	Status5xx []StatusReportHit `json:"5xx"`
}

type StatusReportHit struct {
	Hit  int   `json:"hit"`
	Time int64 `json:"time"`
}
