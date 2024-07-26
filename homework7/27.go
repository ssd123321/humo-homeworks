package main

import "fmt"

type Company struct {
	Name     string
	Location Address
}

func GetCompanyAddress(c Company) {
	fmt.Printf("Название компании: %s, Адрес: %s %s\n", c.Name, c.Location.City, c.Location.Street)
}
