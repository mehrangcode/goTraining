package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"mehrangcode.ir/office/internal/modules/users"
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
	r.Get("/users", users.GetAll)
	fileServer := http.FileServer(http.Dir("./dist"))
	assetsFiles := http.FileServer(http.Dir("./dist/asstest"))
	r.Handle("/dist/*", http.StripPrefix("/dist/", fileServer))
	r.Handle("/assets/*", http.StripPrefix("/assets/", assetsFiles))
	return r
}
