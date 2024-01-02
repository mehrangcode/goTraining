package router

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	user_api "mehrangcode.ir/office/internal/api"
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
	FileServer(r)
	r.Get("/users", user_api.GetAll)
	r.Post("/users", user_api.Create)
	return r
}

func FileServer(router *chi.Mux) {
	root := "./public"
	fs := http.FileServer(http.Dir(root))

	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(root + r.URL.Path); os.IsNotExist(err) {
			http.StripPrefix(r.URL.Path, fs).ServeHTTP(w, r)
		} else {
			fs.ServeHTTP(w, r)
		}
	})
}

// // FileServer is serving static files
// func FileServer(r chi.Router, public string, static string) {

// 	if strings.ContainsAny(public, "{}*") {
// 		panic("FileServer does not permit URL parameters.")
// 	}

// 	root, _ := filepath.Abs(static)
// 	if _, err := os.Stat(root); os.IsNotExist(err) {
// 		panic("Static Documents Directory Not Found")
// 	}

// 	fs := http.StripPrefix(public, http.FileServer(http.Dir(root)))

// 	if public != "/" && public[len(public)-1] != '/' {
// 		r.Get(public, http.RedirectHandler(public+"/", 301).ServeHTTP)
// 		public += "/"
// 	}

// 	r.Get(public+"*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		file := strings.Replace(r.URL.Path, public, "/", 1)
// 		if _, err := os.Stat(root + file); os.IsNotExist(err) {
// 			http.ServeFile(w, r, path.Join(root, "index.html"))
// 			return
// 		}
// 		fs.ServeHTTP(w, r)
// 	}))
// }
