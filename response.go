package freshchat

import (
	"encoding/json"
	"io"
	"net/http"
)

// ResponseFreshchat is the response from Freshchat.
type ResponseFreshchat struct {
	Success        *ResponseSuccess `json:"-"`
	Failed         *ResponseFailed  `json:"-"`
	HTTPStatusCode int              `json:"-"`
	RawData        string           `json:"-"`
}

// assign assigns the values from the given *http.Response.
func (r *ResponseFreshchat) assign(resp *http.Response) (res *ResponseFreshchat, err error) {
	if resp == nil || resp.Body == nil {
		return
	}

	byteBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var ptrDecode interface{}
	if resp.StatusCode >= http.StatusBadRequest {
		ptrDecode = &r.Failed
	} else {
		ptrDecode = &r.Success
	}

	err = json.Unmarshal(byteBody, ptrDecode)
	if err != nil {
		return
	}

	defer func() {
		res = r
	}()
	r.HTTPStatusCode = resp.StatusCode
	r.RawData = convertByteToString(byteBody)
	err = r.getError()
	return
}

func (r *ResponseFreshchat) getError() (err error) {
	if r.HTTPStatusCode >= http.StatusBadRequest {
		err = r.Failed
		return
	}

	return
}

// ResponseSuccess is the success response from Freshchat.
type ResponseSuccess struct {
	RequestID          string              `json:"request_id"`
	RequestProcessTime string              `json:"request_process_time"`
	Link               ResponseSuccessLink `json:"link"`
}

// ResponseSuccessLink is the success link response from Freshchat.
type ResponseSuccessLink struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
	Type string `json:"type"`
}

// ResponseFailed is the failed response from Freshchat.
type ResponseFailed struct {
	Success      bool   `json:"success"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}
