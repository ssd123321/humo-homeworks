package main

import "strings"

func countUniqueWords(s string) int {
	words := strings.Fields(s)
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[strings.ToLower(word)]++
	}
	return len(wordCount)
}
