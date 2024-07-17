package main

import "fmt"
func GetGrade()string{
	var score int
	fmt.Scan(&score)
	if score >= 90{
		return "A"
	}else if score >= 75 && score < 90{
		return "B"
	}else if score < 75{
		return "C"
	}
	return "Нет такой оценки!"
}