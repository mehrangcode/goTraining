package issued_letters

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

func (repo *SqliteStorage) GetAll() ([]types.IssuedLetterViewModel, error) {
	query := `SELECT issued_letters.*,
	subjects.label AS subjectName,
	users.name AS operatorName 
	FROM issued_letters
	LEFT OUTER JOIN subjects ON issued_letters.subjectId = subjects.id
	LEFT OUTER JOIN users ON issued_letters.operatorId = users.id;`
	var list []types.IssuedLetterViewModel
	err := repo.DB.Select(&list, query)
	if err != nil {
		return nil, err
	}
	return list, nil
}
func (repo *SqliteStorage) GetById(itemId string) (types.IssuedLetterViewModel, error) {
	query := `SELECT issued_letters.*,
	subjects.label AS subjectName,
	users.name AS operatorName 
	FROM issued_letters
	LEFT OUTER JOIN subjects ON issued_letters.subjectId = subjects.id
	LEFT OUTER JOIN users ON issued_letters.operatorId = users.id
	WHERE issued_letters.id = ?;`
	var item types.IssuedLetterViewModel
	err := repo.DB.Get(&item, query, itemId)
	if err != nil {
		return types.IssuedLetterViewModel{}, err
	}
	return item, nil
}

func (repo *SqliteStorage) Create(payload types.IssuedLetterDTO) (string, error) {
	query := `INSERT INTO issued_letters 
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
		1,
	)
	itemId := ""
	err = row.Scan(&itemId)
	if err != nil {
		return "", err
	}
	return itemId, nil
}

func (repo *SqliteStorage) Update(itemId string, payload types.IssuedLetterDTO) error {
	query := `UPDATE issued_letters SET 
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
		1,
		itemId).Scan(&id)
	if err != nil {
		return errors.New("issued_letter was not found")
	}
	return nil
}

func (repo *SqliteStorage) Delete(subjetcId string) error {
	query := `DELETE FROM issued_letters WHERE id=?`

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
