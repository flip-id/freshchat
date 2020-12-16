package freshchat_client

type freshchatResponse struct {
	success        *successResponse
	failed         *failedResponse
	httpStatusCode int
	rawData        string
}

type successResponse struct {
	RequestId          string `json:"request_id"`
	RequestProcessTime string `json:"request_process_time"`
	Link               link   `json:"link"`
}

type link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

type failedResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
