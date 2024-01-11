package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"mehrangcode.ir/office/internal/modules/income_letters"
	"mehrangcode.ir/office/internal/modules/issued_letters"
	"mehrangcode.ir/office/internal/modules/subjects"
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
	// FileServer(r)

	user_handler := users.NewHandler(users.NewSqliteRepo())
	r.Route("/users", func(r chi.Router) {
		r.Get("/", user_handler.GetAll)
		r.Post("/", user_handler.Create)
		r.Route("/{userId}", func(r chi.Router) {
			r.Put("/", user_handler.Update)
			r.Delete("/", user_handler.Delete)
		})
	})

	subject_handler := subjects.NewHandlers(subjects.InitialSqliteStorage())
	r.Route("/subjects", func(r chi.Router) {
		r.Get("/", subject_handler.GetAll)
		r.Post("/", subject_handler.Create)
		r.Route("/{subjectId}", func(r chi.Router) {
			r.Put("/", subject_handler.Update)
			r.Delete("/", subject_handler.Delete)
		})
	})

	// LETERS INCOME
	issued_letters_handlers := issued_letters.NewHandlers(issued_letters.InitialSqliteStorage())
	r.Route("/letters/issued", func(r chi.Router) {
		r.Get("/", issued_letters_handlers.GetAll)
		r.Post("/", issued_letters_handlers.Create)
		r.Route("/{letterId}", func(r chi.Router) {
			r.Get("/", issued_letters_handlers.GetById)
			r.Put("/", issued_letters_handlers.Update)
			r.Delete("/", issued_letters_handlers.Delete)
		})
	})

	// LETERS INCOME
	income_letters_handlers := income_letters.NewHandlers(income_letters.InitialSqliteStorage())
	r.Route("/letters/income", func(r chi.Router) {
		r.Get("/", income_letters_handlers.GetAll)
		r.Post("/", income_letters_handlers.Create)
		r.Route("/{letterId}", func(r chi.Router) {
			r.Put("/", income_letters_handlers.Update)
			r.Delete("/", income_letters_handlers.Delete)
		})
	})
	return r
}

// func FileServer(router *chi.Mux) {
// 	root := "./public"
// 	fs := http.FileServer(http.Dir(root))

// 	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
// 		if _, err := os.Stat(root + r.URL.Path); os.IsNotExist(err) {
// 			http.StripPrefix(r.URL.Path, fs).ServeHTTP(w, r)
// 		} else {
// 			fs.ServeHTTP(w, r)
// 		}
// 	})
// }

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
