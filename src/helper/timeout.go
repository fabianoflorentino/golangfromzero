package helper

import (
	"os"
	"time"
)

type Config struct {
	DatabaseTimeout time.Duration
	APITimeout      time.Duration
}

var Default = Config{
	DatabaseTimeout: 5 * time.Second,
	APITimeout:      2 * time.Second,
}

var ConfigTimeout = loadConfig()

func loadConfig() Config {
	return Config{
		DatabaseTimeout: getEnvOrDefaultTimeout("DATABASE_TIMEOUT", 5*time.Second),
		APITimeout:      getEnvOrDefaultTimeout("API_TIMEOUT", 2*time.Second),
	}
}

func getEnvOrDefaultTimeout(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		timeout, err := time.ParseDuration(value)
		if err != nil {
			return defaultValue
		}

		return timeout
	}

	return defaultValue
}
