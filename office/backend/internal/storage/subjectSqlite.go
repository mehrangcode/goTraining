package storage

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"mehrangcode.ir/office/internal/types"
	"mehrangcode.ir/office/pkg/database"
)

type SubjectSqliteRepository struct {
	DB *sqlx.DB
}

func NewSubjectSqliteRepo() *SubjectSqliteRepository {
	return &SubjectSqliteRepository{
		DB: database.Connection(),
	}
}

func (repo *SubjectSqliteRepository) GetAll() ([]types.SubjectViewModel, error) {
	query := `SELECT * FROM subjects`
	var subjectList []types.SubjectViewModel
	err := repo.DB.Select(&subjectList, query)
	if err != nil {
		return nil, err
	}
	return subjectList, nil
}

func (repo *SubjectSqliteRepository) Create(payload types.SubjectDTO) (string, error) {
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

func (repo *SubjectSqliteRepository) Update(payload types.SubjectDTO) error {
	query := "UPDATE subjects SET label=? WHERE id=? RETURNING id"
	stmt, err := repo.DB.Prepare(query)
	if err != nil {
		return err
	}
	var id string
	err = stmt.QueryRow(payload.Label, payload.ID).Scan(&id)
	if err != nil {
		return errors.New("subject was not found")
	}
	return nil
}

func (repo *SubjectSqliteRepository) Delete(subjectId string) error {
	query := `DELETE FROM subjects WHERE id=?`

	stmt, err := repo.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(subjectId)
	if err != nil {
		return err
	}

	return nil
}
