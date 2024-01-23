package models

type UserViewModel struct {
	ID       uint   `json:"id" db:"users_id"`
	Name     string `json:"name" db:"users_name"`
	Email    string `json:"email" db:"users_email"`
	Password string `json:"-"`
}

type UserDTO struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
