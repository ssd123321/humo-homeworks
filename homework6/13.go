package main

import "fmt"

func check_balance(balance *int) {
	if *balance < 50000 {
		fmt.Println("Баланс низкий")
	} else {
		fmt.Printf("Баланс: %d\n", *balance)
	}
}
