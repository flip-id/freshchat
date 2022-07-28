package freshchat

import (
	"github.com/fairyhunter13/reflecthelper/v5"
	"github.com/gojek/heimdall/v7"
	"github.com/gojek/heimdall/v7/hystrix"
	"net/http"
	"strings"
	"time"
)

const (
	// DefaultBaseURLAustralia is the default base URL for the Freshchat API in Australia.
	DefaultBaseURLAustralia = "https://api.au.freshchat.com"
	// DefaultTimeout is the default timeout for the client.
	DefaultTimeout = 30 * time.Second
)

// Option is an option for the Freshchat API
type Option struct {
	BaseURL         string
	NameSpace       string
	APIToken        string
	FromPhoneNumber string
	Timeout         time.Duration
	Client          heimdall.Doer
	HystrixOptions  []hystrix.Option
	client          *hystrix.Client
}

// FnOption is a functional option for the Freshchat API
type FnOption func(o *Option)

// Clone clones the option and returns a new one.
func (o *Option) Clone() *Option {
	opt := *o
	return &opt
}

// Assign assigns the functional options to the option.
func (o *Option) Assign(opts ...FnOption) *Option {
	for _, opt := range opts {
		opt(o)
	}

	return o
}

// Default returns option with default values.
func (o *Option) Default() *Option {
	if o.BaseURL == "" {
		o.BaseURL = DefaultBaseURLAustralia
	}

	o.BaseURL = strings.TrimRight(o.BaseURL, "/")
	if o.Timeout < DefaultTimeout {
		o.Timeout = DefaultTimeout
	}

	if reflecthelper.IsNil(o.Client) {
		o.Client = http.DefaultClient
	}

	o.client = hystrix.NewClient(
		hystrix.WithHystrixTimeout(o.Timeout),
		hystrix.WithHTTPTimeout(o.Timeout),
		hystrix.WithHTTPClient(o.Client),
	)
	return o
}

// WithBaseURL sets the base URL for the Freshchat API.
func WithBaseURL(baseURL string) FnOption {
	return func(o *Option) {
		o.BaseURL = baseURL
	}
}

// WithNameSpace sets the name space for the Freshchat API.
func WithNameSpace(nameSpace string) FnOption {
	return func(o *Option) {
		o.NameSpace = nameSpace
	}
}

// WithAPIToken sets the API token for the Freshchat API.
func WithAPIToken(apiToken string) FnOption {
	return func(o *Option) {
		o.APIToken = apiToken
	}
}

// WithFromPhoneNumber sets the from phone number for the Freshchat API.
func WithFromPhoneNumber(fromPhoneNumber string) FnOption {
	return func(o *Option) {
		o.FromPhoneNumber = fromPhoneNumber
	}
}

// WithTimeout sets the timeout for the Freshchat API.
func WithTimeout(timeout time.Duration) FnOption {
	return func(o *Option) {
		o.Timeout = timeout
	}
}

// WithClient sets the client for the Freshchat API.
func WithClient(client heimdall.Doer) FnOption {
	return func(o *Option) {
		o.Client = client
	}
}

// WithHystrixOptions sets the hystrix options for the Freshchat API.
func WithHystrixOptions(options []hystrix.Option) FnOption {
	return func(o *Option) {
		o.HystrixOptions = options
	}
}
