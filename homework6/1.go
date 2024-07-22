package main

import "fmt"

func add_monthly_balance(balance *int, monthly_addition int) {
    *balance += monthly_addition
    fmt.Printf("Баланс после пополнения: %d\n", *balance)
    if *balance > 500000 {
        fmt.Println("Достигнут лимит накоплений")
    }
}
