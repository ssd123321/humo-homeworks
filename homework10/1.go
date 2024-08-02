package main

import "fmt"

type Book struct {
	title  string
	author string
}
type Library struct {
	books []Book
}

func (b Book) GetDetails() string {
	return fmt.Sprintf("Название книги: %s, Автор: %s", b.title, b.author)
}
func (l *Library) AddBook(book Book) {
	l.books = append(l.books, book)
}
