package storage

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"mehrangcode.ir/resturant/app/database"
	"mehrangcode.ir/resturant/app/models"
)

type TableSqliteDB struct {
	DB *sqlx.DB
}

func NewTableSqliteDB() *TableSqliteDB {
	return &TableSqliteDB{
		DB: database.Connection(),
	}
}

func (repo *TableSqliteDB) GetAll() ([]models.TableViewModel, error) {
	query := `SELECT * FROM tables`
	var list []models.TableViewModel
	err := repo.DB.Select(&list, query)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (repo *TableSqliteDB) GetById(tableId string) (models.TableViewModel, error) {
	query := `SELECT * FROM tables WHERE id=?`
	var item models.TableViewModel
	err := repo.DB.Get(&item, query, tableId)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (repo *TableSqliteDB) Create(payload models.TableDTO) (string, error) {
	query := `INSERT INTO tables (name,capacity,photos,status) VALUES (:name,:capacity,:photos,:status) RETURNING id`
	stmt, err := repo.DB.Prepare(query)
	if err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}
	row := stmt.QueryRow(payload.Name, payload.Capacity, payload.Photos, 1)
	userId := ""
	err = row.Scan(&userId)
	if err != nil {
		return "", err
	}
	return userId, nil
}

func (repo *TableSqliteDB) Update(tableID string, payload models.TableDTO) error {
	query := "UPDATE tables SET name=?,capacity=?,photos=? WHERE id=? RETURNING id"
	stmt, err := repo.DB.Prepare(query)
	if err != nil {
		return err
	}
	var tableId string
	err = stmt.QueryRow(payload.Name, payload.Capacity, payload.Photos, tableID).Scan(&tableId)
	if err != nil {
		return errors.New("table was not found")
	}
	return nil
}

func (repo *TableSqliteDB) ChangeStatus(tableId string, status uint) error {
	query := "UPDATE tables SET status=? WHERE id=? RETURNING id"
	stmt, err := repo.DB.Prepare(query)
	if err != nil {
		return err
	}
	row := stmt.QueryRow(status)
	_tableId := ""
	err = row.Scan(&_tableId)
	if err != nil {
		return err
	}

	return nil
}

func (repo *TableSqliteDB) Delete(tableID string) error {
	query := `DELETE FROM tables WHERE id=?`
	stmt, err := repo.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(tableID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *TableSqliteDB) Reservation(tableId string, payload models.ReservationDTO) error {

	return nil
}
