package repositories

import "mehrangcode.ir/resturant/app/models"

type TableRepo interface {
	GetAll() ([]models.TableViewModel, error)
	GetById(id string) (models.TableViewModel, error)
	Create(payload models.TableDTO) (string, error)
	Update(id string, payload models.TableDTO) error
	ChangeStatus(id string, status uint) error
	Delete(id string) error
}
