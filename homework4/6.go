package main

import "fmt"

func GetTemperatureDescription() string {
	var temperature float64
	fmt.Scan(&temperature)
	if temperature < 10 {
		return "Холодно"
	} else if temperature >= 10 && temperature <= 25 {
		return "Тепло"
	}
	return "Жарко"
}
