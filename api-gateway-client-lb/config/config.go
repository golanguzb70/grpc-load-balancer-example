package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	HTTPPort        string
	PostServiceHost string
	PostServicePort string
}

func Load() *Config {
	cfg := &Config{}

	cfg.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", "8080"))
	cfg.PostServiceHost = cast.ToString(getOrReturnDefault("POST_SERVICE_HOST", "localhost"))
	cfg.PostServicePort = cast.ToString(getOrReturnDefault("POST_SERVICE_PORT", "9000"))

	return cfg
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
