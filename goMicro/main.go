package main

import (
	"fmt"
	"micro/pattern/aggregate"
)

func main() {
	fmt.Println("App Is Start")
	customer, err := aggregate.NewCustomer("Mehran")
	if err != nil {
		fmt.Printf("\nERR: %v", err)
		return
	}
	fmt.Printf("\ncustomer is : %v", customer.GetID())
	_, err = aggregate.NewCustomer("")
	if err != nil {
		fmt.Printf("\nERR: %v", err)
		return
	}
}
