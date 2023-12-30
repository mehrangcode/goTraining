package main

import (
	"net/http"

	"mehrangcode.ir/office/pkg/database"
	"mehrangcode.ir/office/pkg/router"
)

func main() {
	database.Connect("sqlite")
	r := router.RegisterRoutes()
	http.ListenAndServe(":3000", r)
}
