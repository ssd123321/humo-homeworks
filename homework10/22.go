package main

type Account2 struct {
	balance float64
	history []Transaction
}

func (a *Account2) Deposit(amount float64) {
	a.balance += amount
	a.history = append(a.history, Transaction{amount, "Deposit transcation"})
}
func (a *Account2) Withdraw(amount float64) bool {
	if amount > a.balance {
		return false
	}
	a.balance -= amount
	a.history = append(a.history, Transaction{amount, "Withdraw transaction"})
	return true

}
func (a Account2) TransactionHistory() []string {
	data := make([]string, 0)
	for _, v := range a.history {
		data = append(data, v.description)
	}
	return data
}
