package main

type Temperature float64

func CheckTemperature(t Temperature) string {
	if t > 0 {
		return "Больше нуля"
	} else if t < 0 {
		return "Меньше нуля"
	}
	return "Равно нулю"
}
