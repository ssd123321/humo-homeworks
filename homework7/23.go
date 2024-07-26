package main

func ChangeProduct(p *Product, NewName string, NewPrice float64) {
	p.Name = NewName
	p.Price = NewPrice
}
