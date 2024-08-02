package main

type Order struct {
	id          int
	totalAmount float64
}
type OrderManager struct {
	Orders []Order
}

func (o *OrderManager) AddOrder(id int, totalAmount float64) {
	o.Orders = append(o.Orders, Order{id, totalAmount})
}
func (o OrderManager) TotalOrdersAmount() float64 {
	var total float64
	for _, order := range o.Orders {
		total += order.totalAmount
	}
	return total
}
