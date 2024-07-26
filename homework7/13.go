package main

type Weight = float64

func CheckWeight(w Weight) string {
	if w > 80 {
		return "Heavy"
	} else if w >= 60 && w <= 80 {
		return "Medium"
	}
	return "Light"
}
