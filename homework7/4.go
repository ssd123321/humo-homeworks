package main

type Score int

func CheckScore(s Score) string {
	if s > 0 {
		return "Счет положителен"
	} else if s < 0 {
		return "Счет отрицателен"
	}
	return "Счет нулевой"
}
