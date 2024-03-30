package environment

import "net"

type GRPC struct {
	Network string `envconfig:"GRPC_NETWORK" default:"tcp"`
	Host    string `envconfig:"GRPC_HOST" default:"0.0.0.0"`
	Port    string `envconfig:"GRPC_PORT" default:"4770"`
}

func (c GRPC) Addr() string {
	return net.JoinHostPort(c.Host, c.Port)
}
