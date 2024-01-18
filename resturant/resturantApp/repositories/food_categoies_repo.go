package repositories

import "mehrangcode.ir/resturant/app/models"

type FoodCategoriesRepo interface {
	GetAll() ([]models.FoodCategoryViewModel, error)
	GetById(id string) (models.FoodCategoryViewModel, error)
	Create(payload models.FoodCategoryDTO) (string, error)
	Update(id string, payload models.FoodCategoryDTO) error
	ChangeStatus(id string, status int) error
	Delete(id string) error
}
