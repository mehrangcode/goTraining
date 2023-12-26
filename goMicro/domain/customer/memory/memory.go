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

func (mr *MemoryRepository) GetAll() (cs []aggregate.Customer, err error) {
	for _, v := range mr.customers {
		cs = append(cs, v)
	}
	return cs, nil
}
func (mr *MemoryRepository) Get(id uuid.UUID) (c aggregate.Customer, err error) {
	if c, ok := mr.customers[id]; ok {
		return c, nil
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
	if _, ok := mr.customers[c.GetID()]; !ok {
		return customer.ErrCustomerIsNotFound
	}
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}

func (mr *MemoryRepository) Delete(id uuid.UUID) (err error) {
	if _, ok := mr.customers[id]; !ok {
		return customer.ErrCustomerIsNotFound
	}
	mr.Lock()
	delete(mr.customers, id)
	mr.Unlock()
	return nil
}
