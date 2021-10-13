package freshchat

import (
	"encoding/json"
	"time"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

const SEND_MESSAGE_ENDPOINT = "/v2/outbound-messages/whatsapp"
const CONNECTION_TIME_OUT = 15

type Config struct {
	BaseUrl         string
	NameSpace       string
	ApiToken        string
	FromPhoneNumber string
}

type Sender struct {
	Config Config
	Client *resty.Client
}

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

func New(config Config) *Sender {
	return &Sender{
		Config: config,
		Client: resty.New().SetTimeout(time.Second * time.Duration(CONNECTION_TIME_OUT)).SetAuthToken(config.ApiToken),
	}
}

func (s *Sender) SendOtpMessage(otpRequest OtpRequest) (OtpResult, error) {
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

func (s *Sender) makeRequestBody(otpRequest OtpRequest) requestBody {
	body := requestBody{}
	body.initialize(s.Config.NameSpace)
	body.setFrom(s.Config.FromPhoneNumber)
	body.addDestination(otpRequest.ToPhoneNumber)
	body.setTemplateName(otpRequest.TemplateName)
	body.setBodyParams(otpRequest.BodyParams)

	return body
}


func (s *Sender) sendOutboundMessage(body requestBody) (freshchatResponse, error) {
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
