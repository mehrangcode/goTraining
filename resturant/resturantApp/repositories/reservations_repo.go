package repositories

import "mehrangcode.ir/resturant/app/models"

type ReservationRepo interface {
	GetAll() ([]models.ReservationViewModel, error)
	GetById(id string) (models.ReservationViewModel, error)
	Create(payload models.ReservationDTO) (string, error)
	Update(id string, payload models.ReservationDTO) error
	ChangeStatus(id string, status uint) error
	Delete(id string) error
}
