package storage

import (
	"database/sql"
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
	query := `SELECT 
	f.id, f.name, f.description,
	c.id AS "catID", 
	c.title AS "catTitle"
	FROM foods f
	LEFT JOIN food_categories fc ON f.id = fc.food_id
	LEFT JOIN foodCategories c ON fc.category_id = c.id
	ORDER BY f.id DESC`
	rows, err := repo.DB.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	foods := make(map[string]*models.FoodViewModel)
	for rows.Next() {
		var (
			id          string
			name        string
			description sql.NullString
			catID       sql.NullString
			catTitle    sql.NullString
		)
		err := rows.Scan(&id, &name, &description, &catID, &catTitle)
		if err != nil {
			return nil, err
		}
		food, exsist := foods[id]
		if !exsist {
			food = &models.FoodViewModel{
				ID:          id,
				Name:        name,
				Description: &description.String,
			}
			foods[id] = food
		}
		var category *models.FoodCategoryViewModel
		if catID.Valid {
			for i := range food.Categories {
				if food.Categories[i].ID == catID.String {
					category = &food.Categories[i]
					break
				}
			}
			if category == nil {
				category = &models.FoodCategoryViewModel{
					ID:    catID.String,
					Title: catTitle.String,
				}
			}
			food.Categories = append(food.Categories, *category)
		}
	}
	var foodList []models.FoodViewModel
	for _, f := range foods {
		foodList = append(foodList, *f)
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
	tx, err := repo.DB.Beginx()
	if err != nil {
		return "", err
	}
	defer tx.Rollback()
	foodID, err := insertFood(tx, payload)
	if err != nil {
		return "", err
	}
	for _, catID := range payload.Categories {
		err = insertFoodCategories(tx, foodID, catID)
		if err != nil {
			return "", err
		}
	}
	tx.Commit()
	return foodID, nil
}

func (repo *FoodSqliteDB) Update(foodId string, payload models.FoodDTO) error {
	tx, err := repo.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	query := "UPDATE foods SET name=?,description=?,photos=? WHERE id=? RETURNING id"
	stmtx, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	var id string
	err = stmtx.QueryRow(payload.Name, payload.Description, payload.Photos, foodId).Scan(&id)
	if err != nil {
		return errors.New("food was not found")
	}
	deleteQuery := "DELETE FROM food_categories WHERE food_id = ?"
	stmtx, err = tx.Prepare(deleteQuery)
	if err != nil {
		return err
	}
	defer stmtx.Close()
	_, err = stmtx.Exec(foodId)
	if err != nil {
		return err
	}
	for _, catID := range payload.Categories {
		err = insertFoodCategories(tx, foodId, catID)
		if err != nil {
			return err
		}
	}
	tx.Commit()
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
	tx, err := repo.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	deleteQuery := "DELETE FROM food_categories WHERE food_id = ?"
	stmtx, err := tx.Prepare(deleteQuery)
	if err != nil {
		return err
	}
	defer stmtx.Close()
	_, err = stmtx.Exec(foodId)
	if err != nil {
		return err
	}
	query := `DELETE FROM foods WHERE id=?`

	stmtx, err = tx.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmtx.Exec(foodId)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func insertFood(tx *sqlx.Tx, payload models.FoodDTO) (string, error) {
	query := `INSERT INTO foods (name,description,status,photos) VALUES (?,?,?,?) RETURNING id`
	stmt, err := tx.Prepare(query)
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

func insertFoodCategories(tx *sqlx.Tx, foodID, catID string) error {

	query := `INSERT INTO food_categories (food_id,category_id) values(:food_id,:category_id)`
	stmtx, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmtx.Exec(foodID, catID)
	if err != nil {
		return err
	}
	return nil
}
