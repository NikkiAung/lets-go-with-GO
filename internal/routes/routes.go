package routes

import (
	"github.com/NikkiAung/go-fundmentals/internal/app"
	"github.com/go-chi/chi/v5"
)


func SetUpRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()
	// curl localhost:8080/health
	r.Get("/health", app.HealthCheck)
	return r
}