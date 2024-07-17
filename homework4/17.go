package main

import "fmt"

func CheckCondition(n int, action func(int) bool) {
	if action(n) {
		fmt.Println("Условие выполнено")
	} else {
		fmt.Println("Условие невыполнено")
	}
}
func IsPositive(n int) bool {
	return n > 0
}
