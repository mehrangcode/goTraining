package storage

import (
	"mehrang.ir/taski/database"
	"mehrang.ir/taski/models"
	"mehrang.ir/taski/repositories"
)

type categorySqliteDB struct{}

func NewCategorySqliteDB() repositories.CategoryRepo {
	return &categorySqliteDB{}
}

func (r *categorySqliteDB) Create(category *models.Category) error {
	return database.DB.Create(category).Error
}

func (r *categorySqliteDB) GetAll() ([]models.Category, error) {
	var categories []models.Category
	err := database.DB.Find(&categories).Error
	return categories, err
}

func (r *categorySqliteDB) GetByID(id uint) (models.Category, error) {
	var category models.Category
	err := database.DB.First(&category, id).Error
	return category, err
}

func (r *categorySqliteDB) Update(category *models.Category) error {
	return database.DB.Save(category).Error
}

func (r *categorySqliteDB) Delete(id uint) error {
	return database.DB.Delete(&models.Category{}, id).Error
}
