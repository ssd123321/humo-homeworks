package main

import "strings"
func CountWords(str string)map[string]int{
	words := strings.Fields(str)
	m := make(map[string]int)
	for _, value := range words{
		m[value]++
	}
	return m
}