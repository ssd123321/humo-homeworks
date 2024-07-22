package main

import "fmt"

func add_expense(daily_limit *int, expense *int) {
	if *expense > *daily_limit {
		fmt.Printf("Сумма перевода превышает лимит (%d рублей)\n", *daily_limit)
	} else {
		*daily_limit -= *expense
		fmt.Printf("Текущая сумма расходов: %d\n", *daily_limit)
	}
}
