package main

type Operation func(n1, n2 int) int

func Sum(n1, n2 int) int {
	return n1 + n2
}
func Minus(n1, n2 int) int {
	return n1 - n2
}
func Multiply(n1, n2 int) int {
	return n1 * n2
}
func Combined(n1 int, n2 int, array [3]Operation) [3]int {
	var results [3]int
	for i, function := range array {
		results[i] = function(n1, n2)
	}
	return results
}
