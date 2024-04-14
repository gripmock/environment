package environment

import (
	env "github.com/caarlos0/env/v10"
)

type Config struct {
	// Application logging level.
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
	// Ignore the service package.
	Simpler bool `env:"SERVICE_SIMPLER" envDefault:"true"`

	GrpcNetwork string `env:"GRPC_NETWORK" envDefault:"tcp"`
	GrpcHost    string `env:"GRPC_HOST" envDefault:"127.0.0.1"`
	GrpcPort    string `env:"GRPC_PORT" envDefault:"4770"`
	GrpcAddr    string `env:",expand" envDefault:"$GRPC_HOST:$GRPC_PORT"`

	HttpHost string `env:"HTTP_HOST" envDefault:"127.0.0.1"`
	HttpPort string `env:"HTTP_PORT" envDefault:"4771"`
	HttpAddr string `env:",expand" envDefault:"$HTTP_HOST:$HTTP_PORT"`

	OplpHost        string  `env:"OTLP_TRACE_GRPC_HOST" envDefault:"127.0.0.1"`
	OplpPort        string  `env:"OTLP_TRACE_GRPC_PORT" envDefault:"4317"`
	OplpTLS         bool    `env:"OTLP_TRACE_TLS" envDefault:"false"`
	OplpSampleRatio float64 `env:"OTLP_SAMPLE_RATIO" envDefault:"1.0"`
}

func New() (Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
