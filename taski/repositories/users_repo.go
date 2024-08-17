package repositories

import (
	"mehrang.ir/taski/models"
)

type UserRepoInterface interface {
	GetAll() ([]models.User, error)
	GetById(id int) (models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id int) error
}
