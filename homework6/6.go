package main

import "fmt"

func calculate_deposit_compound(deposit_sum *int, interest_rate float64) {
	*deposit_sum += int((float64(*deposit_sum) * interest_rate))
	fmt.Printf("Баланс после начисления процентов: %d\n", *deposit_sum)
	if *deposit_sum > 300000 {
		fmt.Println("Достигнут лимит начислений")
	}
}
