package main

import "fmt"

type Book struct {
	Title  string
	Author string
}
type Library struct {
	Name  string
	Books []Book
}

func CheckLibrary(l Library) {
	fmt.Printf("Название библиотеки %s\n", l.Name)
	for number, book := range l.Books {
		fmt.Printf("Книга номер %d - Название: %s, Автор: %s\n", number+1, book.Title, book.Author)
	}
}
