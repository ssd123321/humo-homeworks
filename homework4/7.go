package main

import "fmt"

func CheckNumber(number int) {
	if number > 0 {
		fmt.Println("Положительное")
	} else if number == 0 {
		fmt.Println("Ноль")
	} else {
		fmt.Println("Отрицательное")
	}
}
