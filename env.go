package environment

import (
	env "github.com/caarlos0/env/v10"
)

type Config struct {
	// Application logging level.
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
	// Ignore the service package.
	Simpler bool `env:"PACKAGE_SIMPLER" envDefault:"true"`
	// Strict mode for checking the name of services and methods.
	StrictMethodTitle bool `env:"STRICT_METHOD_TITLE" envDefault:"true"`

	GRPCNetwork string `env:"GRPC_NETWORK" envDefault:"tcp"`
	GRPCHost    string `env:"GRPC_HOST" envDefault:"0.0.0.0"`
	GRPCPort    string `env:"GRPC_PORT" envDefault:"4770"`
	GRPCAddr    string `env:",expand" envDefault:"$GRPC_HOST:$GRPC_PORT"`

	HTTPHost string `env:"HTTP_HOST" envDefault:"0.0.0.0"`
	HTTPPort string `env:"HTTP_PORT" envDefault:"4771"`
	HTTPAddr string `env:",expand" envDefault:"$HTTP_HOST:$HTTP_PORT"`
}

func New() (Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
