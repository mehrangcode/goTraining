package services

import (
	"fmt"
	"micro/pattern/aggregate"
	"micro/pattern/domain/customer"
	"micro/pattern/domain/customer/memory"
	"micro/pattern/domain/product"
	prodMem "micro/pattern/domain/product/memory"

	"github.com/google/uuid"
)

type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
}

type OrderConfiguration func(os *OrderService) error

func NewOrderService(cfgs ...OrderConfiguration) (os *OrderService, err error) {
	os = &OrderService{}
	for _, cfg := range cfgs {
		err = cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prodMem.New()
		for _, p := range products {
			if err := pr.Add(p); err != nil {
				return err
			}
		}
		os.products = pr
		return nil
	}
}
func (os *OrderService) CreateOrder(customerId uuid.UUID, productIds []uuid.UUID) error {
	fmt.Println("CREATE ORDER _______________________", customerId)
	cc, _ := os.customers.GetAll()
	fmt.Println("LEN: ", len(cc))
	for _, v := range cc {
		fmt.Println("ID", v.GetID())
	}
	customer, err := os.customers.Get(customerId)
	if err != nil {
		return err
	}
	var products []aggregate.Product
	var fee float64
	for _, pId := range productIds {
		p, err := os.products.Get(pId)
		if err != nil {
			return err
		}
		products = append(products, p)
		fee += p.GetPrice()
	}
	fmt.Println("ORDER: ", customer.GetName(), len(products), fee)
	return nil
}
