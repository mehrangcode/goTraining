package apptests

import (
	"fmt"
	"micro/pattern/aggregate"
	"micro/pattern/domain/customer/memory"
	prodMem "micro/pattern/domain/product/memory"
	"micro/pattern/services"

	"github.com/google/uuid"
)

func OrderMock() error {
	cRepo := memory.New()
	pRepo := prodMem.New()
	c, err := aggregate.NewCustomer("Mehran")
	if err != nil {
		return err
	}
	err = cRepo.Add(c)
	if err != nil {
		return err
	}
	p, err := aggregate.NewProduct("Laptop", "Asus brand Laptop", 1200)
	if err != nil {
		return err
	}
	err = pRepo.Add(p)
	if err != nil {
		return err
	}
	car, err := aggregate.NewProduct("Car", "Nissan brand Laptop", 3000)
	if err != nil {
		return err
	}
	err = pRepo.Add(car)
	if err != nil {
		return err
	}
	list, err := pRepo.GetAll()
	if err != nil {
		return err
	}
	var products []aggregate.Product
	products = append(products, list...)
	var productsId []uuid.UUID
	for _, prod := range products {
		productsId = append(productsId, prod.GetID())
	}
	order, err := services.NewOrderService(
		services.WithCustomerRepository(cRepo),
		services.WithMemoryProductRepository(products),
	)
	if err != nil {
		return err
	}
	err = order.CreateOrder(c.GetID(), productsId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("items: %v p: %v", c.GetName(), len(products))
	return nil
}
