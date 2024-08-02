package main

type Order3 struct {
	id       int
	product  string
	quantity int
	price    float64
}
type Order3Manager struct {
	Orders []Order3
}

func (o *Order3Manager) AddOrder(id int, product string, quantity int, price float64) {
	o.Orders = append(o.Orders, Order3{id, product, quantity, price})
}
func (o *Order3Manager) ReturnOrder(id int, quantity int) {
	for _, v := range o.Orders {
		if v.id == id {
			o.Orders = append(o.Orders[:id], o.Orders[id+1:]...)
		}
	}
}
func (o *Order3Manager) TotalActiveOrders() float64 {
	var total float64
	for _, v := range o.Orders {
		total += v.price
	}
	return total
}
