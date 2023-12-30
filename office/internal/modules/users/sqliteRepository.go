package users

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
	"mehrangcode.ir/office/pkg/database"
)

type UserSqliteRepository struct {
	DB *sql.DB
}

func NewSqliteRepo() *UserSqliteRepository {
	return &UserSqliteRepository{
		DB: database.Connection(),
	}
}

func (repo *UserSqliteRepository) GetAll() ([]ViewModel, error) {
	query := `SELECT * FROM users`
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var userList []ViewModel
	// return users, nil
	for rows.Next() {
		var u ViewModel
		err = rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password)
		if err != nil {
			return nil, err
		}
		// converted := NewUser(u.ID, u.Name)
		userList = append(userList, u)
	}
	return userList, nil
}

func (repo *UserSqliteRepository) Create(payload DTO) (string, error) {
	query := `INSERT INTO users (name,email,password) VALUES ($1, $2, $3) RETURNING id`
	stmt, err := repo.DB.Prepare(query)
	if err != nil {
		return "", err
	}
	hasingPassword, err := hashingPassword(payload.Password)
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

func hashingPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hash), err
}
