package main

import "fmt"

type Product struct {
	Name  string
	Price float64
}

func GetProduct(p Product) {
	fmt.Printf("Название продукта: %s, Цена продукта %.2f\n", p.Name, p.Price)
}
