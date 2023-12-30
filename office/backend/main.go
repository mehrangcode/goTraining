package main

import (
	"log"
	"net/http"

	"mehrangcode.ir/office/pkg/database"
	"mehrangcode.ir/office/pkg/router"
)

func main() {
	database.Connect("sqlite")
	r := router.RegisterRoutes()
	log.Println("APP is Running On http://localhost:3000")
	http.ListenAndServe(":3000", r)
}
