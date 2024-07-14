package main

// Функция - как аргумент функции
func Calculate(n1 int, n2 int, action func(value1, value2 int) int) int {
	return action(n1, n2)
}
func Execute(b bool, action func(b bool)) {
	action(b)
}
func simplefunc(n int) int {
	return n * 2
}
func Apply(n int, action func(n int) int) int {
	return action(n)
}