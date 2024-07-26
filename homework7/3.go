package main

type Speed float64

func IsSpeedOk(s Speed) string {
	if s >= 0 && s <= 120 {
		return "Скорость допустима"
	}
	return "Скорость недопустима"
}
