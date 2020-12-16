package freshchat

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

func sendOutboundMessage(body requestBody) (freshchatResponse, error) {
	url := baseUrl + sendMessageUrl
	response, err := client.R().SetBody(body).Post(url)
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

	if responseCode(result.httpStatusCode) != Accepted {
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
