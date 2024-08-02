package main

type Budget struct {
	amount float64
}

func (b *Budget) Allocate(amount float64) bool {
	if amount > b.amount {
		return false
	}
	b.amount -= amount
	return true
}
func (b Budget) Remainig() float64 {
	return b.amount
}
