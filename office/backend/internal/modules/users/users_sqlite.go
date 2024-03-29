package users

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"mehrangcode.ir/office/internal/types"
	"mehrangcode.ir/office/pkg/database"
	"mehrangcode.ir/office/utils"
)

type UserSqliteRepository struct {
	DB *sqlx.DB
}

func NewSqliteRepo() *UserSqliteRepository {
	return &UserSqliteRepository{
		DB: database.Connection(),
	}
}

func (repo *UserSqliteRepository) GetAll() ([]types.UserViewModel, error) {
	query := `SELECT * FROM users`
	var userList []types.UserViewModel
	err := repo.DB.Select(&userList, query)
	if err != nil {
		return nil, err
	}
	// defer rows.Close()
	// // return users, nil
	// for rows.Next() {
	// 	var u types.UserViewModel
	// 	err = rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	// converted := NewUser(u.ID, u.Name)
	// 	userList = append(userList, u)
	// }
	return userList, nil
}

func (repo *UserSqliteRepository) Create(payload types.UserDTO) (string, error) {
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

func (repo *UserSqliteRepository) Update(userId string, userPayload types.UserDTO) error {
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

func (repo *UserSqliteRepository) Delete(userId string) error {
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
