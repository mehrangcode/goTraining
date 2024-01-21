package models

type TableViewModel struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Capacity uint   `json:"Capacity"`
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

type ReservationDTO struct {
	TableID string `json:"table_id"`
	UserID  string `json:"user_id"`
	Date    string `json:"date"`
}
