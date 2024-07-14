package main

import "fmt"

// В функции передаются аргументы, но функция ничего не возвращает
func PrintNumber(number int) {
	fmt.Println(number)
}
func GreetUser(username string) {
	fmt.Println(username)
}
func ToggleLight(value bool) {
	if value {
		fmt.Println("Light on")
	} else {
		fmt.Println("Light off")
	}
}
