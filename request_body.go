package freshchat

import (
	"github.com/fairyhunter13/phone"
	"strings"
)

const (
	// DefaultProviderWhatsapp is the default provider for sending messages in the Freshchat.
	DefaultProviderWhatsapp = "whatsapp"
	// DefaultStorage is the default storage for sending messages.
	DefaultStorage = "none"
	// DefaultLanguagePolicy is the default language policy for sending messages.
	DefaultLanguagePolicy = "deterministic"
	// DefaultLanguageCode is the default language code for sending messages.
	DefaultLanguageCode = "id"
)

// RequestWhatsappMessage is a request for sending a message to a user.
type RequestWhatsappMessage struct {
	From     RequestFrom    `json:"from"`
	Provider string         `json:"provider"`
	To       []*RequestFrom `json:"to"`
	Data     RequestData    `json:"data"`

	// Response Attributes Below
	MessageID     string `json:"message_id,omitempty"`
	RequestID     string `json:"request_id,omitempty"`
	Status        string `json:"status,omitempty"`
	FailureCode   string `json:"failure_code,omitempty"`
	FailureReason string `json:"failure_reason,omitempty"`
	CreatedOn     int    `json:"created_on,omitempty"`
}

// SetFrom sets the from/sender for the message.
func (r *RequestWhatsappMessage) SetFrom(from string) *RequestWhatsappMessage {
	if from == "" {
		return r
	}

	r.From.PhoneNumber = from
	return r
}

// SetProvider sets the provider for the message.
func (r *RequestWhatsappMessage) SetProvider(provider string) *RequestWhatsappMessage {
	if provider == "" {
		return r
	}

	r.Provider = provider
	return r
}

// AddDestination adds the destination for the message.
func (r *RequestWhatsappMessage) AddDestination(phoneNumber string) *RequestWhatsappMessage {
	if phoneNumber == "" {
		return r
	}

	r.To = append(r.To, &RequestFrom{PhoneNumber: phoneNumber})
	return r
}

// SetDestinations sets the destinations for the message.
func (r *RequestWhatsappMessage) SetDestinations(phoneNumbers []string) *RequestWhatsappMessage {
	if len(phoneNumbers) <= 0 {
		return r
	}

	var newTo []*RequestFrom
	for _, number := range phoneNumbers {
		newTo = append(newTo, &RequestFrom{PhoneNumber: number})
	}

	r.To = newTo
	return r
}

// SetStorage sets the storage for the message.
func (r *RequestWhatsappMessage) SetStorage(storage string) *RequestWhatsappMessage {
	if storage == "" {
		return r
	}

	r.Data.MessageTemplate.Storage = storage
	return r
}

// SetTemplateName sets the template name for the message.
func (r *RequestWhatsappMessage) SetTemplateName(tmpl string) *RequestWhatsappMessage {
	if tmpl == "" {
		return r
	}

	r.Data.MessageTemplate.TemplateName = tmpl
	return r
}

// SetNamespace sets the namespace for the message.
func (r *RequestWhatsappMessage) SetNamespace(ns string) *RequestWhatsappMessage {
	if ns == "" {
		return r
	}

	r.Data.MessageTemplate.Namespace = ns
	return r
}

// SetLanguagePolicy sets the language policy for the message.
func (r *RequestWhatsappMessage) SetLanguagePolicy(policy string) *RequestWhatsappMessage {
	if policy == "" {
		return r
	}

	r.Data.MessageTemplate.Language.Policy = policy
	return r
}

// SetLanguageCode sets the language code for the message.
func (r *RequestWhatsappMessage) SetLanguageCode(code string) *RequestWhatsappMessage {
	if code == "" {
		return r
	}

	r.Data.MessageTemplate.Language.Code = code
	return r
}

// SetHeaderMediaURL sets the header media URL for the message.
func (r *RequestWhatsappMessage) SetHeaderMediaURL(url string) *RequestWhatsappMessage {
	if url == "" {
		return r
	}

	if r.Data.MessageTemplate.RichTemplateData.Header == nil {
		r.Data.MessageTemplate.RichTemplateData.Header = new(RequestRichTmplHeader)
	}

	r.Data.MessageTemplate.RichTemplateData.Header.MediaURL = url
	return r
}

// SetHeaderType sets the header type for the message.
func (r *RequestWhatsappMessage) SetHeaderType(headerType string) *RequestWhatsappMessage {
	if headerType == "" {
		return r
	}

	if r.Data.MessageTemplate.RichTemplateData.Header == nil {
		r.Data.MessageTemplate.RichTemplateData.Header = new(RequestRichTmplHeader)
	}

	r.Data.MessageTemplate.RichTemplateData.Header.Type = headerType
	return r
}

