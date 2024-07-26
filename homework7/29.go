package main

import "fmt"

type Event struct {
	Title    string
	Location Address
}

func GetEvent(e Event) {
	fmt.Printf("Название ивента: %s, Адрес: %s %s\n", e.Title, e.Location.City, e.Location.Street)
}
