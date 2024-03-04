package builder

import (
	"errors"
	"net/http"
	"strconv"
)

type Config struct {
	Port int
}

type configBuilder struct {
	port *int
}

func New() *configBuilder {
	return &configBuilder{}
}

func (b *configBuilder) Port(port int) *configBuilder {
	b.port = &port
	return b
}

func (b *configBuilder) Build() (*Config, error) {
	cfg := &Config{}

	if b.port == nil || *b.port == 0 {
		cfg.Port = 8080
	} else if *b.port < 0 {
		return &Config{}, errors.New("port must be positive")
	} else {
		cfg.Port = *b.port
	}

	return cfg, nil
}

func NewServer(addr string, config *Config) (*http.Server, error) {
	return &http.Server{
		Addr: addr + ":" + strconv.Itoa(config.Port),
	}, nil
}
