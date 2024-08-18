package repositories

import (
	"mehrang.ir/taski/models"
)

type PermissionRepo interface {
	Create(permission *models.Permission) error
	GetAll() ([]models.Permission, error)
	GetByID(id uint) (models.Permission, error)
	Update(permission *models.Permission) error
	Delete(id uint) error
}
