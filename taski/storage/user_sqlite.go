package storage

import (
	"mehrang.ir/taski/database"
	"mehrang.ir/taski/models"
	"mehrang.ir/taski/repositories"
)

type userSqliteDB struct{}

func NewUserSqliteDB() repositories.UserRepo {
	return &userSqliteDB{}
}

func (r *userSqliteDB) GetAll() ([]models.User, error) {
	var users []models.User
	if err := database.DB.Preload("Roles.Permissions").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userSqliteDB) GetById(id int) (models.User, error) {
	var user models.User
	if err := database.DB.Preload("Roles.Permissions").First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userSqliteDB) Create(user *models.User) error {
	return database.DB.Create(user).Error
}

func (r *userSqliteDB) Update(user *models.User) error {
	return database.DB.Save(user).Error
}

func (r *userSqliteDB) Delete(id int) error {
	return database.DB.Delete(&models.User{}, id).Error
}
func (r *userSqliteDB) GetByPhone(phone string) (models.User, error) {
	var user models.User
	if err := database.DB.Where("phone = ?", phone).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userSqliteDB) AddRolesToUser(userID uint, roleIDs []uint) (models.User, error) {
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return user, err
	}

	var roles []models.Role
	if err := database.DB.Where("id IN ?", roleIDs).Find(&roles).Error; err != nil {
		return user, err
	}
	if err := database.DB.Model(&user).Association("Roles").Append(&roles); err != nil {
		return user, err
	}
	if err := database.DB.Preload("Roles.Permissions").First(&user, userID).Error; err != nil {
		return user, err
	}

	return user, nil
}
