package storage

import (
	"mehrang.ir/taski/database"
	"mehrang.ir/taski/models"
	"mehrang.ir/taski/repositories"
)

type permissionSqliteDB struct{}

func NewPermissionSqliteDB() repositories.PermissionRepo {
	return &permissionSqliteDB{}
}

func (r *permissionSqliteDB) Create(permission *models.Permission) error {
	return database.DB.Create(permission).Error
}

func (r *permissionSqliteDB) GetAll() ([]models.Permission, error) {
	var permissions []models.Permission
	err := database.DB.Find(&permissions).Error
	return permissions, err
}

func (r *permissionSqliteDB) GetByID(id uint) (models.Permission, error) {
	var permission models.Permission
	err := database.DB.First(&permission, id).Error
	return permission, err
}

func (r *permissionSqliteDB) Update(permission *models.Permission) error {
	return database.DB.Save(permission).Error
}

func (r *permissionSqliteDB) Delete(id uint) error {
	return database.DB.Delete(&models.Permission{}, id).Error
}
