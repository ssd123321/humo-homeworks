package main

import (
	"fmt"
	"math"
)

func main() {
	// 1
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	// 2
	for i := 1; i <= 5; i++ {
		fmt.Println(i * i)
	}
	// 3
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * 3 = %d\n", i, i*3)
	}
	// 4
	a, b := 0, 1
	for i := 0; i < 10; i++ {
		fmt.Println(a)
		a, b = b, a+b
	}
	// 5
	num1, num2 := 24, 18
	for num2 != 0 {
		num1, num2 = num2, num1%num2
	}
	fmt.Println("НОД:", num1)
	// 6
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
	// 7
	num := 17
	isPrime := true
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime && num > 1 {
		fmt.Println(num, "является простым числом")
	} else {
		fmt.Println(num, "не является простым числом")
	}
	// 8
	num = 24
	fmt.Print("Делители", num, ": ")
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
	// 9
	num = 12345
	sum := 0
	for num > 0 {
		digit := num % 10
		sum += digit
		num /= 10
	}
	fmt.Println("Сумма цифр:", sum)
	// 10
	var num10 int
	for {
		fmt.Print("Введите положительное число: ")
		fmt.Scanln(&num10)
		if num10 > 0 {
			break
		}
		fmt.Println("Ошибка: число должно быть положительным.")
	}
	fmt.Println("Вы ввели:", num10)
	// 11
	var n0 int
	fmt.Print("Введите число n: ")
	fmt.Scanln(&n0)
	product := 1
	for i := 1; i <= n0; i++ {
		product *= i
		if product > 1000 {
			fmt.Println("Произведение превышает 1000, выполнение прекращено.")
			break
		}
	}
	fmt.Println("Произведение:", product)
	// 12
	var num12 int
	fmt.Print("Введите число: ")
	fmt.Scanln(&num12)
	original := num12
	reversed := 0
	for num12 > 0 {
		digit := num12 % 10
		reversed = reversed*10 + digit
		num12 /= 10
	}
	if original == reversed {
		fmt.Println("Число является палиндромом")
	} else {
		fmt.Println("Число не является палиндромом")
	}
	// 13
	var height int
	fmt.Print("Введите высоту пирамиды: ")
	fmt.Scanln(&height)
	for i := 1; i <= height; i++ {
		for j := 1; j <= height-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	// 14
	var n1 int
	fmt.Print("Введите число n: ")
	fmt.Scanln(&n1)
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n1; j++ {
			fmt.Printf("%d x %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
	// 15
	var n int
	fmt.Print("Введите высоту треугольника Паскаля: ")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var prev int = 1
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				fmt.Print(1, " ")
			} else {
				curr := prev + prev
				fmt.Print(curr, " ")
				prev = curr
			}
		}
		fmt.Println()
	}

	// 16
	for {
		var num int
		fmt.Print("Enter a number (negative to exit): ")
		fmt.Scanln(&num)

		if num < 0 {
			break
		}

		factorial := 1
		for i := 1; i <= num; i++ {
			factorial *= i
		}

		fmt.Println("Factorial of", num, "is", factorial)
	}
	// 17
	for {
		var num1, num2 int
		fmt.Print("Enter two numbers: ")
		fmt.Scanln(&num1, &num2)
		fmt.Println("The sum is", num1+num2)
	}
	// 18
	for i := 1; i <= 100; i++ {
		if i%int(math.Sqrt(float64(i))) != 0 {
			fmt.Println(i)
		}
	}
	// 19
	for i := 2; i <= 50; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(i)
		}
	}
	// 20
	for i := 1; i <= 30; i++ {
		if i%2 == 0 || i%3 == 0 {
			continue
		}
		fmt.Println(i)
	}
	// 21
	for i := 1; i <= 50; i++ {
		if i*i*i <= 50 {
			fmt.Println(i)
		} else {
			break
		}
	}
	// 22
	sum1 := 0
	for i := 1; i <= 100; i++ {
		if sum1+i > 200 {
			break
		}
		fmt.Println(i)
		sum1 += i
	}
	// 23
	for {
		var num int
		fmt.Print("Enter a number: ")
		fmt.Scanln(&num)

		if num%7 == 0 {
			break
		}
	}
}
