package database

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

var DB *sql.DB
var DBx *sqlx.DB
