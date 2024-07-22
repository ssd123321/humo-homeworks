package main

import "fmt"

func withdraw_money(balance *int, withdrawal_limit int, amount *int) {
	if *amount > withdrawal_limit {
		fmt.Printf("Превышен лимит снятия (%d рублей)\n", withdrawal_limit)
	} else if *amount > *balance {
		fmt.Println("Недостаточно средств")
	} else {
		*balance -= *amount
		fmt.Printf("Новый баланс после снятия: %d\n", *balance)
	}
}
