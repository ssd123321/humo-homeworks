package main

type Book1 struct {
	title  string
	author string
	year   int
}
type Library1 struct {
	Books []Book1
}

func (l *Library1) AddBook(title, author string, year int) {
	l.Books = append(l.Books, Book1{title, author, year})
}
func (l Library1) BooksAfterYear(year int) []Book1 {
	x := make([]Book1, 0)
	for _, v := range l.Books {
		if v.year > year {
			x = append(x, v)
		}
	}
	return x
}
