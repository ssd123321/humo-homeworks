package main

func GetMoreExpensiveProduct(p1 Product, p2 Product) Product {
	if p1.Price > p2.Price {
		return p1
	}
	return p2
}
