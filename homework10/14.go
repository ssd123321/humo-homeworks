package main

type Account struct {
	balance float64
}

func (a *Account) Deposit(amount float64) {
	a.balance += amount
}
func (a *Account) Withdraw(amount float64) bool {
	if amount > a.balance {
		return false
	}
	a.balance -= amount
	return true
}
func (a Account) GetBalance() float64 {
	return a.balance
}
