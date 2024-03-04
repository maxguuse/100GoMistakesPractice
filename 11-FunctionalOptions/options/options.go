package options

import (
	"errors"
	"net/http"
	"strconv"
)

type options struct {
	port *int
}

type Option func(options *options) error

func WithPort(port int) Option {
	return func(o *options) error {
		if port < 0 {
			return errors.New("port must be positive")
		}

		o.port = &port
		return nil
	}
}

func NewServer(addr string, opts ...Option) (*http.Server, error) {
	var options options
	for _, opt := range opts {
		if err := opt(&options); err != nil {
			return nil, err
		}
	}

	var port int
	if options.port == nil || *options.port == 0 {
		port = 8080
	} else {
		port = *options.port
	}

	return &http.Server{
		Addr: addr + ":" + strconv.Itoa(port),
	}, nil
}
