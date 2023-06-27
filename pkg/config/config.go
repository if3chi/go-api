package config

import "os"

type AppConfig struct {
	Name    string
	Version string
	Port    string
}

type HTTPConfig struct {
	Content string
	Problem string
}

type Config struct {
	App  *AppConfig
	HTTP *HTTPConfig
}

func Load() *Config {
	return &Config{
		App: &AppConfig{
			Name:    env("APP_NAME", "Go Api"),
			Version: env("APP_VERSION", "2023.1.0"),
			Port:    env("APP_PORT", "8080"),
		},
		HTTP: &HTTPConfig{
			Content: env("HTTP_CONTENT_TYPE", "application/json"),
			Problem: env("HTTP_PROBLEM", "application/problem+json"),
		},
	}
}

func env(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
