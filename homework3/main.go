package main

import (
	"fmt"
)

func main() {
	// В функции не передаются аргументы и ничего не возвращает фукция
	PrintGreeting()
	DisplaySeparator()
	NotifyStart()
	// В функции не передаются аргументы, но функция возвращает значение (значения)
	fmt.Println(GetWelcomeMessage())
	fmt.Println(GetPiValue())
	fmt.Println(IsServerActive())
	// В функции передаются аргументы, но функция ничего не возвращает
	PrintNumber(123)
	GreetUser("Said")
	ToggleLight(true)
	// В функции передаются аргументы, но функция возвращает значение (значения)
	fmt.Println(Add(2, 3))
	fmt.Println("Hello, ", "World!")
	fmt.Println(IsEven(7))
	// Функция - как переменная
	var adder func(n1, n2 int) int
	adder = func(n1, n2 int) int {
		return n1 + n2
	}
	adder(7, 9)
	var concatenator func(str1, str2 string) string
	concatenator = func(str1, str2 string) string {
		return str1 + str2
	}
	concatenator("Hello, ", "Guys!")
	var isPositive func(number int) bool
	isPositive = func(number int) bool {
		if number > 0 {
			return true
		} else {
			return false
		}
	}
	isPositive(-9)
	// Функция - как аргумент функции
	fmt.Println(Calculate(23, 55, Add))
	Execute(true, ToggleLight)
	fmt.Println(Apply(32, simplefunc))
	// Функция – как значение функции (callback)
	MultiplyBy32 := Multiplier(32)
	fmt.Println(MultiplyBy32(32))
	StrRepeaterBy5 := StringRepeater(5)
	fmt.Println(StrRepeaterBy5("Hello!"))
	Printer := ConditionalPrinter(true)
	Printer("It's true")

}
