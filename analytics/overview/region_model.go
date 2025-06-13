package overview

type RegionRequest struct {
	OrganizationUUID string   `json:"organization_uuid"`
	From             string   `json:"from"`
	To               string   `json:"to"`
	Resources        []string `json:"resources"`
}

type RegionResponse struct {
	Status      bool                  `json:"status"`
	OperationID string                `json:"operation_id"`
	Payload     RegionResponsePayload `json:"data"`
}

type RegionResponsePayload struct {
	Meaa   RegionStatisticsResponse `json:"MEAA"`
	Eu     RegionStatisticsResponse `json:"EU"`
	Na     RegionStatisticsResponse `json:"NA"`
	Ap2    RegionStatisticsResponse `json:"AP2"`
	Others RegionStatisticsResponse `json:"others"`
	Ap3    RegionStatisticsResponse `json:"AP3"`
	Ap1    RegionStatisticsResponse `json:"AP1"`
	Sa     RegionStatisticsResponse `json:"SA"`
}

type RegionStatisticsResponse struct {
	TotalTraffic          int    `json:"total-traffic"`
	TotalHit              int    `json:"total-hit"`
	TotalTrafficFormatted string `json:"total-traffic-formatted"`
	TotalHitFormatted     string `json:"total-hit-formatted"`
}
