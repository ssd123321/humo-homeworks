package main

import (
    "fmt"
)

type Library interface {
    AddBook(title, author string)
    GetBooks() []string
}

type BookLibrary struct {
    books []book
}

type book struct {
    title  string
    author string
}
func (l *BookLibrary) AddBook(title, author string) {
    l.books = append(l.books, book{
        title:  title,
        author: author,
    })
}

func (l *BookLibrary) GetBooks() []string {
    var bookTitles []string
    for _, b := range l.books {
        bookTitles = append(bookTitles, b.title)
    }
    return bookTitles
}

func manageLibrary(l Library) {
    l.AddBook("To Kill a Mockingbird", "Harper Lee")
    l.AddBook("1984", "George Orwell")
    l.AddBook("The Great Gatsby", "F. Scott Fitzgerald")

    books := l.GetBooks()
    fmt.Println("Books in the library:")
    for _, title := range books {
        fmt.Println(title)
    }
}


