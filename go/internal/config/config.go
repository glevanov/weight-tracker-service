package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port        int
	FrontendURL string
}

func Load() *Config {
	port := 3000
	if p := os.Getenv("PORT"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil {
			port = parsed
		}
	}

	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:5173"
	}

	return &Config{
		Port:        port,
		FrontendURL: frontendURL,
	}
}
