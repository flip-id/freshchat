package freshchat

// RequestWhatsappMessage is a request for sending a message to a user.
type RequestWhatsappMessage struct {
	From     RequestFrom   `json:"from"`
	Provider string        `json:"provider"`
	To       []RequestFrom `json:"to"`
	Data     RequestData   `json:"data"`
}

// RequestFrom is a request for specifying the sender number.
type RequestFrom struct {
	PhoneNumber string `json:"phone_number"`
}

// RequestData is a request for specifying the message data.
type RequestData struct {
	MessageTemplate RequestMessageTemplate `json:"message_template"`
}

// RequestMessageTemplate is a request for specifying the message template.
type RequestMessageTemplate struct {
	Storage          string                  `json:"storage"`
	TemplateName     string                  `json:"template_name"`
	Namespace        string                  `json:"namespace"`
	Language         RequestLanguage         `json:"language"`
	RichTemplateData RequestRichTemplateData `json:"rich_template_data"`
}

// RequestLanguage is a request for specifying the language.
type RequestLanguage struct {
	Policy string `json:"policy"`
	Code   string `json:"code"`
}

// RequestRichTemplateData is a request for specifying the rich template data.
type RequestRichTemplateData struct {
	Body RequestRichTmplBody `json:"body"`
}

// RequestRichTmplBody is a request for specifying the rich template body.
type RequestRichTmplBody struct {
	Params []RequestRichTmplParamsData `json:"params"`
}

type RequestRichTmplParamsData struct {
	Data string `json:"data"`
}

func (rb *requestBody) initialize(namespace string) {
	rb.Provider = "whatsapp"
	rb.Data.MessageTemplate.Storage = "none"
	rb.Data.MessageTemplate.Namespace = namespace
	rb.Data.MessageTemplate.Language.Policy = "deterministic"
	rb.Data.MessageTemplate.Language.Code = "id"
}

func (rb *requestBody) setFrom(number string) {
	rb.From.PhoneNumber = number
}

func (rb *requestBody) addDestination(number string) {
	rb.To = append(rb.To, PhoneNumber{PhoneNumber: number})
}

func (rb *requestBody) setTemplateName(templateName string) {
	rb.Data.MessageTemplate.TemplateName = templateName
}

func (rb *requestBody) setBodyParams(params []string) {
	var bodyParams []RequestRichTmplParamsData

	for _, inputParam := range params {
		bodyParams = append(bodyParams, RequestRichTmplParamsData{Data: inputParam})
	}

	rb.Data.MessageTemplate.RichTemplateData.Body.Params = bodyParams
}
