package main

import "fmt"

func calculate_payment(credit_sum *int, interest_rate float64) {
	payment := float64(*credit_sum) * (interest_rate / 12)
	*credit_sum -= int(payment)
	fmt.Printf("Остаток по кредиту после выплаты: %d\n", *credit_sum)
	if *credit_sum < 500000 {
		fmt.Println("Почти погашен кредит")
	}
}
