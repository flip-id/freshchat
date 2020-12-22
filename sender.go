package freshchat

type OtpRequest struct {
	ToPhoneNumber string
	TemplateName  string
	BodyParams    []string
}

type OtpResult struct {
	IsSuccess      bool
	HttpStatusCode int
	MessageId      string
	Message        string
	RawData        string
}

func SendOtpMessage(otpRequest OtpRequest) (OtpResult, error) {
	body := makeRequestBody(otpRequest)

	response, err := sendOutboundMessage(body)
	var otpResult OtpResult

	if &response == nil {
		return otpResult, err
	}

	otpResult.HttpStatusCode = response.httpStatusCode
	otpResult.RawData = response.rawData

	if response.success != nil {
		otpResult.IsSuccess = true
		otpResult.MessageId = response.success.RequestId
	} else if response.failed != nil {
		otpResult.IsSuccess = false
		otpResult.Message = response.failed.ErrorMessage
	}

	return otpResult, err
}

func makeRequestBody(otpRequest OtpRequest) requestBody {
	body := requestBody{}
	body.initialize()
	body.setFrom(fromPhoneNumber)
	body.addDestination(otpRequest.ToPhoneNumber)
	body.setTemplateName(otpRequest.TemplateName)
	body.setBodyParams(otpRequest.BodyParams)

	return body
}
