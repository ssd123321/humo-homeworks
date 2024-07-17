package main

import "fmt"

func main() {
	PrintGreeting()
	PrintWeather()
	PrintTrafficLight()
	fmt.Println(GetGrade())
	fmt.Println(GetDiscount())
	fmt.Println(GetTemperatureDescription())
	CheckNumber(32)
	CheckAge(23)
	CheckPassword("4234")
	fmt.Println(Add(2, 3))
	fmt.Println(CompareStrings("ttr", "ltr"))
	fmt.Println(Max(23, 45))
	fmt.Println(ApplyOperation(32, 22, action))
	CheckCondition(45, IsPositive)
	FormatAndPrint("hellO", formatFunc)
	x := CreateMultiplier(3)
	fmt.Println(x())
	x1 := CreateGreeter("stt")
	fmt.Println(x1("hello"))
	x2 := CreateValidator("ertyu")
	fmt.Println(x2("rrrt"))
	// Задание 13
	var operation func(n1, n2 int) int
	operation = func(n1, n2 int) int {
		if n1 >= n2 {
			return n1 - n2
		}
		return n2 - n1
	}
	fmt.Println(operation(13, 13))
	// Задание 14
	var concat func(str1, str2 string) string
	concat = func(str1, str2 string) string {
		if str1 != "" && str2 != "" {
			return str1 + " " + str2
		}
		return str1 + str2
	}
	fmt.Println(concat("jj", "dd"))
	// Задание 15
	var multiply func(n1, n2 int) int
	multiply = func(n1, n2 int) int {
		if n1 < 0 {
			n1 *= -1
		}
		if n2 < 0 {
			n2 *= -1
		}
		return n1 * n2
	}
	fmt.Println(multiply(-6, -5))
}
