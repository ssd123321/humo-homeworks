package main

import "fmt"

func calculate_investment_yearly(investment_sum *int, interest_rate float64) {
	*investment_sum += int((float64(*investment_sum) * interest_rate))
	fmt.Printf("Баланс после начисления процентов: %d\n", *investment_sum)
	if *investment_sum > 400000 {
		fmt.Println("Достигнут лимит дохода")
	}
}
