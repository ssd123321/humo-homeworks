package main

import "fmt"

func CheckPassword(password string) {
	if password == "1234" {
		fmt.Println("Пароль верный")
	} else {
		fmt.Println("Пароль неверный")
	}
}
