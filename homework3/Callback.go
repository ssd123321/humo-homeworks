package main

import (
	"fmt"
	"strings"
)

// Функция – как значение функции (callback)
func Multiplier(factor int) func(n int) int {
	return func(n int) int {
		return factor * n
	}
}
func StringRepeater(n int) func(str string) string {
	return func(str string) string {
		return strings.Repeat(str, n)
	}
}
func ConditionalPrinter(condition bool) func(str string) {
	return func(str string) {
		if condition {
			fmt.Println(str)
		}
	}
}
