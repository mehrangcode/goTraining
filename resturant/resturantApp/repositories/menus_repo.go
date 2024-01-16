package repositories

import "mehrangcode.ir/resturant/app/models"

type MenuRepo interface {
	GetAll() ([]models.MenuViewModel, error)
	// GetById(id string) (models.MenuViewModel, error)
	Create(payload models.MenuDTO) (string, error)
	// Update(userId string, payload models.MenuDTO) error
	// Delete(userId string) error
}
