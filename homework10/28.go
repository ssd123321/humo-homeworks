package main

type Contract struct {
	contractID int
	client     string
	amount     float64
}
type ContractManager struct {
	Contracts []Contract
}

func (c *ContractManager) AddContract(contractID int, client string, amount float64) {
	c.Contracts = append(c.Contracts, Contract{contractID, client, amount})
}
func (c ContractManager) TotalAmountForClient(client string) float64 {
	var total float64
	for _, v := range c.Contracts {
		if v.client == client {
			total += v.amount
		}
	}
	return total
}
