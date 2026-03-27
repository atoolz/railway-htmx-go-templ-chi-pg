package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/atoolz/railway-htmx-go-templ-chi-pg/internal/database"
	"github.com/atoolz/railway-htmx-go-templ-chi-pg/internal/handlers"
)

func main() {
	ctx := context.Background()

	pool, err := database.Connect(ctx)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	defer pool.Close()

	if err := database.Migrate(ctx, pool); err != nil {
		log.Fatal("Failed to run migrations: ", err)
	}

	h := handlers.New(pool)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", h.Home)
	r.Post("/todos", h.CreateTodo)
	r.Patch("/todos/{id}/toggle", h.ToggleTodo)
	r.Delete("/todos/{id}", h.DeleteTodo)
	r.Get("/health", h.HealthCheck)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
