package main

import "fmt"

func PrintWeather() {
	var weatherType string
	fmt.Scan(&weatherType)
	switch weatherType {
	case "солнечно":
		fmt.Println("Солнечно")
	case "облачно":
		fmt.Println("Облачно")
	case "дождливо":
		fmt.Println("Дождливо")
	}
}
