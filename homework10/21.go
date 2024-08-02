package main

type Order2 struct {
	id       int
	product  string
	quantity int
	price    float64
}
type Discount struct {
	percentage float64
}

func (o *Order2) ApplyDiscount(discount Discount) {
	discountAmount := o.price * (discount.percentage / 100)
	o.price = o.price - discountAmount
}

func (o Order2) TotalAmount() float64 {
	return o.price * float64(o.quantity)
}
