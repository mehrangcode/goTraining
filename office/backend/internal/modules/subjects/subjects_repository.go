package subjects

import "mehrangcode.ir/office/internal/types"

type Repository interface {
	GetAll() ([]types.SubjectViewModel, error)
	Create(payload types.SubjectDTO) (string, error)
	Update(subjectId string, payload types.SubjectDTO) error
	Delete(subjectId string) error
}
