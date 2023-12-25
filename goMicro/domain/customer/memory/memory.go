package memory

import (
	"micro/pattern/aggregate"
	"micro/pattern/domain/customer"
	"sync"

	"github.com/google/uuid"
)

type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (mr *MemoryRepository) Get(id uuid.UUID) (c aggregate.Customer, err error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}
	return c, customer.ErrCustomerIsNotFound
}

func (mr *MemoryRepository) Add(c aggregate.Customer) error {
	if mr.customers == nil {
		mr.Lock()
		mr.customers = make(map[uuid.UUID]aggregate.Customer)
		mr.Unlock()
	}
	if _, ok := mr.customers[c.GetID()]; ok {
		return customer.ErrFaildToAddCustomer
	}
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}

func (mr *MemoryRepository) Update(c aggregate.Customer) error {
	return nil
}

func (mr *MemoryRepository) Delete(id uuid.UUID) (err error) {
	return err
}
