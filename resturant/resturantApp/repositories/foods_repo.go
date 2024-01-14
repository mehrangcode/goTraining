package repositories

import "mehrangcode.ir/resturant/app/models"

type FoodRepo interface {
	GetAll() ([]models.FoodViewModel, error)
	GetById(id string) (models.FoodViewModel, error)
	Create(payload models.FoodDTO) (string, error)
	Update(id string, payload models.FoodDTO) error
	ChangeStatus(id string, status int) error
	Delete(id string) error
}
