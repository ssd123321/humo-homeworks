package main

type Grade = int

func CheckGrade(g Grade) string {
	if g >= 50 {
		return "Оценка удовлетворительна"
	}
	return "Оценка неудовлетворительна"
}
