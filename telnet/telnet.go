package telnet

import (
	"errors"
	"io"
	"time"

	"github.com/digitalysin/go-telnet/client"
)

type (
	Telnet interface {
		Run(input io.Reader, output io.Writer) error
		Close() error
	}

	Option struct {
		TelnetHost                  string
		TelnetPort                  uint64
		ConnectTimeout, FeedTimeout time.Duration
	}

	impl struct {
		client *client.TelnetClient
	}
)

func (o *Option) Host() string {
	return o.TelnetHost
}

func (o *Option) Port() uint64 {
	return o.TelnetPort
}

func (o *Option) Timeout() time.Duration {
	return o.FeedTimeout
}

func (o *Option) DialTimeout() time.Duration {
	return o.ConnectTimeout
}

func (i *impl) Run(input io.Reader, output io.Writer) error {
	return i.client.ProcessData(input, output)
}

func (i *impl) Close() error {
	return i.client.Close()
}

func New(opts *Option) (Telnet, error) {
	if opts.TelnetHost == "" {
		return nil, errors.New("telnet host is required")
	}

	var (
		tln, err = client.NewTelnetClient(opts)
	)

	if err != nil {
		return nil, err
	}

	return &impl{
		client: tln,
	}, nil
}
