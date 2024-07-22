package main

import "fmt"

func make_transfer(balance *int, transfer_limit int, amount int) {
	if amount > transfer_limit {
		fmt.Printf("Сумма перевода превышает лимит (%d рублей)\n", transfer_limit)
	} else {
		*balance -= amount
		fmt.Printf("Остаток на счете после перевода: %d\n", *balance)
	}
}
