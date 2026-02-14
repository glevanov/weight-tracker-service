package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"weight-tracker-service/internal/config"
	"weight-tracker-service/internal/handlers"
)

func main() {
	cfg := config.Load()

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{cfg.FrontendURL},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	r.Get("/health-check", handlers.HealthCheck)

	addr := fmt.Sprintf(":%d", cfg.Port)
	fmt.Printf("Server is running on http://localhost%s\n", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
