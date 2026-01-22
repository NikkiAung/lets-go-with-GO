package routes

import (
	"github.com/NikkiAung/go-fundmentals/internal/app"
	"github.com/go-chi/chi/v5"
)


func SetUpRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()
	// curl localhost:8080/health
	r.Get("/health", app.HealthCheck)

	// curl localhost:8080/posts/1
	r.Get("/posts/{id}", app.PostHandler.HandleGetPostById)

	// curl -X POST localhost:8080/posts
	r.Post("/posts", app.PostHandler.HandleCreatePost)

	r.Get("/posts", app.PostHandler.HandleGetPosts)

	return r
}