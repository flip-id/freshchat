package freshchat

import (
	"encoding/json"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var config = Config{
	BaseUrl:         "https://api.au.freshchat.com/v2/",
	NameSpace:       "8fa2f8b4_cd9a_40d7_8c22_5cd8e19a938e",
	ApiToken:        "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJUOW9ENnZjSDRfbEVUOU8xNlhDd2NzWmkzT3FLX2NtaFl4cjV0TXlCd3FJIn0.eyJqdGkiOiI2YTMwNGU1Ny03OWZhLTQ1ZjMtODY0NC00NTM2OWJkZWM3N2UiLCJleHAiOjE5MjA1MzQyMzksIm5iZiI6MCwiaWF0IjoxNjA1MTc0MjM5LCJpc3MiOiJodHRwOi8vaW50ZXJuYWwtZmMtYXBzZTItMDAtYWxiLWtleWNsb2FrLTEzNDM0NTk2MjIuYXAtc291dGhlYXN0LTIuZWxiLmFtYXpvbmF3cy5jb20vYXV0aC9yZWFsbXMvcHJvZHVjdGlvbiIsImF1ZCI6ImFmMGQ0ODZhLWVkMGYtNDMyOC1hMGY5LTdjOWNmODk4MWE2MiIsInN1YiI6IjI3NTFmOWQwLWI0ZDItNGMwZS04MjViLTRiNDA1ZmNhZDIyNiIsInR5cCI6IkJlYXJlciIsImF6cCI6ImFmMGQ0ODZhLWVkMGYtNDMyOC1hMGY5LTdjOWNmODk4MWE2MiIsImF1dGhfdGltZSI6MCwic2Vzc2lvbl9zdGF0ZSI6IjdhZTI0MTZhLTAwMGEtNDY5Yi1hZjk4LWM5YjlmYmMyZWFlMiIsImFjciI6IjEiLCJhbGxvd2VkLW9yaWdpbnMiOltdLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsib2ZmbGluZV9hY2Nlc3MiLCJ1bWFfYXV0aG9yaXphdGlvbiJdfSwicmVzb3VyY2VfYWNjZXNzIjp7ImFjY291bnQiOnsicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdfX0sInNjb3BlIjoib3V0Ym91bmRtZXNzYWdlOmdldCByZXBvcnRzOmZldGNoIHJlcG9ydHM6ZXh0cmFjdCBhZ2VudDpjcmVhdGUgcm9sZTpyZWFkIGRhc2hib2FyZDpyZWFkIHVzZXI6dXBkYXRlIGFnZW50OmRlbGV0ZSBhZ2VudDpyZWFkIHVzZXI6cmVhZCBhZ2VudDp1cGRhdGUgdXNlcjpjcmVhdGUgbWVzc2FnZTpjcmVhdGUgY29udmVyc2F0aW9uOnJlYWQgb3V0Ym91bmRtZXNzYWdlOnNlbmQgZmlsdGVyaW5ib3g6cmVhZCBjb252ZXJzYXRpb246Y3JlYXRlIHVzZXI6ZGVsZXRlIG1lc3NhZ2U6Z2V0IHJlcG9ydHM6cmVhZCBmaWx0ZXJpbmJveDpjb3VudDpyZWFkIGJpbGxpbmc6dXBkYXRlIGNvbnZlcnNhdGlvbjp1cGRhdGUgcmVwb3J0czpleHRyYWN0OnJlYWQiLCJjbGllbnRIb3N0IjoiMTAuNjkuMTAuMzQiLCJjbGllbnRJZCI6ImFmMGQ0ODZhLWVkMGYtNDMyOC1hMGY5LTdjOWNmODk4MWE2MiIsImNsaWVudEFkZHJlc3MiOiIxMC42OS4xMC4zNCJ9.mDFezqGUHOAe09THagAdyDRggXbZ68GQfoNe-MFBOLz9GvgoSo4z1KtRsf3gdtOgKzaqO0ixyQAqTT-7uSPPrATPNriMrYJRrHORWBHQFitR7qNDHuGohdISMqILxqAyWCYnoNEgBJ4AV0XdwPYt-0mze9J_4jqAtOxu-szsa8Cpcw7WbPZbUw3g7GpePokbAPS3wcQz-xBmxiNHE4EZi6yfmmTb8v06nEcvDVWWGZ2f2WcqqgJ0whA4VsWhgumfBtAqq9wD_o8mP9k7PqEkSUScdyfZOE9APJHQYAJQ2gPT5s0bF3aVf1T_LBoszrHbsz0cVMshZa07n8TMER38Jw",
	FromPhoneNumber: "+6282181526987",
}

func TestSendOtpMessage(t *testing.T) {
	t.Run("success case", func(t *testing.T) {
		sender := New(config)
		httpmock.ActivateNonDefault(sender.Client.GetClient())
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
		url := sender.Config.BaseUrl + SEND_MESSAGE_ENDPOINT
		httpmock.RegisterResponder("POST", url, responder)

		request := OtpRequest{
			ToPhoneNumber: "+628910111213",
			TemplateName:  "account_registration",
			BodyParams:    []string{"14045"},
		}

		result, err := sender.SendOtpMessage(request)

		assertNoError(t, err)

		assert.Equal(t, true, result.IsSuccess, "IsSuccess")
		assert.Equal(t, 202, result.HttpStatusCode, "HttpStatusCode")
		assert.Equal(t, "0fcdd6b6-1f80-4643-a294-8e0625ce30dd", result.MessageId, "MessageId")
		assert.Equal(t, "", result.Message, "Message")
		assert.Equal(t, mockResponse, result.RawData, "RawData")
	})

	t.Run("failed case", func(t *testing.T) {
		sender := New(config)
		httpmock.ActivateNonDefault(sender.Client.GetClient())
		defer httpmock.DeactivateAndReset()

		mockResponse := `{ 
			"success": false,
			"error_code": 404,
			"error_message": "invalid request format"
		}`
		responder := httpmock.NewStringResponder(404, mockResponse)
		url := sender.Config.BaseUrl + SEND_MESSAGE_ENDPOINT
		httpmock.RegisterResponder("POST", url, responder)

		request := OtpRequest{
			ToPhoneNumber: "+628910111213",
			TemplateName:  "account_registration",
			BodyParams:    []string{"14045"},
		}

		result, err := sender.SendOtpMessage(request)

		assertNoError(t, err)

		assert.Equal(t, false, result.IsSuccess, "IsSuccess")
		assert.Equal(t, 404, result.HttpStatusCode, "HttpStatusCode")
		assert.Equal(t, "", result.MessageId, "MessageId")
		assert.Equal(t, "invalid request format", result.Message, "Message")
		assert.Equal(t, mockResponse, result.RawData, "RawData")
	})

	t.Run("error case", func(t *testing.T) {
		sender := New(config)
		httpmock.ActivateNonDefault(sender.Client.GetClient())
		defer httpmock.DeactivateAndReset()

		mockResponse := `{ 
			"Internal server error"
		}`
		responder := httpmock.NewStringResponder(500, mockResponse)
		url := sender.Config.BaseUrl + SEND_MESSAGE_ENDPOINT
		httpmock.RegisterResponder("POST", url, responder)

		request := OtpRequest{
			ToPhoneNumber: "+62891011121",
			TemplateName:  "account_registration",
			BodyParams:    []string{"Test"},
		}

		_, err := sender.SendOtpMessage(request)

		if err == nil {
			t.Errorf("Want an error but didn't get one")
		}
	})
}

func TestMakeRequestBody(t *testing.T) {
	t.Run("success case - one parameter", func(t *testing.T) {
		request := OtpRequest{
			ToPhoneNumber: "+628910111213",
			TemplateName:  "account_registration",
			BodyParams:    []string{"14045"},
		}

		sender := New(config)
		body := sender.makeRequestBody(request)

		byteJsonRequest, _ := json.Marshal(body)
		actualJsonRequest := string(byteJsonRequest)

		expectedJsonRequestFormatted := `{
			"from": {
				"phone_number": "` + sender.Config.FromPhoneNumber + `"
			},
			"provider": "whatsapp",
			"to": [
				{ "phone_number": "+628910111213" }
			],
			"RequestData": {
				"message_template": {
					"storage": "none",
					"template_name": "account_registration",
					"namespace": "` + sender.Config.NameSpace + `",
					"RequestLanguage": {
						"policy": "deterministic",
						"code": "id"
					},
					"rich_template_data": {
						"RequestRichTmplBody": {
							"params": [
								{ "RequestData": "14045" }
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

	t.Run("success case - many parameter", func(t *testing.T) {
		request := OtpRequest{
			ToPhoneNumber: "+628910111213",
			TemplateName:  "info_transfer",
			BodyParams:    []string{"sender_name", "receiver_name", "amount", "YYYY-MM-DD_HH:mm", "https://receipt.link/sample", "flip.id"},
		}

		sender := New(config)
		body := sender.makeRequestBody(request)

		byteJsonRequest, _ := json.Marshal(body)
		actualJsonRequest := string(byteJsonRequest)

		expectedJsonRequestFormatted := `{
			"from": {
				"phone_number": "` + sender.Config.FromPhoneNumber + `"
			},
			"provider": "whatsapp",
			"to": [
				{ "phone_number": "+628910111213" }
			],
			"RequestData": {
				"message_template": {
					"storage": "none",
					"template_name": "info_transfer",
					"namespace": "` + sender.Config.NameSpace + `",
					"RequestLanguage": {
						"policy": "deterministic",
						"code": "id"
					},
					"rich_template_data": {
						"RequestRichTmplBody": {
							"params": [
								{ "RequestData": "sender_name" },
								{ "RequestData": "receiver_name" },
								{ "RequestData": "amount" },
								{ "RequestData": "YYYY-MM-DD_HH:mm" },
								{ "RequestData": "https://receipt.link/sample" },
								{ "RequestData": "flip.id" }
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
