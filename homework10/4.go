package main

type Item struct {
	name  string
	price float64
}
type Inventory struct {
	Items []Item
}

func (i *Inventory) AddItem(item Item) {
	i.Items = append(i.Items, item)
}

func (i Inventory) TotalValue() float64 {
	var total float64
	for _, item := range i.Items {
		total += item.price
	}
	return total
}
