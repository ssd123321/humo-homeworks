package main

import (
    "fmt"
)

type OrderManager interface {
    AddOrder(orderID string, amount float64)
    GetTotalSales() float64
}

type OnlineStore struct {
    orders map[string]float64
    totalSales float64
}
func (s *OnlineStore) AddOrder(orderID string, amount float64) {
    s.orders[orderID] = amount
    s.totalSales += amount
}

func (s *OnlineStore) GetTotalSales() float64 {
    return s.totalSales
}

func manageOrders(om OrderManager) {
    om.AddOrder("order1", 100.0)
    om.AddOrder("order2", 200.0)
    om.AddOrder("order3", 150.0)

    totalSales := om.GetTotalSales()
    fmt.Printf("Total Sales: %.2f\n", totalSales)
}


