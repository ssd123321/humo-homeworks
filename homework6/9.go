package main

import "fmt"

func add_transaction(balance *int, commission_rate float64, amount *int) {
	*amount += int(float64(*amount) * commission_rate)
	*balance += *amount
	fmt.Printf("Баланс после добавления транзакции: %d\n", *balance)
}
