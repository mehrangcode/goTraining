package storage

import (
	"github.com/jmoiron/sqlx"
	"mehrangcode.ir/office/internal/types"
	"mehrangcode.ir/office/pkg/database"
)

type IncomeLettersSqliteRepository struct {
	DB *sqlx.DB
}

func NewIncomSqliteRepository() *IncomeLettersSqliteRepository {
	return &IncomeLettersSqliteRepository{
		DB: database.Connection(),
	}
}

func (repo *IncomeLettersSqliteRepository) GetAll() ([]types.IncomeLetterViewModel, error) {
	query := `SELECT * FROM income_letters`
	var letters []types.IncomeLetterViewModel
	err := repo.DB.Select(&letters, query)
	if err != nil {
		return nil, err
	}
	return letters, nil
}

func (repo *IncomeLettersSqliteRepository) GetById(letterId string) (types.IncomeLetterViewModel, error) {

	var letter types.IncomeLetterViewModel
	return letter, nil
}

func (repo *IncomeLettersSqliteRepository) Create(payload types.IncomeLetterDTO) (string, error) {

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
	var letterId string
	err = row.Scan(&letterId)
	if err != nil {
		return "", err
	}
	return letterId, nil
}

func (repo *IncomeLettersSqliteRepository) Update(letterId string, payload types.IncomeLetterDTO) error {

	return nil
}

func (repo *IncomeLettersSqliteRepository) Delete(letterId string) error {

	return nil
}
