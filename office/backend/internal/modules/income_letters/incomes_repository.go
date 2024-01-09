package income_letters

import "mehrangcode.ir/office/internal/types"

type Repository interface {
	GetAll() ([]types.IncomeLetterViewModel, error)
	Create(payload types.IncomeLetterDTO) (string, error)
	Update(itemId string, payload types.IncomeLetterDTO) error
	Delete(itemId string) error
}
