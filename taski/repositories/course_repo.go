package repositories

import (
	"mehrang.ir/taski/models"
)

type CourseRepo interface {
	Create(course *models.Course) error
	GetAll() ([]models.Course, error)
	GetByID(id uint) (models.Course, error)
	Update(course *models.Course) error
	Delete(id uint) error
}
