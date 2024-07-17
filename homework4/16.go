package main

func ApplyOperation(n1, n2 int, action func(n1, n2 int) int) int {
	if n1 < 0 {
		n1 *= -1
	}
	if n2 < 0 {
		n2 *= -1
	}
	return action(n1, n2)
}
func action(n1, n2 int) int {
	return n1 + n2
}
