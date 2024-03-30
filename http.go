package environment

import "net"

type HTTP struct {
	Host string `envconfig:"HTTP_HOST" default:"0.0.0.0"`
	Port string `envconfig:"HTTP_PORT" default:"4771"`
}

func (c HTTP) Addr() string {
	return net.JoinHostPort(c.Host, c.Port)
}
