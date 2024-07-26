package main

type Item struct {
	Name     string
	Quantity int
}

func GetQuantities(i ...Item) int {
	var sum int
	for _, Item := range i {
		sum += Item.Quantity
	}
	return sum
}
