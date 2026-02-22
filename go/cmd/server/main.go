package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"weight-tracker-service/internal/auth"
	"weight-tracker-service/internal/config"
	"weight-tracker-service/internal/database"
	"weight-tracker-service/internal/handlers"
	"weight-tracker-service/internal/logger"
)

func main() {
	cfg := config.Load()

	if err := database.Connect(cfg.ConnectionURI, cfg.DBName); err != nil {
		logger.Error("database connection error", "error", err)
		return
	}
	defer database.Disconnect(context.Background())

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{cfg.FrontendURL},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	// Public routes
	r.Get("/health-check", handlers.HealthCheck)
	r.Post("/login", handlers.Login(cfg))

	// Protected routes (require authentication)
	r.Group(func(r chi.Router) {
		r.Use(auth.Middleware(cfg.JWTSecret))
		r.Get("/weights", handlers.GetWeights)
		r.Post("/weights", handlers.AddWeight)
		r.Get("/session-check", handlers.SessionCheck)
	})

	addr := fmt.Sprintf(":%d", cfg.Port)
	logger.Info("server starting", "port", cfg.Port)

	if err := http.ListenAndServe(addr, r); err != nil {
		logger.Error("server error", "error", err)
	}
}
