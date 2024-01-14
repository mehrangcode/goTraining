package storage

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"mehrangcode.ir/resturant/app/database"
	"mehrangcode.ir/resturant/app/models"
)

type FoodSqliteDB struct {
	DB *sqlx.DB
}

func NewFoodSqliteDB() *FoodSqliteDB {
	return &FoodSqliteDB{
		DB: database.Connection(),
	}
}

func (repo *FoodSqliteDB) GetAll() ([]models.FoodViewModel, error) {
	query := `SELECT * FROM foods`
	var foodList []models.FoodViewModel
	err := repo.DB.Select(&foodList, query)
	if err != nil {
		return nil, err
	}
	return foodList, nil
}
func (repo *FoodSqliteDB) GetById(id string) (models.FoodViewModel, error) {
	query := `SELECT * FROM foods WHERE id=?`
	var food models.FoodViewModel
	err := repo.DB.Get(&food, query, id)
	if err != nil {
		return food, err
	}
	return food, nil
}

func (repo *FoodSqliteDB) Create(payload models.FoodDTO) (string, error) {
	query := `INSERT INTO foods (name,description,status,photos) VALUES (?,?,?,?) RETURNING id`
	stmt, err := repo.DB.Prepare(query)
	if err != nil {
		return "", err
	}
	row := stmt.QueryRow(payload.Name, payload.Description, 1, payload.Photos)
	foodId := ""
	err = row.Scan(&foodId)
	if err != nil {
		return "", err
	}
	return foodId, nil
}

func (repo *FoodSqliteDB) Update(foodId string, payload models.FoodDTO) error {
	query := "UPDATE foods SET name=?,description=?,photos=? WHERE id=? RETURNING id"
	stmt, err := repo.DB.Prepare(query)
	if err != nil {
		return err
	}
	var id string
	err = stmt.QueryRow(payload.Name, payload.Description, payload.Photos, foodId).Scan(&id)
	if err != nil {
		return errors.New("food was not found")
	}
	return nil
}
func (repo *FoodSqliteDB) ChangeStatus(foodId string, status int) error {
	query := "UPDATE foods SET status=?  WHERE id=? RETURNING id"
	stmt, err := repo.DB.Prepare(query)
	if err != nil {
		return err
	}
	var id string
	err = stmt.QueryRow(status, foodId).Scan(&id)
	if err != nil {
		return errors.New("food was not found")
	}
	return nil
}

func (repo *FoodSqliteDB) Delete(foodId string) error {
	query := `DELETE FROM foods WHERE id=?`

	stmt, err := repo.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(foodId)
	if err != nil {
		return err
	}

	return nil
}
