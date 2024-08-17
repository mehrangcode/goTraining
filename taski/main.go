package main

import (
	"net/http"

	"mehrang.ir/taski/database"
	"mehrang.ir/taski/router"
)

func main() {
	database.InitDatabase()
	r := router.RegisterRoutes()
	http.ListenAndServe(":3000", r)
}
