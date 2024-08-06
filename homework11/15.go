package main

import "strings"

func CountWords(s1 string) int {
	return len(strings.Split(s1, " "))
}
