package main

import (
	"fmt"
)

type BankAccount interface {
	Deposit(amount float64)
	Withdraw(amount float64) error
	GetBalance() float64
}

type Account struct {
	balance float64
}

func (a *Account) Deposit(amount float64) {
	a.balance += amount
}

func (a *Account) Withdraw(amount float64) error {
	if a.balance < amount {
		return fmt.Errorf("Insufficient funds. Current balance: %.2f", a.balance)
	}
	a.balance -= amount
	return nil
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func manageAccount(a BankAccount) {
	a.Deposit(1000.0)
	err := a.Withdraw(500.0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Current Balance: %.2f\n", a.GetBalance())

	err = a.Withdraw(600.0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Current Balance: %.2f\n", a.GetBalance())
}
