package freshchat

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/fairyhunter13/pool"
	"github.com/gofiber/fiber/v2"
)

const (
	// EndpointSendMessage is the endpoint for sending a message.
	EndpointSendMessage = "/v2/outbound-messages/whatsapp"
)

// Client is the client for sending messages in the Freshchat.
type Client interface {
	SendMessage(ctx context.Context, req *RequestWhatsappMessage) (res *ResponseFreshchat, err error)
}

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

// SendMessage sends a Whatsapp message to the Freshchat.
func (c *client) SendMessage(ctx context.Context, req *RequestWhatsappMessage) (res *ResponseFreshchat, err error) {
	if req == nil {
		err = ErrNilArguments
		return
	}

	resp, err := c.doRequest(ctx, EndpointSendMessage, req.Default(c.opt))
	defer c.closeResponse(resp)
	if err != nil {
		return
	}

	res, err = (new(ResponseFreshchat)).assign(resp)
	return
}

func (c *client) closeResponse(resp *http.Response) {
	if resp == nil || resp.Body == nil {
		return
	}

	_ = resp.Body.Close()
}

const (
	// HeaderBearerPrefix is the prefix for the Bearer token.
	HeaderBearerPrefix = "Bearer "
)

func (c *client) prepareRequest(ctx context.Context, req *http.Request) *http.Request {
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	req.Header.Set(fiber.HeaderAuthorization, HeaderBearerPrefix+c.opt.APIToken)
	req = req.WithContext(ctx)
	return req
}

func (c *client) doRequest(ctx context.Context, endpoint string, message interface{}) (resp *http.Response, err error) {
	buff := pool.GetBuffer()
	defer pool.Put(buff)

	err = json.NewEncoder(buff).Encode(message)
	if err != nil {
		return
	}

	url := c.opt.BaseURL + endpoint
	req, err := http.NewRequest(http.MethodPost, url, buff)
	if err != nil {
		return
	}

	resp, err = c.opt.client.Do(c.prepareRequest(ctx, req))
	return
}
