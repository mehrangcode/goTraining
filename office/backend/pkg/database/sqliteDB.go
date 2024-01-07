package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func sqliteDbConnect() {
	db, err := sql.Open("sqlite3", "data.db")
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
	query = `
		CREATE TABLE IF NOT EXISTS income_letters(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			number INTEGER NOT NULL,
			title TEXT,
			subjectId TEXT DEFAULT 1,
			owner TEXT,
			destination TEXT,
			status INTEGER DEFAULT 1,
			content TEXT,
			operatorId TEXT,
			created_at DATE DEFAULT CURRENT_DATE
		);
		`
	_, err = DB.Exec(query)
	if err != nil {
		panic(err)
	}
	// query = `
	// 	CREATE TABLE IF NOT EXISTS letters(
	// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 		title TEXT NOT NULL,
	// 		content TEXT NOT NULL,
	// 		ownerId TEXT NOT NULL,
	// 	);
	// 	`
	// _, err = DB.Exec(query)
	// if err != nil {
	// 	panic(err)
	// }
	return err
}
