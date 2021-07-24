package nats

import (
	"errors"
	"time"

	gnats "github.com/nats-io/nats.go"
)

type (
	Option struct {
		Servers        []string
		User, Password string
		ConnectTimeout time.Duration
	}

	Connection interface {
		Publish(subject string, data []byte) error
		Subscribe(subject string, callback gnats.MsgHandler) (*gnats.Subscription, error)
		QueueSubscribe(subject, group string, callback gnats.MsgHandler) (*gnats.Subscription, error)
		Flush() error
		Close() error
	}

	impl struct {
		client *gnats.Conn
	}
)

func (i *impl) Publish(subject string, data []byte) error {
	return i.client.Publish(subject, data)
}

func (i *impl) Subscribe(subject string, callback gnats.MsgHandler) (*gnats.Subscription, error) {
	return i.client.Subscribe(subject, callback)
}

func (i *impl) QueueSubscribe(subject, group string, callback gnats.MsgHandler) (*gnats.Subscription, error) {
	return i.client.QueueSubscribe(subject, group, callback)
}

func (i *impl) Flush() error {
	return i.client.Flush()
}

func (i *impl) Close() error {
	i.client.Close()
	return nil
}

func New(opts *Option) (Connection, error) {
	if len(opts.Servers) == 0 {
		return nil, errors.New("nats servers is required")
	}

	client, err := gnats.Options{
		Servers:  opts.Servers,
		User:     opts.User,
		Password: opts.Password,
		Timeout:  opts.ConnectTimeout,
	}.Connect()

	if err != nil {
		return nil, err
	}

	return &impl{client}, nil
}
