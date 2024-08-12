package main

import (
	"fmt"
	"strings"
)

type DuplicateRemover interface {
	RemoveDuplicates() string
	RemoveDuplicatesCaseInsensitive() string
}

type TextDuplicateRemover struct {
	text string
}

func (r *TextDuplicateRemover) RemoveDuplicates() string {
	words := strings.Split(r.text, " ")
	uniqueWords := make(map[string]bool)
	for _, word := range words {
		uniqueWords[word] = true
	}
	uniqueWordSlice := make([]string, 0, len(uniqueWords))
	for word := range uniqueWords {
		uniqueWordSlice = append(uniqueWordSlice, word)
	}
	return strings.Join(uniqueWordSlice, " ")
}

func (r *TextDuplicateRemover) RemoveDuplicatesCaseInsensitive() string {
	words := strings.Split(r.text, " ")
	uniqueWords := make(map[string]bool)
	for _, word := range words {
		uniqueWords[strings.ToLower(word)] = true
	}
	uniqueWordSlice := make([]string, 0, len(uniqueWords))
	for word := range uniqueWords {
		uniqueWordSlice = append(uniqueWordSlice, word)
	}
	return strings.Join(uniqueWordSlice, " ")
}

func removeDuplicates(r DuplicateRemover) {
	fmt.Println("Removed Duplicates:", r.RemoveDuplicates())
	fmt.Println("Removed Duplicates (Case Insensitive):", r.RemoveDuplicatesCaseInsensitive())
}
