package errorreport

type ErrorReportStatusCodesRequest struct {
	OrganizationUUID string   `json:"organization_uuid"`
	From             string   `json:"from"`
	To               string   `json:"to"`
	Resources        []string `json:"resources"`
}

type ErrorReportStatusCodesResponse struct {
	Status      bool                       `json:"status"`
	OperationID string                     `json:"operation_id"`
	Payload     ErrorReportStatusCodesData `json:"data"`
}

type ErrorReportStatusCodesSuccessfulResponse struct {
	SuccessResponses         int    `json:"success_responses"`
	OtherSuccessResponses    int    `json:"other_success_responses"`
	TotalSuccessfulResponse  int    `json:"total_successful_response"`
	SuccessfulResponsesRatio string `json:"successful_responses_ratio"`
}

type ErrorReportStatusCodesRedirectResponse struct {
	TotalRedirectResponses      int    `json:"total_redirect_responses"`
	MovedPermanentlyResponses   int    `json:"moved_permanently_responses"`
	FoundResponses              int    `json:"found_responses"`
	OtherRedirectResponses      int    `json:"other_redirect_responses"`
	TotalRedirectResponsesRatio string `json:"total_redirect_responses_ratio"`
}

type ErrorReportStatusCodesClientErrorResponse struct {
	TotalClientErrorResponses int    `json:"total_client_error_responses"`
	ForbiddenResponses        int    `json:"forbidden_responses"`
	NotFoundResponses         int    `json:"not_found_responses"`
	ToManyRequest             int    `json:"to_many_request"`
	OtherClientErrorResponses int    `json:"other_client_error_responses"`
	TotalClientErrorRatio     string `json:"total_client_error_ratio"`
}

type ErrorReportStatusCodesServerErrorResponse struct {
	TotalServerErrorResponses    int    `json:"total_server_error_responses"`
	InternalServerErrorResponses int    `json:"internal_server_error_responses"`
	BadGateway                   int    `json:"bad_gateway"`
	GatewayTimeout               int    `json:"gateway_timeout"`
	OtherServerErrorResponses    int    `json:"other_server_error_responses"`
	TotalServerErrorRatio        string `json:"total_server_error_ratio"`
}

type ErrorReportStatusCodesData struct {
	TotalRequest         int                                       `json:"total_request"`
	SuccessfulResponses  ErrorReportStatusCodesSuccessfulResponse  `json:"successful_responses"`
	RedirectResponses    ErrorReportStatusCodesRedirectResponse    `json:"redirect_responses"`
	ClientErrorResponses ErrorReportStatusCodesClientErrorResponse `json:"client_error_responses"`
	ServerErrorResponses ErrorReportStatusCodesServerErrorResponse `json:"server_error_responses"`
}
