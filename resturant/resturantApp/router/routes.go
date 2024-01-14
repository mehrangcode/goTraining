package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"mehrangcode.ir/resturant/app/handlers"
)

func RegisterRoutes() http.Handler {

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Use(middleware.Logger)

	user_handler := handlers.NewUsersHandler()
	r.Route("/api/users", func(r chi.Router) {
		r.Get("/", user_handler.GetAll)
		r.Post("/", user_handler.Create)
		r.Route("/{userId}", func(r chi.Router) {
			r.Put("/", user_handler.Update)
			r.Delete("/", user_handler.Delete)
		})
	})
	food_handler := handlers.NewFoodsHandler()
	r.Route("/api/foods", func(r chi.Router) {
		r.Get("/", food_handler.GetAll)
		r.Post("/", food_handler.Create)
		r.Route("/{foodId}", func(r chi.Router) {
			r.Put("/changeStatus/{status}", food_handler.ChangeStatus)
			r.Put("/", food_handler.Update)
			r.Delete("/", food_handler.Delete)
		})
	})
	return r
}
