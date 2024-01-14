package models

type SectionViewModel struct {
	ID          string
	Title       string
	Description string
	Foods       []struct {
		Price int
		*FoodViewModel
	}
}

type SectionDTO struct {
	ID          string
	Title       string
	Description string
	Foods       []struct {
		Price  int
		FoodID string
	}
}
type MenuViewModel struct {
	ID          string
	Title       string
	Description string
	Status      uint
	Sections    []SectionViewModel
}

type MenuDTO struct {
	ID          string
	Title       string
	Description string
	Status      uint
	Sections    []SectionDTO
}
