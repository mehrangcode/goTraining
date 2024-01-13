package database

import (
	"github.com/jmoiron/sqlx"
)

func Connection() *sqlx.DB {
	return DB
}
