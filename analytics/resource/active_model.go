package resource

// ActiveDataResponse struct, used by active_traffic and active_hit
type ActiveDataResponse struct {
	Status  bool       `json:"status"`
	Payload ActiveData `json:"data"`
}

type ActiveData struct {
	PerResources map[string]ActiveDataPoint `json:"per_resources"`
	Overall      ActiveDataPoint            `json:"overall"`
}

type ActiveDataPoint struct {
	Current Number `json:"current"`
	Max     Number `json:"max"`
	Unit    string `json:"unit"`
}
