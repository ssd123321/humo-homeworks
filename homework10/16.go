package main

type ShoppingItem struct {
	name  string
	price float64
}
type ShoppingList struct {
	ShoppingItems []ShoppingItem
}

func (s *ShoppingList) AddItem(name string, price float64) {
	s.ShoppingItems = append(s.ShoppingItems, ShoppingItem{name, price})
}
func (s ShoppingList) TotalPrice() float64 {
	var total float64
	for _, item := range s.ShoppingItems {
		total += item.price
	}
	return total
}