// SetHeaderParams sets the header params for the message.
func (r *RequestWhatsappMessage) SetHeaderParams(params []string) *RequestWhatsappMessage {
	if len(params) <= 0 ||
		r.Data.MessageTemplate.RichTemplateData.Header == nil ||
		r.Data.MessageTemplate.RichTemplateData.Header.Type != "text" {
		return r
	}

	var newParams []RequestRichTmplParams
	for _, param := range params {
		newParams = append(newParams, RequestRichTmplParams{Data: param})
	}

	r.Data.MessageTemplate.RichTemplateData.Header.Params = newParams
	return r
}

// SetParams sets the params for the message.
func (r *RequestWhatsappMessage) SetParams(params []string) *RequestWhatsappMessage {
	if len(params) <= 0 {
		return r
	}

	var newParams []RequestRichTmplParams
	for _, param := range params {
		newParams = append(newParams, RequestRichTmplParams{Data: param})
	}

	r.Data.MessageTemplate.RichTemplateData.Body.Params = newParams
	return r
}

// Default is a default for RequestWhatsappMessage.
func (r *RequestWhatsappMessage) Default(o *Option) *RequestWhatsappMessage {
	r.From.Default(o, TypePhoneSender)
	if r.Provider == "" {
		r.Provider = DefaultProviderWhatsapp
	}

	for idx, to := range r.To {
		r.To[idx] = to.Default(o, TypePhoneDestination)
	}

	r.Data.Default(o)
	return r
}

// RequestFrom is a request for specifying the sender number.
type RequestFrom struct {
	PhoneNumber string `json:"phone_number"`
}

const (
	// TypePhoneSender is the type for phone sender.
	TypePhoneSender = "sender"
	// TypePhoneDestination is the type for phone destination.
	TypePhoneDestination = "destination"
)

// Default is a default for RequestFrom.
func (r *RequestFrom) Default(o *Option, typePhone string) *RequestFrom {
	if r.PhoneNumber == "" && o != nil && typePhone == TypePhoneSender {
		r.PhoneNumber = o.FromPhoneNumber
	}

	r.PhoneNumber = strings.TrimLeft(r.PhoneNumber, "+")
	r.PhoneNumber = "+" + phone.NormalizeID(r.PhoneNumber, 0)
	return r
}

// RequestData is a request for specifying the message data.
type RequestData struct {
	MessageTemplate RequestMessageTemplate `json:"message_template"`
}

// Default is a default for RequestData.
func (r *RequestData) Default(o *Option) *RequestData {
	r.MessageTemplate.Default(o)
	return r
}

// RequestMessageTemplate is a request for specifying the message template.
type RequestMessageTemplate struct {
	Storage          string                  `json:"storage"`
	TemplateName     string                  `json:"template_name"`
	Namespace        string                  `json:"namespace"`
	Language         RequestLanguage         `json:"language"`
	RichTemplateData RequestRichTemplateData `json:"rich_template_data"`
}

// Default is a default for RequestMessageTemplate.
func (r *RequestMessageTemplate) Default(o *Option) *RequestMessageTemplate {
	if r.Storage == "" {
		r.Storage = DefaultStorage
	}

	if r.Namespace == "" && o != nil {
		r.Namespace = o.NameSpace
	}

	r.Language.Default(o)
	return r
}

// RequestLanguage is a request for specifying the language.
type RequestLanguage struct {
	Policy string `json:"policy"`
	Code   string `json:"code"`
}

// Default is a default for RequestLanguage.
func (r *RequestLanguage) Default(o *Option) *RequestLanguage {
	if r.Policy == "" {
		r.Policy = DefaultLanguagePolicy
	}

	if r.Code == "" {
		r.Code = DefaultLanguageCode
	}
	return r
}

// RequestRichTemplateData is a request for specifying the rich template data.
type RequestRichTemplateData struct {
	Header *RequestRichTmplHeader `json:"header,omitempty"`
	Body   RequestRichTmplBody    `json:"body"`
}

// RequestRichTmplHeader is a request for specifying the rich template header.
type RequestRichTmplHeader struct {
	Type     string                  `json:"type,omitempty"`
	MediaURL string                  `json:"media_url,omitempty"`
	Params   []RequestRichTmplParams `json:"params,omitempty"`
}

// RequestRichTmplBody is a request for specifying the rich template body.
type RequestRichTmplBody struct {
	Params []RequestRichTmplParams `json:"params"`
}

// RequestRichTmplParams is a request for specifying the rich template params data.
type RequestRichTmplParams struct {
	Data string `json:"data"`
}
