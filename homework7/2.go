package main

type Age int

func ReturnAge(a Age) string {
	if a >= 13 && a <= 19 {
		return "Подросток"
	}
	return "Неподросток"
}
