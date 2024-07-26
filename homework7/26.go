package main

import "fmt"

type Address struct {
	Street string
	City   string
}
type Person2 struct {
	Name string
	Address
}

func GetPersonsAddress(p Person2) {
	fmt.Printf("Имя: %s Адрес: %s %s\n", p.Name, p.City, p.Street)
}
