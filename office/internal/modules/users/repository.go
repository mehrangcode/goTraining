package users

type Repository interface {
	GetAll() ([]ViewModel, error)
	GetByID(id string) (ViewModel, error)
	Create(payload ViewModel) error
	Update(payload ViewModel) error
	Delete(id string) error
}
