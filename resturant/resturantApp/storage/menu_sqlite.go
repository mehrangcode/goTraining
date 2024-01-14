package storage

import (
	"github.com/jmoiron/sqlx"
	"mehrangcode.ir/resturant/app/database"
	"mehrangcode.ir/resturant/app/models"
)

type MenuSqliteDB struct {
	DB *sqlx.DB
}

func NewMenuSqliteDB() *MenuSqliteDB {
	return &MenuSqliteDB{
		DB: database.Connection(),
	}
}
func Create(payload models.MenuDTO) {
	query := `INSERT INTO menus (title,description,status) VALUES(?,?,?)`
}
