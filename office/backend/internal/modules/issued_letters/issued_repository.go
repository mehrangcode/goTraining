package issued_letters

import "mehrangcode.ir/office/internal/types"

type Repository interface {
	GetAll() ([]types.IssuedLetterViewModel, error)
	Create(payload types.IssuedLetterDTO) (string, error)
	Update(itemId string, payload types.IssuedLetterDTO) error
	Delete(itemId string) error
}
