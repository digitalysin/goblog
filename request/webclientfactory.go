package request

import (
	"time"

	"github.com/gojek/heimdall/v7/httpclient"
)

type (
	WebClientFactory interface {
		Create(timeout time.Duration) Client
	}
	clientFactory struct {
		timeout time.Duration
	}
)

func (cf *clientFactory) Create(timeout time.Duration) Client {
	return &client{Doer: *httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))}
}

func NewClientFactory() WebClientFactory {
	return &clientFactory{}
}
