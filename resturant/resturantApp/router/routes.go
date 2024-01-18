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
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
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
			r.Put("/", food_handler.Update)
			r.Patch("/changeStatus/{status}", food_handler.ChangeStatus)
			r.Delete("/", food_handler.Delete)
		})
	})
	food_category_handler := handlers.NewFoodCategoriesHandler()
	r.Route("/api/foodCategories", func(r chi.Router) {
		r.Get("/", food_category_handler.GetAll)
		r.Post("/", food_category_handler.Create)
		r.Route("/{catId}", func(r chi.Router) {
			r.Put("/", food_category_handler.Update)
			r.Patch("/changeStatus/{status}", food_category_handler.ChangeStatus)
			r.Delete("/", food_category_handler.Delete)
		})
	})
	menu_handler := handlers.NewMenusHandler()
	r.Route("/api/menus", func(r chi.Router) {
		r.Get("/", menu_handler.GetAll)
		r.Post("/", menu_handler.Create)
		r.Route("/{menuId}", func(r chi.Router) {
			// r.Put("/", menu_handler.Update)
			r.Delete("/", menu_handler.Delete)
		})
	})
	return r
}
