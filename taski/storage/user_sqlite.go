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
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userSqliteDB) GetById(id int) (models.User, error) {
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
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
