package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func sqliteDbConnect() {
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
		return
	}
	DBx = db
	err = MigrateDB()
	if err != nil {
		panic(err)
	}
	DBx.SetMaxOpenConns(10)
	DBx.SetMaxIdleConns(5)
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
	_, err := DBx.Exec(query)
	if err != nil {
		panic(err)
	}
	query = `
		CREATE TABLE IF NOT EXISTS income_letters(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			number INTEGER NOT NULL,
			title TEXT,
			content TEXT,
			subjectId TEXT DEFAULT 1,
			created_at DATE DEFAULT CURRENT_DATE,
			owner TEXT,
			destination TEXT,
			operatorId TEXT,
			status INTEGER DEFAULT 1
		);
		`
	_, err = DBx.Exec(query)
	if err != nil {
		panic(err)
	}
	query = `
	CREATE TABLE subjects (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		label TEXT NOT NULL,
		archive BOOLEAN DEFAULT FALSE
		);
		`
	_, err = DBx.Exec(query)
	if err != nil {
		panic(err)
	}

	return err
}
