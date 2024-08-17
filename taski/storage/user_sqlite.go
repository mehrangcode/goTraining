package storage

import (
	"mehrang.ir/taski/database"
	"mehrang.ir/taski/models"
	"mehrang.ir/taski/repositories"
)

type userRepo struct{}

func NewUserSqliteDB() repositories.UserRepoInterface {
	return &userRepo{}
}

func (r *userRepo) GetAll() ([]models.User, error) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) GetById(id int) (models.User, error) {
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepo) Create(user *models.User) error {
	return database.DB.Create(user).Error
}

func (r *userRepo) Update(user *models.User) error {
	return database.DB.Save(user).Error
}

func (r *userRepo) Delete(id int) error {
	return database.DB.Delete(&models.User{}, id).Error
}
