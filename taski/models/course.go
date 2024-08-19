package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Title      string     `json:"title"`
	Teacher    string     `json:"teacher"`
	Level      uint       `json:"level"`
	Tags       string     `json:"tags"`
	Categories []Category `json:"categories" gorm:"many2many:course_categories;"`
	Lectures   []Lecture  `json:"lectures" gorm:"foreignKey:CourseID"`
}
