//go:build integration
// +build integration

package freshchat

import (
	"context"
	"flag"
	"github.com/fairyhunter13/dotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"os"
	"sync"
	"testing"
	"time"
)

var (
	c    Client
	once sync.Once
)

func setupClient() {
	once.Do(func() {
		err := dotenv.Load2(
			dotenv.WithPaths(".env"),
		)
		if err != nil {
			log.Fatalln(err)
		}

		c = NewClient(
			WithApiToken(os.Getenv("FRESHCHAT_API_TOKEN")),
			WithNameSpace(os.Getenv("FRESHCHAT_NAMESPACE")),
			WithFromPhoneNumber(os.Getenv("FRESHCHAT_FROM_PHONE_NUMBER")),
		)
	})
}

// Run integration tests.
// Notes: Run this test only on local, not on CI/CD.
func TestMain(m *testing.M) {
	flag.Parse()
	setupClient()

	os.Exit(m.Run())
}

func formatTime(in time.Time) string {
	return in.Format("2006-01-02 15:04:05")
}

func TestSendWhatsappMessageSuccess(t *testing.T) {
	ctx := context.Background()
	req := (new(RequestWhatsappMessage)).
		AddDestination(os.Getenv("PHONE_NUMBER")).
		SetTemplateName("otp_info").
		SetParams([]string{
		formatTime(time.Now()),
	})
	resp, err := c.SendMessage(ctx, req)

	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Success.RequestID)
	assert.NotEmpty(t, resp.Success.Link.Href)
	assert.NotEmpty(t, resp.Success.Link.Rel)
	assert.NotEmpty(t, resp.Success.Link.Type)
	assert.NotEmpty(t, resp.Success.RequestProcessTime)
	assert.Nil(t, resp.Failed)
	assert.Equal(t, http.StatusAccepted, resp.HTTPStatusCode)
	assert.NotEmpty(t, resp.RawData)
}

func TestSendWhatsappMessageFailed(t *testing.T) {
	t.Run("no template name", func(t *testing.T) {
		ctx := context.Background()
		req := (new(RequestWhatsappMessage)).
			AddDestination(os.Getenv("PHONE_NUMBER")).
			SetParams([]string{
			formatTime(time.Now()),
		})
		resp, err := c.SendMessage(ctx, req)

		assert.NotNil(t, err)
		assert.NotNil(t, resp)
		assert.Nil(t, resp.Success)
		assert.NotNil(t, resp.Failed)
		assert.False(t, resp.Failed.Success)
		assert.Equal(t, 5, resp.Failed.ErrorCode)
		assert.Equal(t, "template_name is required", resp.Failed.ErrorMessage)
		assert.Equal(t, http.StatusBadRequest, resp.HTTPStatusCode)
		assert.NotEmpty(t, resp.RawData)
	})
}
