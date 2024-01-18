package models

type FoodCategoryViewModel struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Status      uint    `json:"status"`
	Avatar      *string `json:"avatar"`
}

type FoodCategoryDTO struct {
	ID          string
	Title       string
	Description *string
	Status      uint
	Avatar      *string
}
