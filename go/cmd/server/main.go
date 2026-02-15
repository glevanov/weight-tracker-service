package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"weight-tracker-service/internal/config"
	"weight-tracker-service/internal/database"
	"weight-tracker-service/internal/handlers"
)

func main() {
	cfg := config.Load()

	if err := database.Connect(cfg.ConnectionURI, cfg.DBName); err != nil {
		fmt.Printf("Database connection error: %v\n", err)
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

	r.Get("/health-check", handlers.HealthCheck)
	r.Get("/weights", handlers.GetWeights)
	r.Post("/weights", handlers.AddWeight)

	addr := fmt.Sprintf(":%d", cfg.Port)
	fmt.Printf("Server is running on http://localhost%s\n", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
