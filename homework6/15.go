package main

import "fmt"

func transfer_funds(account1_balance *int, account2_balance *int, amount *int) {
	if *amount > *account1_balance {
		fmt.Println("Недостаточно средств на первом счете")
	} else {
		*account1_balance -= *amount
		*account2_balance += *amount
		fmt.Printf("Баланс первого счета: %d\n", *account1_balance)
		fmt.Printf("Баланс второго счета: %d\n", *account2_balance)
	}
}
