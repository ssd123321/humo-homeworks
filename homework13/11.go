package main

import (
	"fmt"
)

type Budget interface {
	AddIncome(amount float64)
	AddExpense(amount float64)
	GetBalance() float64
}

type MonthlyBudget struct {
	income  float64
	expense float64
}

func (b *MonthlyBudget) AddIncome(amount float64) {
	b.income += amount
}

func (b *MonthlyBudget) AddExpense(amount float64) {
	b.expense += amount
}

func (b *MonthlyBudget) GetBalance() float64 {
	return b.income - b.expense
}

func manageMonthlyBudget(b Budget) {
	b.AddIncome(5000.0)
	b.AddExpense(3000.0)
	b.AddExpense(1500.0)
	fmt.Printf("Current Balance: %.2f\n", b.GetBalance())
}
