package freshchat_client

type requestBody struct {
	From     from   `json:"from"`
	Provider string `json:"provider"`
	To       to     `json:"to"`
	Data     data   `json:"data"`
}

type from struct {
	PhoneNumber string `json:"phone_number"`
}

type to struct {
	PhoneNumber string `json:"phone_number"`
}

type data struct {
	MessageTemplate messageTemplate `json:"message_template"`
}

type messageTemplate struct {
	TemplateName     string           `json:"template_name"`
	Namespace        string           `json:"namespace"`
	Language         language         `json:"language"`
	RichTemplateData richTemplateData `json:"rich_template_data"`
}

type language struct {
	Policy string `json:"policy"`
	Code   string `json:"code"`
}

type richTemplateData struct {
	Header header `json:"header"`
	Body   body   `json:"body"`
}

type header struct {
	Type   string  `json:"type"`
	Params []param `json:"params"`
}

type body struct {
	Params []param `json:"params"`
}

type param struct {
	Data string `json:"data"`
}

func (rb *requestBody) initialize() {
	rb.Provider = "whatsapp"
	rb.Data.MessageTemplate.Namespace = namespace
	rb.Data.MessageTemplate.Language.Policy = "deterministic"
	rb.Data.MessageTemplate.Language.Code = "id"
	rb.Data.MessageTemplate.RichTemplateData.Header.Type = "text"
}

func (rb *requestBody) setFrom(phoneNumber string) {
	rb.From.PhoneNumber = phoneNumber
}

func (rb *requestBody) setTo(phoneNumber string) {
	rb.To.PhoneNumber = phoneNumber
}

func (rb *requestBody) setTemplateName(templateName string) {
	rb.Data.MessageTemplate.TemplateName = templateName
}

func (rb *requestBody) setHeaderParams(params []string) {
	var headerParams []param

	for _, inputParam := range params {
		headerParams = append(headerParams, param{Data: inputParam})
	}

	rb.Data.MessageTemplate.RichTemplateData.Header.Params = headerParams
}

func (rb *requestBody) setBodyParams(params []string) {
	var bodyParams []param

	for _, inputParam := range params {
		bodyParams = append(bodyParams, param{Data: inputParam})
	}

	rb.Data.MessageTemplate.RichTemplateData.Body.Params = bodyParams
}
