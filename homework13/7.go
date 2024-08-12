package main

import (
	"fmt"
	"strings"
)

type StringReverser interface {
	Reverse() string
	ReverseWords() string
}

type TextReverser struct {
	text string
}

func (r *TextReverser) Reverse() string {
	runes := []rune(r.text)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func (r *TextReverser) ReverseWords() string {
	words := strings.Split(r.text, " ")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func reverseText(r StringReverser) {
	fmt.Println("Reversed:", r.Reverse())
	fmt.Println("Reversed Words:", r.ReverseWords())
}
