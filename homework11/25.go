package main

import (
	"sort"
	"strings"
)

func sortWordsByLength(input string) []string {
	words := strings.Fields(input)
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	return words
}
