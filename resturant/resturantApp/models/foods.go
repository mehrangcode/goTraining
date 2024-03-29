package models

type FoodViewModel struct {
	ID          string                  `json:"id"`
	Name        string                  `json:"name"`
	Description *string                 `json:"description"`
	Status      uint                    `json:"status"`
	Photos      *string                 `json:"photos"`
	Categories  []FoodCategoryViewModel `json:"categories"`
}

type FoodDTO struct {
	ID          string
	Name        string
	Description *string
	Status      uint
	Photos      *string
	Categories  []string
}
