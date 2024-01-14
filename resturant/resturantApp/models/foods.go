package models

type FoodViewModel struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Status      uint    `json:"status"`
	Photos      *string `json:"photos"`
}

type FoodDTO struct {
	ID          string
	Name        string
	Description *string
	Status      uint
	Photos      *string
}
