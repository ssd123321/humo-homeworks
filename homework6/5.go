package main

import "fmt"

func update_rate(usd_rate *float64) {
	*usd_rate += 0.5
	fmt.Printf("Новый курс доллара к рублю: %.2f\n", *usd_rate)
}

func convert_currency(usd_rate float64, amount float64) {
	rub_amount := amount * usd_rate
	fmt.Printf("%.2f долларов равны %.2f рублей\n", amount, rub_amount)
	if usd_rate < 70 {
		fmt.Println("Курс слишком низкий")
	}
}
