package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"mehrang.ir/taski/handlers"
)

func RegisterRoutes() http.Handler {

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Use(middleware.Logger)

	userHandlers := handlers.NewUsersHandler()

	r.Route("/api/users", func(r chi.Router) {
		r.Get("/", userHandlers.GetAll)
		r.Post("/", userHandlers.Create)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", userHandlers.GetById)
			r.Put("/", userHandlers.Update)
			r.Delete("/", userHandlers.Delete)
		})
	})
	return r
}
