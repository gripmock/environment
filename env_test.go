package environment_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gripmock/environment"
)

func TestConfig_Defaults(t *testing.T) {
	conf, err := environment.New()
	require.NoError(t, err)

	require.Equal(t, "info", conf.LogLevel)

	require.True(t, conf.Simpler)

	require.Equal(t, "tcp", conf.GRPCNetwork)
	require.Equal(t, "127.0.0.1", conf.GRPCHost)
	require.Equal(t, "4770", conf.GRPCPort)
	require.Equal(t, "127.0.0.1:4770", conf.GRPCAddr)

	require.Equal(t, "127.0.0.1", conf.HTTPHost)
	require.Equal(t, "4771", conf.HTTPPort)
	require.Equal(t, "127.0.0.1:4771", conf.HTTPAddr)

	require.Equal(t, "127.0.0.1", conf.OtlpHost)
	require.Equal(t, "4317", conf.OtlpPort)
	require.False(t, conf.OtlpTLS)
	require.InEpsilon(t, 1.0, conf.OtlpRatio, 0.0001)
}

func TestConfig_Override(t *testing.T) {
	env := map[string]string{
		"LOG_LEVEL":            "trace",
		"SERVICE_SIMPLER":      "false",
		"GRPC_NETWORK":         "udp",
		"GRPC_HOST":            "192.168.1.1",
		"GRPC_PORT":            "1000",
		"GRPC_ADDR":            "111.111.111.111:1111", // no override
		"HTTP_HOST":            "192.168.1.2",
		"HTTP_PORT":            "2000",
		"HTTP_ADDR":            "222.222.222.222:2222",
		"OTLP_TRACE_GRPC_HOST": "1.1.1.1",
		"OTLP_TRACE_GRPC_PORT": "8888",
		"OTLP_TRACE_TLS":       "true",
		"OTLP_SAMPLE_RATIO":    "0.001",
	}

	for k, v := range env {
		err := os.Setenv(k, v)
		require.NoError(t, err)
	}

	conf, err := environment.New()
	require.NoError(t, err)

	require.Equal(t, "trace", conf.LogLevel)

	require.False(t, conf.Simpler)

	require.Equal(t, "udp", conf.GRPCNetwork)
	require.Equal(t, "192.168.1.1", conf.GRPCHost)
	require.Equal(t, "1000", conf.GRPCPort)
	require.Equal(t, "192.168.1.1:1000", conf.GRPCAddr)

	require.Equal(t, "192.168.1.2", conf.HTTPHost)
	require.Equal(t, "2000", conf.HTTPPort)
	require.Equal(t, "192.168.1.2:2000", conf.HTTPAddr)

	require.Equal(t, "1.1.1.1", conf.OtlpHost)
	require.Equal(t, "8888", conf.OtlpPort)
	require.True(t, conf.OtlpTLS)
	require.InEpsilon(t, 0.001, conf.OtlpRatio, 0.0001)
}
