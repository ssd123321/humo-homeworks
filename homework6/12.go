package main

import "fmt"

func withdraw_balance(balance *int, amount *int) {
	if *amount > *balance {
		fmt.Println("Недостаточно средств")
	} else {
		*balance -= *amount
		fmt.Printf("Новый баланс после снятия: %d\n", *balance)
	}
}
