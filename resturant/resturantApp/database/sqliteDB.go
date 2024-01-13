package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"mehrangcode.ir/resturant/app/utils"
)

func sqliteDbConnect() {
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
		return
	}
	DB = db
	err = MigrateDB()
	if err != nil {
		panic(err)
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	log.Println("sqlite DB is Connected")
}

func MigrateDB() error {
	query := `
		CREATE TABLE IF NOT EXISTS users(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL,
			password TEXT NOT NULL
		);
		`
	_, err := DB.Exec(query)
	if err != nil {
		panic(err)
	}
	query = ` INSERT INTO users (name,email,password) values(?,?,?)`
	hash, err := utils.HashingPassword("1234")
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(query, "Mehran Ganji", "Mehran@mail.com", hash)
	if err != nil {
		panic(err)
	}
	return err
}
