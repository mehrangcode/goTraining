package models

import "gorm.io/gorm"

type Section struct {
	gorm.Model
	Title     string `json:"title"`
	Content   string `json:"content"`
	Video     string `json:"video"`
	LectureID uint
}
