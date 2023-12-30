package users

import (
	"database/sql"

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

// func ConvertToRepoModel(u ViewModel) UserViewModel {
// 	return UserViewModel{
// 		ID:   u.person.ID,
// 		Name: u.person.Name,
// 	}
// }
