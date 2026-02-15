package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port          int
	FrontendURL   string
	ConnectionURI string
	DBName        string
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

	connectionURI := os.Getenv("CONNECTION_URI")
	if connectionURI == "" {
		connectionURI = "mongodb://localhost:27017"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "weight-tracker"
	}

	return &Config{
		Port:          port,
		FrontendURL:   frontendURL,
		ConnectionURI: connectionURI,
		DBName:        dbName,
	}
}
