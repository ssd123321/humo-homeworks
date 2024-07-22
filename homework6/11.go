package main

import "fmt"

func add_balance(balance *int, amount *int) {
	*balance += *amount
	fmt.Printf("Новый баланс после пополнения: %d\n", *balance)
}
