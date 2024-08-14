package main

import "strings"

func countWords(s string) int {
	words := strings.Split(s, " ")
	return len(words)
}
