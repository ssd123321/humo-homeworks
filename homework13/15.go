package main

import (
	"fmt"
)

type Invoice interface {
	AddItem(name string, price float64)
	Total() float64
}

type SimpleInvoice struct {
	items []item
}

type item struct {
	name  string
	price float64
}

func (i *SimpleInvoice) AddItem(name string, price float64) {
	i.items = append(i.items, item{
		name:  name,
		price: price,
	})
}

func (i *SimpleInvoice) Total() float64 {
	var total float64
	for _, item := range i.items {
		total += item.price
	}
	return total
}

func generateInvoice(i Invoice) {
	i.AddItem("Item 1", 10.0)
	i.AddItem("Item 2", 15.0)
	i.AddItem("Item 3", 20.0)
	fmt.Printf("Total: %.2f\n", i.Total())
}
