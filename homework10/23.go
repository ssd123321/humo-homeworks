package main

type Sale struct {
	item   string
	amount float64
}
type SalesReport struct {
	Sales []Sale
}

func (s *SalesReport) AddSale(item string, amount float64) {
	s.Sales = append(s.Sales, Sale{item, amount})
}
func (s SalesReport) TotalSales() float64 {
	var total float64
	for _, sale := range s.Sales {
		total += sale.amount
	}
	return total
}
