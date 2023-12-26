package aggregate

import (
	"errors"
	"micro/pattern/entity"

	"github.com/google/uuid"
)

var (
	ErrMissingValue    = errors.New("name and description is required")
	ErrProductIsSxists = errors.New("product is already exists")
)

type Product struct {
	item     *entity.Item
	price    float64
	quantity int
}

func NewProduct(name string, description string, price float64) (p Product, err error) {
	if name == "" || description == "" {
		return p, ErrMissingValue
	}
	p = Product{
		item: &entity.Item{
			Id:          uuid.New(),
			Name:        name,
			Description: description,
		},
		price:    price,
		quantity: 0,
	}
	return p, nil
}

func (p *Product) GetID() uuid.UUID {
	return p.item.Id
}

func (p *Product) GetItem() *entity.Item {
	return p.item
}

func (p *Product) GetPrice() float64 {
	return p.price
}
