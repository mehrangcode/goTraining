package models

type SectionFoodType struct {
	Price       uint    `json:"price"`
	Section_id  string  `json:"section_id"`
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Status      uint    `json:"status"`
	Photos      *string `json:"photos"`
}
type SectionViewModel struct {
	ID          string            `json:"id"`
	Menu_id     string            `json:"menu_id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Foods       []SectionFoodType `json:"foods"`
}

type SectionDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Foods       []struct {
		Price   uint   `json:"price"`
		Food_id string `json:"food_id"`
	} `json:"foods"`
}
type MenuViewModel struct {
	ID          string             `json:"id"`
	MenuID      string             `json:"menu_id,omitempty"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Status      uint               `json:"status"`
	Sections    []SectionViewModel `json:"sections"`
}

type MenuDTO struct {
	ID          string       `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Status      uint         `json:"status"`
	Sections    []SectionDTO `json:"sections"`
}
