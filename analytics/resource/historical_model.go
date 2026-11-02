package resource

// HistoricalResponse struct.
// The API fills only the fields related to the requested type
// (hit -> hit, traffic -> traffic, bandwidth -> bandwidth, requests -> count).
type HistoricalResponse struct {
	Status  bool           `json:"status"`
	Payload HistoricalData `json:"data"`
}

// HistoricalHitRatioResponse struct is the response of the hit_ratio type,
// which returns a flat ratio per resource instead of the timed series.
type HistoricalHitRatioResponse struct {
	Status  bool                   `json:"status"`
	Payload HistoricalHitRatioData `json:"data"`
}

type HistoricalHitRatioData struct {
	PerResources map[string]HitRatioValue `json:"per_resources"`
	Overall      HitRatioValue            `json:"overall"`
}

type HitRatioValue struct {
	HitRatio Number `json:"hit_ratio"`
}

// HistoricalGPUTransactionResponse struct is the response of the gpu_transaction type,
// which returns a flat count per resource instead of the timed series.
type HistoricalGPUTransactionResponse struct {
	Status  bool                         `json:"status"`
	Payload HistoricalGPUTransactionData `json:"data"`
}

type HistoricalGPUTransactionData struct {
	PerResources map[string]GPUTransactionValue `json:"per_resources"`
	Overall      GPUTransactionValue            `json:"overall"`
}

type GPUTransactionValue struct {
	GPUTransaction Number `json:"gpu_transaction"`
}

type HistoricalData struct {
	PerResources map[string]HistoricalMetricSeries `json:"per_resources"`
	Overall      HistoricalMetricOverall           `json:"overall"`
	TimePeriods  []string                          `json:"time_periods"`
}

type HistoricalMetricOverall struct {
	Status  bool                   `json:"status"`
	Payload HistoricalMetricSeries `json:"data"`
}

type HistoricalMetricSeries struct {
	SelectedTime HistoricalMetricPeriod  `json:"selected_time"`
	Previous     HistoricalMetricPeriod  `json:"previous"`
	Indicate     Number                  `json:"indicate"`
	ThisMonth    *HistoricalMetricPeriod `json:"this_month"`
	LastMonth    *HistoricalMetricPeriod `json:"last_month"`
}

type HistoricalMetricPeriod struct {
	From               string `json:"from"`
	To                 string `json:"to"`
	Traffic            Number `json:"traffic"`
	Hit                Number `json:"hit"`
	Bandwidth          Number `json:"bandwidth"`
	Count              Number `json:"count"`
	TrafficFormatted   string `json:"traffic_formatted"`
	HitFormatted       string `json:"hit_formatted"`
	BandwidthFormatted string `json:"bandwidth_formatted"`
	CountFormatted     string `json:"count_formatted"`
}
