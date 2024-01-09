package database

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func ConnectionDB() *sql.DB {
	return DB
}
func Connection() *sqlx.DB {
	return DBx
}
