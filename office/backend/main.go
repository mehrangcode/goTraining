package main

import (
	"log"
	"mime"
	"net/http"

	"mehrangcode.ir/office/pkg/database"
	"mehrangcode.ir/office/pkg/router"
)

// Execute before the service runs.
func init() {
	_ = mime.AddExtensionType(".js", "text/javascript")
	_ = mime.AddExtensionType(".css", "text/css")
}

func main() {
	database.Connect("sqlite")
	r := router.RegisterRoutes()
	log.Println("APP is Running On http://localhost:3000")
	http.ListenAndServe(":3000", r)
}
