package config

import (
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	GinMode    string
	LogLevel   string
	ESHost     string
	ESUser     string
	ESPassword string
}

func Load() (*Config, error) {
	// Загружаем .env только если он существует
	_ = godotenv.Load()

	return &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		GinMode:    getEnv("GIN_MODE", "release"),
		LogLevel:   getEnv("LOG_LEVEL", "warning"),
		ESHost:     getEnv("ES_HOST", ""),
		ESUser:     getEnv("ES_USER", ""),
		ESPassword: getEnv("ES_PASSWORD", ""),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
