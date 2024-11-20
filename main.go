package main

import (
	"github.com/sunilsoni88/lowleveldesign/observer"
)

func main() {
	item1 := observer.NewItem("item1")
	item2 := observer.NewItem("item2")

	customer1 := observer.NewCustomer("c1")
	customer2 := observer.NewCustomer("c2")

	item1.RegisterObserver(customer1)
	item2.RegisterObserver(customer2)

	item1.UpdateAvailability()
	item2.UpdateAvailability()

}
