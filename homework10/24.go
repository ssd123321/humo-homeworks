package main

type Transaction2 struct {
	transactionType string
	amount          float64
}
type Account3 struct {
	Transactions []Transaction2
}

func (t *Account3) AddTransaction(transactionType string, amount float64) {
	t.Transactions = append(t.Transactions, Transaction2{transactionType, amount})
}
func (t Account3) Balance() float64 {
	var sum float64
	for _, v := range t.Transactions {
		sum += v.amount
	}
	return sum
}
