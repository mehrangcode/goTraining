package storage

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"mehrangcode.ir/resturant/app/database"
	"mehrangcode.ir/resturant/app/models"
)

type ReservationSqliteDB struct {
	DB *sqlx.DB
}

func NewReservationSqliteDB() *ReservationSqliteDB {
	return &ReservationSqliteDB{
		DB: database.Connection(),
	}
}

func (repo *ReservationSqliteDB) GetAll() ([]models.ReservationViewModel, error) {
	query := `SELECT 
	r.id as "id", 
	r.guests as "guests", 
	r.date as "date",
	user.id as "user.id", 
	user.name as "user.name", 
	user.email as "user.email"
	FROM reservations r
	JOIN users user ON r.user_id = user.id`
	var list []models.ReservationViewModel
	err := repo.DB.Select(&list, query)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (repo *ReservationSqliteDB) GetById(reservationId string) (models.ReservationViewModel, error) {
	query := `SELECT 
	r.id as "id", 
	r.guests as "guests", 
	r.date as "date",
	user.id as "user.id", 
	user.name as "user.name", 
	user.email as "user.email"
	FROM reservations r
	JOIN users user ON r.user_id = user.id
	WHERE r.id=?`
	var item models.ReservationViewModel
	err := repo.DB.Get(&item, query, reservationId)
	if err != nil {
		return item, err
	}
	return item, nil
}
func (repo *ReservationSqliteDB) GetByUserId(userId string) ([]models.ReservationViewModel, error) {
	query := `SELECT 
	r.id as "id", 
	r.guests as "guests", 
	r.date as "date",
	user.id as "user.id", 
	user.name as "user.name", 
	user.email as "user.email"
	FROM reservations r
	JOIN users user ON r.user_id = user.id
	WHERE r.user_id=?`
	var list []models.ReservationViewModel
	err := repo.DB.Select(&list, query, userId)
	if err != nil {
		return list, err
	}
	return list, nil
}

func (repo *ReservationSqliteDB) Create(payload models.ReservationDTO) (string, error) {
	query := `INSERT INTO reservations (user_id,guests,date,status) VALUES (:user_id,:guests,:date,:status) RETURNING id`
	stmt, err := repo.DB.Prepare(query)
	if err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}
	row := stmt.QueryRow(payload.UserID, payload.Guests, payload.Date, 1)
	userId := ""
	err = row.Scan(&userId)
	if err != nil {
		return "", err
	}
	return userId, nil
}

func (repo *ReservationSqliteDB) Update(reservationID string, payload models.ReservationDTO) error {
	query := "UPDATE reservations SET guests=?,date=? WHERE id=? RETURNING id"
	stmt, err := repo.DB.Prepare(query)
	if err != nil {
		return err
	}
	var reservationId string
	err = stmt.QueryRow(payload.Guests, payload.Date, reservationID).Scan(&reservationId)
	if err != nil {
		return errors.New("reservation was not found")
	}
	return nil
}

func (repo *ReservationSqliteDB) ChangeStatus(reservationId string, status uint) error {
	query := "UPDATE reservations SET status=? WHERE id=? RETURNING id"
	stmt, err := repo.DB.Prepare(query)
	if err != nil {
		return err
	}
	row := stmt.QueryRow(status)
	_reservationId := ""
	err = row.Scan(&_reservationId)
	if err != nil {
		return err
	}

	return nil
}

func (repo *ReservationSqliteDB) Delete(reservationID string) error {
	query := `DELETE FROM reservations WHERE id=?`
	stmt, err := repo.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(reservationID)
	if err != nil {
		return err
	}
	return nil
}
