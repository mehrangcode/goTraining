package repositories

import (
	"mehrang.ir/taski/models"
)

type CategoryRepo interface {
	Create(category *models.Category) error
	GetAll() ([]models.Category, error)
	GetByID(id uint) (models.Category, error)
	Update(category *models.Category) error
	Delete(id uint) error
}
