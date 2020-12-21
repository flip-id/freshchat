package freshchat

type requestBody struct {
	From     phoneNumber   `json:"from"`
	Provider string        `json:"provider"`
	To       []phoneNumber `json:"to"`
	Data     data          `json:"data"`
}

type phoneNumber struct {
	PhoneNumber string `json:"phone_number"`
}

type data struct {
	MessageTemplate messageTemplate `json:"message_template"`
}

type messageTemplate struct {
	Storage          string           `json:"storage"`
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
	Body   body   `json:"body"`
}

type body struct {
	Params []param `json:"params"`
}

type param struct {
	Data string `json:"data"`
}

func (rb *requestBody) initialize() {
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
	rb.To = append(rb.To, phoneNumber{PhoneNumber: number})
}

func (rb *requestBody) setTemplateName(templateName string) {
	rb.Data.MessageTemplate.TemplateName = templateName
}

func (rb *requestBody) setBodyParams(params []string) {
	var bodyParams []param

	for _, inputParam := range params {
		bodyParams = append(bodyParams, param{Data: inputParam})
	}

	rb.Data.MessageTemplate.RichTemplateData.Body.Params = bodyParams
}
