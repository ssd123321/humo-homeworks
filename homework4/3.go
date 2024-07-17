package main

import "fmt"

func PrintTrafficLight() {
	var lightColor string
	fmt.Scan(&lightColor)
	switch lightColor {
	case "красный":
		fmt.Println("Стоп")
	case "желтый":
		fmt.Println("Внимание")
	case "зеленый":
		fmt.Println("Идите")
	}
}
