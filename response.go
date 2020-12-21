package freshchat

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
	Type string `json:"type"`
}

type failedResponse struct {
	Success      bool   `json:"success"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}
