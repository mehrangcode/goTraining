package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func postgresDbConnect() {
	dns := "postgresql://root:TfDOLXcFMh6glxXAK2IpHZzz@kilimanjaro.liara.cloud:34267/postgres?sslmode=disable"

	db, MyError := sql.Open("postgres", dns)
	if MyError != nil {
		log.Fatal(MyError)
		return
	}
	DB = db
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	log.Println("postgres DB is Connected")
}
