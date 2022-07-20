package freshchat

import (
	"encoding/json"
	"time"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

const (
	// EndpointSendMessage is the endpoint for sending a message.
	EndpointSendMessage = "/v2/outbound-messages/whatsapp"
)

// Client is the client for sending messages in the Freshchat.
type Client interface{}

type client struct {
	opt *Option
}

func (c *client) Assign(o *Option) *client {
	if o == nil {
		return c
	}

	c.opt = o.Clone()
	return c
}

// NewClient creates a new Freshchat client.
func NewClient(opts ...FnOption) Client {
	o := (new(Option)).Assign(opts...).Default()
	return (new(client)).Assign(o)
}

func (c *client) SendMessage()

func (s *client) SendOtpMessage(otpRequest OtpRequest) (OtpResult, error) {
	body := s.makeRequestBody(otpRequest)

	response, err := s.sendOutboundMessage(body)
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

func (s *client) makeRequestBody(otpRequest OtpRequest) requestBody {
	body := requestBody{}
	body.initialize(s.Config.NameSpace)
	body.setFrom(s.Config.FromPhoneNumber)
	body.addDestination(otpRequest.ToPhoneNumber)
	body.setTemplateName(otpRequest.TemplateName)
	body.setBodyParams(otpRequest.BodyParams)

	return body
}

func (s *client) sendOutboundMessage(body requestBody) (freshchatResponse, error) {
	url := s.Config.BaseUrl + SEND_MESSAGE_ENDPOINT
	response, err := s.Client.R().SetBody(body).Post(url)
	result := freshchatResponse{
		success: nil,
		failed:  nil,
	}

	if response != nil {
		result.httpStatusCode = response.StatusCode()
		result.rawData = string(response.Body())
	}

	if err != nil {
		log.Error(err)
		return result, err
	}

	if ResponseCode(result.httpStatusCode) != Accepted {
		err = json.Unmarshal(response.Body(), &result.failed)

		log.WithFields(log.Fields{
			"message":  "Failed to send WhatsappMessage via Freshchat",
			"response": response,
		}).Warn()
	} else {
		err = json.Unmarshal(response.Body(), &result.success)
	}

	return result, err
}
