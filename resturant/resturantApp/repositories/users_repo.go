package repositories

import "mehrangcode.ir/resturant/app/models"

type UserRepo interface {
	GetAll() ([]models.UserViewModel, error)
	GetById(id string) (models.UserViewModel, error)
	Create(payload models.UserDTO) (string, error)
	Update(userId string, payload models.UserDTO) error
	Delete(userId string) error
}
