package subjects

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"mehrangcode.ir/office/internal/types"
	"mehrangcode.ir/office/pkg/database"
)

type SqliteStorage struct {
	DB *sqlx.DB
}

func InitialSqliteStorage() *SqliteStorage {
	return &SqliteStorage{
		DB: database.Connection(),
	}
}

func (repo *SqliteStorage) GetAll() ([]types.SubjectViewModel, error) {
	query := `SELECT * FROM subjects`
	var list []types.SubjectViewModel
	err := repo.DB.Select(&list, query)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (repo *SqliteStorage) Create(payload types.SubjectDTO) (string, error) {
	query := `INSERT INTO subjects (label) VALUES (?) RETURNING id`
	stmt, err := repo.DB.Prepare(query)
	if err != nil {
		return "", err
	}
	row := stmt.QueryRow(payload.Label)
	itemId := ""
	err = row.Scan(&itemId)
	if err != nil {
		return "", err
	}
	return itemId, nil
}

func (repo *SqliteStorage) Update(itemId string, payload types.SubjectDTO) error {
	query := "UPDATE subjects SET label=? WHERE id=? RETURNING id"
	stmt, err := repo.DB.Prepare(query)
	if err != nil {
		return err
	}
	var id string
	err = stmt.QueryRow(payload.Label, itemId).Scan(&id)
	if err != nil {
		return errors.New("subject was not found")
	}
	return nil
}

func (repo *SqliteStorage) Delete(subjetcId string) error {
	query := `DELETE FROM subjects WHERE id=?`

	stmt, err := repo.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(subjetcId)
	if err != nil {
		return err
	}

	return nil
}
