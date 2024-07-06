package environment

import (
	env "github.com/caarlos0/env/v10"
)

type Config struct {
	// Application logging level.
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
	// Ignore the service package.
	Simpler bool `env:"SERVICE_SIMPLER" envDefault:"true"`

	GRPCNetwork string `env:"GRPC_NETWORK" envDefault:"tcp"`
	GRPCHost    string `env:"GRPC_HOST" envDefault:"0.0.0.0"`
	GRPCPort    string `env:"GRPC_PORT" envDefault:"4770"`
	GRPCAddr    string `env:",expand" envDefault:"$GRPC_HOST:$GRPC_PORT"`

	HTTPHost string `env:"HTTP_HOST" envDefault:"0.0.0.0"`
	HTTPPort string `env:"HTTP_PORT" envDefault:"4771"`
	HTTPAddr string `env:",expand" envDefault:"$HTTP_HOST:$HTTP_PORT"`

	OtlpHost  string  `env:"OTLP_TRACE_GRPC_HOST" envDefault:"127.0.0.1"`
	OtlpPort  string  `env:"OTLP_TRACE_GRPC_PORT" envDefault:"4317"`
	OtlpTLS   bool    `env:"OTLP_TRACE_TLS" envDefault:"false"`
	OtlpRatio float64 `env:"OTLP_SAMPLE_RATIO" envDefault:"0.0"`
}

func New() (Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
