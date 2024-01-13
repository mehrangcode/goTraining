package storage

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"mehrangcode.ir/resturant/app/database"
	"mehrangcode.ir/resturant/app/models"
	"mehrangcode.ir/resturant/app/utils"
)

type UserSqliteDB struct {
	DB *sqlx.DB
}

func NewUserSqliteDB() *UserSqliteDB {
	return &UserSqliteDB{
		DB: database.Connection(),
	}
}

func (repo *UserSqliteDB) GetAll() ([]models.UserViewModel, error) {
	query := `SELECT * FROM users`
	var userList []models.UserViewModel
	err := repo.DB.Select(&userList, query)
	if err != nil {
		return nil, err
	}
	return userList, nil
}
func (repo *UserSqliteDB) GetById(id string) (models.UserViewModel, error) {
	query := `SELECT * FROM users WHERE id=?`
	var user models.UserViewModel
	err := repo.DB.Get(&user, query, id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (repo *UserSqliteDB) Create(payload models.UserDTO) (string, error) {
	query := `INSERT INTO users (name,email,password) VALUES ($1, $2, $3) RETURNING id`
	stmt, err := repo.DB.Prepare(query)
	if err != nil {
		return "", err
	}
	hasingPassword, err := utils.HashingPassword(payload.Password)
	if err != nil {
		return "", err
	}
	row := stmt.QueryRow(payload.Name, payload.Email, hasingPassword)
	userId := ""
	err = row.Scan(&userId)
	if err != nil {
		return "", err
	}
	return userId, nil
}

func (repo *UserSqliteDB) Update(userId string, userPayload models.UserDTO) error {
	query := "UPDATE users SET name=?, email=? WHERE id=? RETURNING id"
	stmt, err := repo.DB.Prepare(query)
	if err != nil {
		return err
	}
	var id string
	err = stmt.QueryRow(userPayload.Name, userPayload.Email, userId).Scan(&id)
	if err != nil {
		return errors.New("user was not found")
	}
	return nil
}

func (repo *UserSqliteDB) Delete(userId string) error {
	query := `DELETE FROM users WHERE id=?`

	stmt, err := repo.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userId)
	if err != nil {
		return err
	}

	return nil
}
