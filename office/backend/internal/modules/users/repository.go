package users

import "mehrangcode.ir/office/internal/types"

type UserRepository interface {
	GetAll() ([]types.UserViewModel, error)
	Create(payload types.UserDTO) (string, error)
	Update(userId string, payload types.UserDTO) error
	Delete(userId string) error
}
