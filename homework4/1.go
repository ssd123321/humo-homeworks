package main

import "fmt"

func PrintGreeting() {
	var dayType string
	fmt.Scan(&dayType)
	switch dayType {
	case "Утро":
		fmt.Println("Доброе утро!")
	case "День":
		fmt.Println("Добрый день!")
	case "Вечер":
		fmt.Println("Добрый вечер!")
	}
}
