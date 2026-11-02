package resource

// ErrorsResponse struct
type ErrorsResponse struct {
	Status  bool       `json:"status"`
	Payload ErrorsData `json:"data"`
}

type ErrorsData struct {
	PerResources map[string][]ErrorCountItem `json:"per_resources"`
	Overall      []ErrorCountItem            `json:"overall"`
}

type ErrorCountItem struct {
	StatusCode string `json:"status_code"`
	Total      Number `json:"total"`
}

// VisitorsCountriesResponse struct
type VisitorsCountriesResponse struct {
	Status  bool                  `json:"status"`
	Payload VisitorsCountriesData `json:"data"`
}

type VisitorsCountriesData struct {
	PerResources map[string][]CountryVisitorItem `json:"per_resources"`
	Overall      []CountryVisitorItem            `json:"overall"`
}

type CountryVisitorItem struct {
	Country          string `json:"country"`
	Traffic          Number `json:"traffic"`
	TrafficGb        Number `json:"traffic_gb"`
	TrafficFormatted string `json:"traffic_formatted"`
	Hit              Number `json:"hit"`
	HitFormatted     string `json:"hit_formatted"`
}

// StatusResponse struct
type StatusResponse struct {
	Status  bool       `json:"status"`
	Payload StatusData `json:"data"`
}

type StatusData struct {
	PerResources map[string]StatusPerResourceItem `json:"per_resources"`
	Overall      StatusOverallItem                `json:"overall"`
}

type StatusOverallItem struct {
	TotalRequest         int64                      `json:"total_request"`
	SuccessfulResponses  SuccessfulResponsesOverall `json:"successful_responses"`
	RedirectResponses    RedirectResponses          `json:"redirect_responses"`
	ClientErrorResponses ClientErrorResponses       `json:"client_error_responses"`
	ServerErrorResponses ServerErrorResponses       `json:"server_error_responses"`
}

type StatusPerResourceItem struct {
	SuccessfulResponses  SuccessfulResponsesPerResource `json:"successful_responses"`
	RedirectResponses    RedirectResponses              `json:"redirect_responses"`
	ClientErrorResponses ClientErrorResponses           `json:"client_error_responses"`
	ServerErrorResponses ServerErrorResponses           `json:"server_error_responses"`
}

type SuccessfulResponsesOverall struct {
	TotalSuccessfulResponse  int64  `json:"total_successful_response"`
	SuccessResponses         int64  `json:"success_responses"`
	OtherSuccessResponses    int64  `json:"other_success_responses"`
	SuccessfulResponsesRatio string `json:"successful_responses_ratio"`
}

type SuccessfulResponsesPerResource struct {
	TotalSuccessfulResponses int64  `json:"total_successful_responses"`
	SuccessResponses         int64  `json:"success_responses"`
	OtherSuccessResponses    int64  `json:"other_success_responses"`
	SuccessfulResponsesRatio string `json:"successful_responses_ratio"`
}

type RedirectResponses struct {
	TotalRedirectResponses      int64  `json:"total_redirect_responses"`
	MovedPermanentlyResponses   int64  `json:"moved_permanently_responses"`
	FoundResponses              int64  `json:"found_responses"`
	OtherRedirectResponses      int64  `json:"other_redirect_responses"`
	TotalRedirectResponsesRatio string `json:"total_redirect_responses_ratio"`
}

type ClientErrorResponses struct {
	TotalClientErrorResponses     int64  `json:"total_client_error_responses"`
	UnsupportedMediaTypeResponses int64  `json:"unsupported_media_type_responses"`
	ForbiddenResponses            int64  `json:"forbidden_responses"`
	NotFoundResponses             int64  `json:"not_found_responses"`
	ToManyRequest                 int64  `json:"to_many_request"`
	OtherClientErrorResponses     int64  `json:"other_client_error_responses"`
	TotalClientErrorRatio         string `json:"total_client_error_ratio"`
}

type ServerErrorResponses struct {
	TotalServerErrorResponses    int64  `json:"total_server_error_responses"`
	InternalServerErrorResponses int64  `json:"internal_server_error_responses"`
	BadGateway                   int64  `json:"bad_gateway"`
	GatewayTimeout               int64  `json:"gateway_timeout"`
	OtherServerErrorResponses    int64  `json:"other_server_error_responses"`
	TotalServerErrorRatio        string `json:"total_server_error_ratio"`
}

// RegionResponse struct
type RegionResponse struct {
	Status  bool       `json:"status"`
	Payload RegionData `json:"data"`
}

type RegionData struct {
	PerResources map[string]map[string]RegionAggregate `json:"per_resources"`
	Overall      map[string]RegionAggregate            `json:"overall"`
}

type RegionAggregate struct {
	TotalTraffic          Number `json:"total-traffic"`
	TotalHit              Number `json:"total-hit"`
	TotalTrafficFormatted string `json:"total-traffic-formatted"`
	TotalHitFormatted     string `json:"total-hit-formatted"`
}
