package models

type ReservationViewModel struct {
	ID     uint          `json:"id"`
	User   UserViewModel `json:"user"`
	Guests uint          `json:"guests"`
	Date   string        `json:"date"`
}

type ReservationDTO struct {
	ID     uint   `json:"id"`
	UserID string `json:"user_id"`
	Guests uint   `json:"guests"`
	Date   string `json:"date"`
}
