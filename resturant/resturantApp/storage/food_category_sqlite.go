package storage

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"mehrangcode.ir/resturant/app/database"
	"mehrangcode.ir/resturant/app/models"
)

type FoodCategorySqliteDB struct {
	DB *sqlx.DB
}

func NewFoodCategorySqliteDB() *FoodCategorySqliteDB {
	return &FoodCategorySqliteDB{
		DB: database.Connection(),
	}
}

func (repo *FoodCategorySqliteDB) GetAll() ([]models.FoodCategoryViewModel, error) {
	query := `SELECT * FROM food_categories`
	var list []models.FoodCategoryViewModel
	err := repo.DB.Select(&list, query)
	if err != nil {
		return nil, err
	}
	return list, nil
}
func (repo *FoodCategorySqliteDB) GetById(id string) (models.FoodCategoryViewModel, error) {
	query := `SELECT * FROM food_categories WHERE id=?`
	var item models.FoodCategoryViewModel
	err := repo.DB.Get(&item, query, id)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (repo *FoodCategorySqliteDB) Create(payload models.FoodCategoryDTO) (string, error) {
	query := `INSERT INTO food_categories (title,description,status,avatar) VALUES (?,?,?,?) RETURNING id`
	stmt, err := repo.DB.Prepare(query)
	if err != nil {
		return "", err
	}
	row := stmt.QueryRow(payload.Title, payload.Description, 1, payload.Avatar)
	catId := ""
	err = row.Scan(&catId)
	if err != nil {
		return "", err
	}
	return catId, nil
}

func (repo *FoodCategorySqliteDB) Update(catId string, payload models.FoodCategoryDTO) error {
	query := "UPDATE food_categories SET title=?,description=?,avatar=? WHERE id=? RETURNING id"
	stmt, err := repo.DB.Prepare(query)
	if err != nil {
		return err
	}
	var id string
	err = stmt.QueryRow(payload.Title, payload.Description, payload.Avatar, catId).Scan(&id)
	if err != nil {
		return errors.New("category was not found")
	}
	return nil
}
func (repo *FoodCategorySqliteDB) ChangeStatus(catId string, status int) error {
	query := "UPDATE food_categories SET status=?  WHERE id=? RETURNING id"
	stmt, err := repo.DB.Prepare(query)
	if err != nil {
		return err
	}
	var id string
	err = stmt.QueryRow(status, catId).Scan(&id)
	if err != nil {
		return errors.New("category was not found")
	}
	return nil
}

func (repo *FoodCategorySqliteDB) Delete(catId string) error {
	query := `DELETE FROM food_categories WHERE id=?`

	stmt, err := repo.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(catId)
	if err != nil {
		return err
	}

	return nil
}
