package freshchat

type WhatsappRequest struct {
	FromPhoneNumber string
	ToPhoneNumber   string
	TemplateName    string
	HeaderParams    []string
	BodyParams      []string
}

type WhatsappResult struct {
	IsSuccess      bool
	HttpStatusCode int
	MessageId      string
	Message        string
	RawData        string
}

func SendWhatsappMessage(waRequest WhatsappRequest) (WhatsappResult, error) {
	body := requestBody{}
	body.initialize()
	body.setFrom(waRequest.FromPhoneNumber)
	body.setTo(waRequest.ToPhoneNumber)
	body.setTemplateName(waRequest.TemplateName)
	body.setHeaderParams(waRequest.HeaderParams)
	body.setBodyParams(waRequest.BodyParams)

	response, err := sendOutboundMessage(body)
	var waResult WhatsappResult

	if &response == nil {
		return waResult, err
	}

	waResult.HttpStatusCode = response.httpStatusCode
	waResult.RawData = response.rawData

	if response.success != nil {
		waResult.IsSuccess = true
		waResult.MessageId = response.success.RequestId
	} else if response.failed != nil {
		waResult.IsSuccess = false
		waResult.Message = response.failed.Message
	}

	return waResult, err
}
