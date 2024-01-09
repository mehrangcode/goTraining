package income_letters

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

func (repo *SqliteStorage) GetAll() ([]types.IncomeLetterViewModel, error) {
	query := `SELECT * FROM income_letters`
	var list []types.IncomeLetterViewModel
	err := repo.DB.Select(&list, query)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (repo *SqliteStorage) Create(payload types.IncomeLetterDTO) (string, error) {
	query := `INSERT INTO income_letters 
	(number,title,content,subjectId,owner,destination,status,operatorId) 
	values($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id`
	stmt, err := repo.DB.Prepare(query)
	if err != nil {
		return "", err
	}
	row := stmt.QueryRow(
		payload.Number,
		payload.Title,
		payload.Content,
		payload.SubjectId,
		payload.Owner,
		payload.Destination,
		payload.Status,
		payload.OperatorId,
	)
	itemId := ""
	err = row.Scan(&itemId)
	if err != nil {
		return "", err
	}
	return itemId, nil
}

func (repo *SqliteStorage) Update(itemId string, payload types.IncomeLetterDTO) error {
	query := `UPDATE income_letters SET 
	title=?,
	content=?,
	subjectId=?,
	owner=?,
	destination=?,
	status=?,
	operatorId=?
	WHERE id=? 
	RETURNING id`
	stmt, err := repo.DB.Prepare(query)
	if err != nil {
		return err
	}
	var id string
	err = stmt.QueryRow(
		payload.Title,
		payload.Content,
		payload.SubjectId,
		payload.Owner,
		payload.Destination,
		payload.Status,
		payload.OperatorId,
		itemId).Scan(&id)
	if err != nil {
		return errors.New("income_letter was not found")
	}
	return nil
}

func (repo *SqliteStorage) Delete(subjetcId string) error {
	query := `DELETE FROM income_letters WHERE id=?`

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
