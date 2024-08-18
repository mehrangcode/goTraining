package repositories

import (
	"mehrang.ir/taski/models"
)

type RoleRepo interface {
	Create(role *models.Role) error
	GetAll() ([]models.Role, error)
	GetByID(id uint) (models.Role, error)
	Update(role *models.Role) error
	Delete(id uint) error
	AddPermissionsToRole(roleID int, permissionIDs []int) (models.Role, error)
}
