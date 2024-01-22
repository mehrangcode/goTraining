package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"mehrangcode.ir/resturant/app/utils"
)

func sqliteDbConnect() {
	db, err := sqlx.Open("sqlite3", ":memory:")
	// db, err := sqlx.Open("sqlite3", "data.db")
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
		CREATE TABLE IF NOT EXISTS foods(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description,
			status INTEGER DEFAULT 1,
			photos
		);
		`
	_, err = DB.Exec(query)
	if err != nil {
		panic(err)
	}
	query = `
		CREATE TABLE IF NOT EXISTS foodCategories(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			description,
			status INTEGER DEFAULT 1,
			avatar
		);
		`
	_, err = DB.Exec(query)
	if err != nil {
		panic(err)
	}
	query = `
		CREATE TABLE IF NOT EXISTS food_categories(
			food_id TEXT NOT NULL,
			category_id TEXT NOT NULL
		);
		`
	_, err = DB.Exec(query)
	if err != nil {
		panic(err)
	}
	query = `
		CREATE TABLE IF NOT EXISTS menus(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			description,
			status INTEGER DEFAULT 1
		);
		`
	_, err = DB.Exec(query)
	if err != nil {
		panic(err)
	}
	query = `
		CREATE TABLE IF NOT EXISTS sections(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			menu_id TEXT NOT NULL,
			description,
			status INTEGER DEFAULT 1
		);
		`
	_, err = DB.Exec(query)
	if err != nil {
		panic(err)
	}
	query = `
		CREATE TABLE IF NOT EXISTS section_foods(
			food_id TEXT NOT NULL,
			section_id TEXT NOT NULL,
			price
		);
		`
	_, err = DB.Exec(query)
	if err != nil {
		panic(err)
	}
	query = `
		CREATE TABLE IF NOT EXISTS tables(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			capacity INTEGER NOT NULL,
			photos TEXT,
			status INTEGER DEFAULT 1
		);
		`
	_, err = DB.Exec(query)
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
	query = ` INSERT INTO foodCategories (title,status) values(?,?)`
	_, err = DB.Exec(query, "Fast Food", 1)
	if err != nil {
		panic(err)
	}
	query = ` INSERT INTO foodCategories (title,status) values(?,?)`
	_, err = DB.Exec(query, "Drinks", 1)
	if err != nil {
		panic(err)
	}
	query = ` INSERT INTO foods (name,status) values(?,?)`
	_, err = DB.Exec(query, "Pizza", 1)
	if err != nil {
		panic(err)
	}
	query = ` INSERT INTO foods (name,status) values(?,?)`
	_, err = DB.Exec(query, "Hot dog", 1)
	if err != nil {
		panic(err)
	}
	query = ` INSERT INTO foods (name,status) values(?,?)`
	_, err = DB.Exec(query, "Spagety", 1)
	if err != nil {
		panic(err)
	}
	query = ` INSERT INTO food_categories (food_id,category_id) values(?,?)`
	_, err = DB.Exec(query, 1, 1)
	if err != nil {
		panic(err)
	}
	query = ` INSERT INTO food_categories (food_id,category_id) values(?,?)`
	_, err = DB.Exec(query, 2, 1)
	if err != nil {
		panic(err)
	}
	query = ` INSERT INTO food_categories (food_id,category_id) values(?,?)`
	_, err = DB.Exec(query, 3, 1)
	if err != nil {
		panic(err)
	}
	return err
}
