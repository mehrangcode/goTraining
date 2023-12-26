package memory

import (
	"micro/pattern/aggregate"
	"micro/pattern/domain/product"
	"sync"

	"github.com/google/uuid"
)

type ProductMemoryRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func New() *ProductMemoryRepository {
	return &ProductMemoryRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

func (pm *ProductMemoryRepository) GetAll() (ps []aggregate.Product, err error) {
	for _, p := range pm.products {
		ps = append(ps, p)
	}
	return ps, nil
}

func (pm *ProductMemoryRepository) Get(id uuid.UUID) (p aggregate.Product, err error) {
	if p, ok := pm.products[id]; ok {
		return p, nil
	}
	return p, product.ErrProductNotFound
}

func (pm *ProductMemoryRepository) Add(p aggregate.Product) error {
	pm.Lock()
	defer pm.Unlock()
	if pm.products == nil {
		pm.products = make(map[uuid.UUID]aggregate.Product)
	}
	if _, ok := pm.products[p.GetID()]; ok {
		return aggregate.ErrProductIsSxists
	}
	pm.products[p.GetID()] = p
	return nil
}

func (pm *ProductMemoryRepository) Update(p aggregate.Product) error {
	pm.Lock()
	defer pm.Unlock()
	if _, ok := pm.products[p.GetID()]; !ok {
		return product.ErrProductNotFound
	}
	pm.products[p.GetID()] = p
	return nil
}

func (pm *ProductMemoryRepository) Delete(id uuid.UUID) error {
	pm.Lock()
	defer pm.Unlock()
	if _, ok := pm.products[id]; !ok {
		return product.ErrProductNotFound
	}
	delete(pm.products, id)
	return nil
}
