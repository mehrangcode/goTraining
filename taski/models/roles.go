package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string       `json:"name" gorm:"unique"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}
