package product

import (
	"errors"
	"micro/pattern/aggregate"

	"github.com/google/uuid"
)

var (
	ErrProductNotFound = errors.New("product not found")
)

type ProductRepository interface {
	GetAll() (p []aggregate.Product, err error)
	Get(id uuid.UUID) (p aggregate.Product, err error)
	Add(p aggregate.Product) error
	Update(p aggregate.Product) error
	Delete(id uuid.UUID) error
}
