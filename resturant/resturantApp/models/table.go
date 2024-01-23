package models

type TableViewModel struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Capacity uint   `json:"capacity"`
	Photos   string `json:"photos"`
	Status   uint   `json:"status"`
}
type TableDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Capacity uint   `json:"Capacity"`
	Photos   string `json:"photos"`
	Status   uint   `json:"status"`
}
