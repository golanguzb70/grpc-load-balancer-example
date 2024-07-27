package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	GrpcPort string
}

func Load() *Config {
	cfg := &Config{}

	cfg.GrpcPort = cast.ToString(getOrReturnDefault("GRPC_PORT", "80"))

	return cfg
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
