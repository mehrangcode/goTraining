package models

import "gorm.io/gorm"

type Lecture struct {
	gorm.Model
	Title    string    `json:"title"`
	Lead     string    `json:"lead"`
	Sections []Section `json:"sections" gorm:"foreignKey:LectureID"`
	CourseID uint
}
