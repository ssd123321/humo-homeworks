package main

type Product struct {
	name     string
	quantity int
}
type Warehouse struct {
	Products []Product
}

func (w *Warehouse) AddProduct(name string, quantity int) {
	w.Products = append(w.Products, Product{name, quantity})
}
func (w Warehouse) TotalQuantity() int {
	total := 0
	for _, product := range w.Products {
		total += product.quantity
	}
	return total
}
