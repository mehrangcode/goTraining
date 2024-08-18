// models/user.go
package models

import (
	"gorm.io/gorm"
)

type UserVieModel struct {
	ID       uint
	Phone    string
	FullName string
	Roles    []Role `json:"rules"`
}
type User struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone" gorm:"unique"`
	Password  string `json:"password"`
	FullName  string `json:"fullName" gorm:"-"`
	Roles     []Role `gorm:"many2many:user_roles;"`
}
