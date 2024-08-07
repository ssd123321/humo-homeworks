package main

import "strings"

func ReverseWords(s string) string {
	words := strings.Split(s, " ")
	NewSequence := make([]string, 0)
	for i := len(words) - 1; i >= 0; i-- {
		NewSequence = append(NewSequence, words[i])
	}
	return strings.Join(NewSequence, " ")
}
