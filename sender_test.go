package freshchat

import (
	"encoding/json"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestSendOtpMessage(t *testing.T) {
	t.Run("success case", func(t *testing.T) {
		httpmock.ActivateNonDefault(client.GetClient())
		defer httpmock.DeactivateAndReset()

		mockResponse := `{ 
			"request_id": "0fcdd6b6-1f80-4643-a294-8e0625ce30dd", 
			"request_process_time": "1", 
			"link": {
				"rel": "string",
				"href": "string",
				"type": "GET"
			}
		}`
		responder := httpmock.NewStringResponder(202, mockResponse)
		url := baseUrl + sendMessageUrl
		httpmock.RegisterResponder("POST", url, responder)

		request := OtpRequest{
			ToPhoneNumber: "+628910111213",
			TemplateName:  "account_registration",
			BodyParams:    []string{"14045"},
		}

		result, err := SendOtpMessage(request)

		assertNoError(t, err)

		assert.Equal(t, true, result.IsSuccess, "IsSuccess")
		assert.Equal(t, 202, result.HttpStatusCode, "HttpStatusCode")
		assert.Equal(t, "0fcdd6b6-1f80-4643-a294-8e0625ce30dd", result.MessageId, "MessageId")
		assert.Equal(t, "", result.Message, "Message")
		assert.Equal(t, mockResponse, result.RawData, "RawData")
	})

	t.Run("failed case", func(t *testing.T) {
		httpmock.ActivateNonDefault(client.GetClient())
		defer httpmock.DeactivateAndReset()

		mockResponse := `{ 
			"success": false,
			"error_code": 404,
			"error_message": "invalid request format"
		}`
		responder := httpmock.NewStringResponder(404, mockResponse)
		url := baseUrl + sendMessageUrl
		httpmock.RegisterResponder("POST", url, responder)

		request := OtpRequest{
			ToPhoneNumber: "+628910111213",
			TemplateName:  "account_registration",
			BodyParams:    []string{"14045"},
		}

		result, err := SendOtpMessage(request)

		assertNoError(t, err)

		assert.Equal(t, false, result.IsSuccess, "IsSuccess")
		assert.Equal(t, 404, result.HttpStatusCode, "HttpStatusCode")
		assert.Equal(t, "", result.MessageId, "MessageId")
		assert.Equal(t, "invalid request format", result.Message, "Message")
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

		request := OtpRequest{
			ToPhoneNumber: "+62891011121",
			TemplateName:  "account_registration",
			BodyParams:    []string{"Test"},
		}

		_, err := SendOtpMessage(request)

		if err == nil {
			t.Errorf("Want an error but didn't get one")
		}
	})
}

func TestMakeRequestBody(t *testing.T) {
	t.Run("success case", func(t *testing.T) {
		request := OtpRequest{
			ToPhoneNumber: "+628910111213",
			TemplateName:  "account_registration",
			BodyParams:    []string{"14045"},
		}

		body := makeRequestBody(request)

		byteJsonRequest, _ := json.Marshal(body)
		actualJsonRequest := string(byteJsonRequest)

		expectedJsonRequestFormatted := `{
			"from": {
				"phone_number": "` + fromPhoneNumber + `"
			},
			"provider": "whatsapp",
			"to": [
				{ "phone_number": "+628910111213" }
			],
			"data": {
				"message_template": {
					"storage": "none",
					"template_name": "account_registration",
					"namespace": "` + namespace + `",
					"language": {
						"policy": "deterministic",
						"code": "id"
					},
					"rich_template_data": {
						"body": {
							"params": [
								{ "data": "14045" }
							]
						}
					}
				}
			}
		}`
		expectedJsonRequest := strings.ReplaceAll(expectedJsonRequestFormatted, "\t", "")
		expectedJsonRequest = strings.ReplaceAll(expectedJsonRequest, "\n", "")
		expectedJsonRequest = strings.ReplaceAll(expectedJsonRequest, " ", "") // Will also remove whitespace in params' value, be careful when changing this test

		assert.Equal(t, expectedJsonRequest, actualJsonRequest)
	})
}

func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("Got an error but didn't want one. The error is %s", err)
	}
}
