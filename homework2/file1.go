package main

import (
	"fmt"
	"math"
)

func main() {
	// Задание 1
	var a float64
	fmt.Scan(&a)
	P := 4 * a
	fmt.Printf("Периметр a: %.2f\n", P)
	// Задание 2
	fmt.Scan(&a)
	S := math.Pow((a), 2)
	fmt.Printf("Площадь S: %.2f\n", S)
	// Задание 3
	var b float64
	fmt.Scan(&a, &b)
	S = a * b
	P = (a + b) * 2
	fmt.Printf("Периметр: %.2f Площадь: %.2f\n", P, S)
	// Задание 4
	var d float64
	fmt.Scan(d)
	L := math.Pi * d
	fmt.Println(L)
	// Задание 5
	var V float64
	fmt.Scan(&a)
	V = math.Pow(a, 3)
	S = 6 * a
	fmt.Printf("Обьем куба: %.2f Площадь поверхности: %.2f\n", V, S)
	// Задание 6
	var c float64
	fmt.Scan(&a, &b, &c)
	V = a * b * c
	S = 2 * (a*b + b*c + a*c)
	fmt.Printf("Обьем параллелепипеда: %.2f Площадь поверхности: %.2f\n", V, S)
	// Задание 7
	var R float64
	fmt.Scan(&R)
	L = 2 * math.Pi * R
	S2 := math.Pi * R
	fmt.Printf("Длина окружности: %.2f Площадь круга: %f\n", L, S2)
	// Задание 8
	fmt.Scan(&a, &b)
	fmt.Printf("Среднее арифметическое a и b: %.2f\n", (a+b)/2)
	// Задание 9
	fmt.Scan(&a, &b)
	fmt.Printf("Среднее геометрическое a и b: %.2f\n", math.Sqrt(a*b))
	// Задание 10
	var number1, number2 float64
	fmt.Scan(&number1, &number2)
	fmt.Printf("Сумма: %.2f Разность: %.2f Произведение: %.2f Частное %.2f\n",
		number1*number1+number2*number2, number1*number1-number2*number2, (number1*number1)*(number2*number2), (number1*number1)/(number2*number2))
	// Задание 11
	fmt.Scan(&L)
	fmt.Printf("Количество метров: %d\n", int(L)/100)
	// Задание 12
	var M int
	fmt.Scan(&M)
	fmt.Printf("Количество тонн: %d\n", M/1000)
	// Задание 13
	var F int
	fmt.Scan(&F)
	fmt.Printf("Количество килобайт: %d\n", F/1024)
	// Задание 14
	var A, B int
	fmt.Scan(&A, &B)
	fmt.Printf("Количество возможных отрезков B на отрезке A: %d\n", A/B)
	// Задание 15
	fmt.Scan(&A, &B)
	fmt.Printf("Длина незанятой части A: %d\n", A%B)
}
