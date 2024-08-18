package storage

import (
	"mehrang.ir/taski/database"
	"mehrang.ir/taski/models"
	"mehrang.ir/taski/repositories"
)

type roleSqliteDB struct{}

func NewRoleSqliteDB() repositories.RoleRepo {
	return &roleSqliteDB{}
}

func (r *roleSqliteDB) Create(role *models.Role) error {
	return database.DB.Create(role).Error
}

func (r *roleSqliteDB) GetAll() ([]models.Role, error) {
	var roles []models.Role
	err := database.DB.Preload("Permissions").Find(&roles).Error
	return roles, err
}

func (r *roleSqliteDB) GetByID(id uint) (models.Role, error) {
	var role models.Role
	err := database.DB.Preload("Permissions").First(&role, id).Error
	return role, err
}

func (r *roleSqliteDB) Update(role *models.Role) error {
	return database.DB.Save(role).Error
}

func (r *roleSqliteDB) Delete(id uint) error {
	return database.DB.Delete(&models.Role{}, id).Error
}

func (r *roleSqliteDB) AddPermissionsToRole(roleID int, permissionIDs []int) (models.Role, error) {

	var role models.Role
	if err := database.DB.First(&role, roleID).Error; err != nil {
		return role, err
	}

	var permissions []models.Permission
	if err := database.DB.Where("id IN ?", permissionIDs).Find(&permissions).Error; err != nil {
		return role, err
	}

	if err := database.DB.Model(&role).Association("Permissions").Append(&permissions); err != nil {
		return role, err
	}

	return role, nil
}
