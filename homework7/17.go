package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func PrintPerson(p Person) {
	fmt.Printf("%s %s %d лет\n", p.FirstName, p.LastName, p.Age)
}
