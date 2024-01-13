package main

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"mime"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/go-chi/chi/v5"
	"mehrangcode.ir/resturant/app/database"
	"mehrangcode.ir/resturant/app/router"
	"mehrangcode.ir/resturant/app/utils"
)

// Execute before the service runs.
func init() {
	_ = mime.AddExtensionType(".js", "text/javascript")
	_ = mime.AddExtensionType(".css", "text/css")
}

//go:embed client/*
var spaFiles embed.FS

func main() {
	database.Connect("sqlite")
	// var r *chi.Mux
	r := router.RegisterRoutes().(*chi.Mux)
	r.Get("/api/*", func(w http.ResponseWriter, r *http.Request) {
		utils.ResponseToError(w, errors.New("not found"), http.StatusNotFound)
	})
	r.Handle("/*", SPAHandler())
	log.Println("APP is Running On http://localhost:3000")
	http.ListenAndServe(":3000", r)
}

func SPAHandler() http.HandlerFunc {
	spaFS, err := fs.Sub(spaFiles, "client")
	if err != nil {
		panic(fmt.Errorf("failed getting the sub tree for the site files: %w", err))
	}
	return func(w http.ResponseWriter, r *http.Request) {
		f, err := spaFS.Open(strings.TrimPrefix(path.Clean(r.URL.Path), "/"))
		if err == nil {
			defer f.Close()
		}
		if os.IsNotExist(err) {
			r.URL.Path = "/"
		}
		http.FileServer(http.FS(spaFS)).ServeHTTP(w, r)
	}
}
