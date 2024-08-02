package main

type Order1 struct {
	orderID  int
	product  string
	quantity int
}
type Score struct {
	Orders []Order1
}

func (o *Score) AddOrder(orderID int, product string, quantity int) {
	o.Orders = append(o.Orders, Order1{orderID, product, quantity})
}
func (o *Score) TotalQuantityByProduct(product string) int {
	total := 0
	for _, order := range o.Orders {
		if order.product == product {
			total += order.quantity
		}
	}
	return total
}
