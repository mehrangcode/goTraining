package storage

import (
	"mehrang.ir/taski/database"
	"mehrang.ir/taski/models"
	"mehrang.ir/taski/repositories"
)

type courseSqliteDB struct{}

func NewCourseSqliteDB() repositories.CourseRepo {
	return &courseSqliteDB{}
}

func (r *courseSqliteDB) Create(course *models.Course) error {
	return database.DB.Create(course).Error
}

func (r *courseSqliteDB) GetAll() ([]models.Course, error) {
	var courses []models.Course
	err := database.DB.Find(&courses).Error
	return courses, err
}

func (r *courseSqliteDB) GetByID(id uint) (models.Course, error) {
	var course models.Course
	err := database.DB.First(&course, id).Error
	return course, err
}

func (r *courseSqliteDB) Update(course *models.Course) error {
	return database.DB.Save(course).Error
}

func (r *courseSqliteDB) Delete(id uint) error {
	return database.DB.Delete(&models.Course{}, id).Error
}
