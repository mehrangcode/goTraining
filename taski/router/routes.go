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

	r.Route("/api/users", func(r chi.Router) {
		userHandler := handlers.NewUsersHandler()
		r.Post("/login", userHandler.Login)
		r.Group(func(r chi.Router) {
			// r.Use(utils.Authenticate)
			r.Get("/", userHandler.GetAll)
			r.Post("/", userHandler.Create)
			r.Route("/{userID}", func(r chi.Router) {
				r.Get("/", userHandler.GetById)
				r.Put("/", userHandler.Update)
				r.Delete("/", userHandler.Delete)
				r.Put("/roles", userHandler.AddRolesToUser)
			})
		})
	})

	r.Route("/api/roles", func(r chi.Router) {
		roleHandler := handlers.NewRoleHandler()
		r.Group(func(r chi.Router) {
			// r.Use(utils.Authenticate)
			r.Post("/", roleHandler.Create)
			r.Get("/", roleHandler.GetAll)
			r.Route("/{roleID}", func(r chi.Router) {
				r.Get("/", roleHandler.GetByID)
				r.Put("/", roleHandler.Update)
				r.Delete("/", roleHandler.Delete)
				r.Put("/permissions", roleHandler.AddPermissionsToRole)
			})
		})
	})

	r.Route("/api/permissions", func(r chi.Router) {
		permissionHandler := handlers.NewPermissionHandler()
		r.Group(func(r chi.Router) {
			// r.Use(utils.Authenticate)
			r.Post("/", permissionHandler.Create)
			r.Get("/", permissionHandler.GetAll)
			r.Route("/{permissionID}", func(r chi.Router) {
				r.Get("/", permissionHandler.GetByID)
				r.Put("/", permissionHandler.Update)
				r.Delete("/", permissionHandler.Delete)
			})
		})
	})
	return r
}
