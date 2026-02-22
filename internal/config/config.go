package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port            int
	FrontendURL     string
	ConnectionURI   string
	DBName          string
	JWTSecret       string
	SessionDuration time.Duration
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
		panic(fmt.Errorf("CONNECTION_URI environment variable is required"))
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "weight-tracker"
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		panic(fmt.Errorf("JWT_SECRET environment variable is required"))
	}

	// 5 years session duration
	sessionDuration := 5 * 365 * 24 * time.Hour

	return &Config{
		Port:            port,
		FrontendURL:     frontendURL,
		ConnectionURI:   connectionURI,
		DBName:          dbName,
		JWTSecret:       jwtSecret,
		SessionDuration: sessionDuration,
	}
}
