package main

import "fmt"

func calculate_interest(balance *int, interest_rate float64) {
	interest := float64(*balance) * interest_rate
	*balance += int(interest)
	fmt.Printf("Баланс после начисления процентов: %d\n", *balance)
}
