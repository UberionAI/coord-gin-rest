package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config holds all application configuration
type Config struct {
	ServerPort string
	GinMode    string
	LogLevel   string

	// Elasticsearch config (for future use)
	// ESHost     string
	// ESUser     string
	// ESPassword string
}

// Load reads configuration from environment variables
func Load() (*Config, error) {
	// Load .env file if exists (optional in production)
	_ = godotenv.Load()

	cfg := &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		GinMode:    getEnv("GIN_MODE", "release"),
		LogLevel:   getEnv("LOG_LEVEL", "warning"),

		// Future Elasticsearch config
		// ESHost:     getEnv("ES_HOST", ""),
		// ESUser:     getEnv("ES_USER", ""),
		// ESPassword: getEnv("ES_PASSWORD", ""),
	}

	return cfg, nil
}

// getEnv reads an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
