package main

import (
	"fmt"
	"strings"
)

type StringProcessor interface {
	Length() int
	WordCount() int
}

type MyString struct {
	text string
}

func (s *MyString) Length() int {
	return len(s.text)
}

func (s *MyString) WordCount() int {
	return len(strings.Fields(s.text))
}

func processString(sp StringProcessor) {
	fmt.Println("Length:", sp.Length())
	fmt.Println("Word Count:", sp.WordCount())
}
