package main

import "fmt"

func GetDiscount() float64 {
	var amount float64
	fmt.Scan(&amount)
	if amount > 1000{
		return (amount / 100) * 10
	}else if amount <= 1000 && amount > 500{
		return (amount / 100) * 5
	}
	return amount
}
