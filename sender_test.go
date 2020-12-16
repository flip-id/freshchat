package freshchat

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendWhatsappMessage(t *testing.T) {
	t.Run("success case", func(t *testing.T) {
		httpmock.ActivateNonDefault(client.GetClient())
		defer httpmock.DeactivateAndReset()

		mockResponse := `{ 
			"request_id": "0fcdd6b6-1f80-4643-a294-8e0625ce30dd", 
			"request_process_time": "1", 
			"link": {
				"rel": "string",
				"href": "string"
			}
		}`
		responder := httpmock.NewStringResponder(202, mockResponse)
		url := baseUrl + sendMessageUrl
		httpmock.RegisterResponder("POST", url, responder)

		request := WhatsappRequest{
			FromPhoneNumber: "+62876543210",
			ToPhoneNumber:   "+62891011121",
			TemplateName:    "account_registration",
			HeaderParams:    []string{"Test"},
			BodyParams:      []string{"Test"},
		}

		result, err := SendWhatsappMessage(request)

		assertNoError(t, err)

		assert.Equal(t, true, result.IsSuccess, "IsSuccess")
		assert.Equal(t, 202, result.HttpStatusCode, "HttpStatusCode")
		assert.Equal(t, "0fcdd6b6-1f80-4643-a294-8e0625ce30dd", result.Message, "Message")
		assert.Equal(t, mockResponse, result.RawData, "RawData")
	})

	t.Run("failed case", func(t *testing.T) {
		httpmock.ActivateNonDefault(client.GetClient())
		defer httpmock.DeactivateAndReset()

		mockResponse := `{ 
			"code": 404,
			"status": "AGENT_NOT_FOUND",
			"message": "agent not found"
		}`
		responder := httpmock.NewStringResponder(404, mockResponse)
		url := baseUrl + sendMessageUrl
		httpmock.RegisterResponder("POST", url, responder)

		request := WhatsappRequest{
			FromPhoneNumber: "+62876543210",
			ToPhoneNumber:   "+62891011121",
			TemplateName:    "account_registration",
			HeaderParams:    []string{"Test"},
			BodyParams:      []string{"Test"},
		}

		result, err := SendWhatsappMessage(request)

		assertNoError(t, err)

		assert.Equal(t, false, result.IsSuccess, "IsSuccess")
		assert.Equal(t, 404, result.HttpStatusCode, "HttpStatusCode")
		assert.Equal(t, "agent not found", result.Message, "Message")
		assert.Equal(t, mockResponse, result.RawData, "RawData")
	})

	t.Run("error case", func(t *testing.T) {
		httpmock.ActivateNonDefault(client.GetClient())
		defer httpmock.DeactivateAndReset()

		mockResponse := `{ 
			"Internal server error"
		}`
		responder := httpmock.NewStringResponder(500, mockResponse)
		url := baseUrl + sendMessageUrl
		httpmock.RegisterResponder("POST", url, responder)

		request := WhatsappRequest{
			FromPhoneNumber: "+62876543210",
			ToPhoneNumber:   "+62891011121",
			TemplateName:    "account_registration",
			HeaderParams:    []string{"Test"},
			BodyParams:      []string{"Test"},
		}

		_, err := SendWhatsappMessage(request)

		if err == nil {
			t.Errorf("Want an error but didn't get one")
		}
	})
}

func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("Got an error but didn't want one. The error is %s", err)
	}
}
