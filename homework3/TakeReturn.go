package main

// В функции передаются аргументы, но функция возвращает значение (значения)
func Add(number1, number2 int) int {
	return number1 + number2
}
func Concat(str1, str2 string) string {
	return str1 + str2
}
func IsEven(number int) bool {
	if number%2 == 0 {
		return true
	} else {
		return false
	}
}
