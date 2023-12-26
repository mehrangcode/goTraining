package customer

import (
	"errors"
	"micro/pattern/aggregate"

	"github.com/google/uuid"
)

var (
	ErrCustomerIsNotFound = errors.New("customer is not found")
	ErrFaildToAddCustomer = errors.New("faild to add customer")
)

type CustomerRepository interface {
	GetAll() ([]aggregate.Customer, error)
	Get(id uuid.UUID) (aggregate.Customer, error)
	Add(c aggregate.Customer) error
	Update(c aggregate.Customer) error
	Delete(customerId uuid.UUID) error
}
