package models

type FoodViewModel struct {
	ID          string
	Name        string
	Description *string
	Status      uint
	Photos      *string
}

type FoodDTO struct {
	ID          string
	Name        string
	Description *string
	Status      uint
	Photos      *string
}
