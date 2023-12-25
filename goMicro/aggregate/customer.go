package aggregate

import (
	"errors"
	"micro/pattern/entity"
	valueobject "micro/pattern/valueObject"

	"github.com/google/uuid"
)

var (
	ErrNameIsRequired = errors.New("name is required")
)

type Customer struct {
	person   *entity.Person
	products []*entity.Item

	transactions []*valueobject.Transaction
}

func NewCustomer(name string) (c Customer, err error) {
	if name == "" {
		return c, ErrNameIsRequired
	}
	person := &entity.Person{
		Id:   uuid.New(),
		Name: name,
		Age:  38,
	}

	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]*valueobject.Transaction, 0),
	}, nil
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Id = id
}

func (c *Customer) GetID() uuid.UUID {
	return c.person.Id
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Name = name
}
func (c *Customer) GetName() string {
	return c.person.Name
}
