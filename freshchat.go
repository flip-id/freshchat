package freshchat

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

type ResponseCode int

const (
	Success             ResponseCode = 200
	Accepted            ResponseCode = 202
	BadRequest          ResponseCode = 400
	Unauthenticated     ResponseCode = 401
	Forbidden           ResponseCode = 403
	NotFound            ResponseCode = 404
	TooManyRequests     ResponseCode = 429
	InternalServerError ResponseCode = 500
	ServiceUnavailable  ResponseCode = 503
)

// Request body for sending Whatsapp message via Freshchat
type RequestBody struct {
	ClientId        string            `json:"client_id"`
	ProjectId       string            `json:"project_id"`
	Type            string            `json:"type"`
	RecipientNumber string            `json:"recipient_number"`
	Params          map[string]string `json:"params"`
}

// Success response from Freshchat
type SuccessResponse struct {
	RequestId          string `json:"request_id"`
	RequestProcessTime string `json:"request_process_time"`
	Rel                string `json:"rel"`
	Href               string `json:"href"`
}

// Error response from Freshchat
type ErrorResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Response body after post request to Freshchat
type ResMessage struct {
	MessageId      string `json:"message_id"`
	Status         string `json:"status"`
	Message        string `json:"message"`
	HttpStatusCode int
	RawData        string
}

// Send Whatsapp message
func sendWaMessage(req RequestBody) (ResMessage, error) {
	res, err := postToFreshchat(sendMessageUrl, req)

	return res, err
}

// Post request to Freshchat service
func postToFreshchat(endpoint string, body interface{}) (ResMessage, error) {
	url := baseUrl + endpoint
	res, err := client.R().SetBody(body).Post(url)
	resMessage := ResMessage{}

	if res != nil {
		resMessage.HttpStatusCode = res.StatusCode()
		resMessage.RawData = string(res.Body())
	}

	if err != nil {
		log.Error(err)
		return resMessage, err
	}

	if res.StatusCode() != 200 {
		log.WithFields(log.Fields{
			"msg": "Status code is not we expected",
			"res": res,
		}).Warn()
	}

	if err := json.Unmarshal(res.Body(), &resMessage); err != nil {
		return resMessage, err
	}

	return resMessage, err
}
