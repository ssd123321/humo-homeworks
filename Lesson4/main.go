package main

import "fmt"

func main() {
	// Условные конструкции
	n := 5
	b := n > 0
	fmt.Println(b)
	x := 0
	b1 := (x > 0) || (x < 0)
	fmt.Println(b1)
	// if-else
	num := 15
	if num > 9 && num < 100 {
		fmt.Println("Двухзначное число")
	}
	num = 6
	if num%2 != 0 {
		fmt.Println("Число нечетное")
	} else {
		fmt.Println("Число четное")
	}
	// if -else
	num2 := 10
	if num2 >= 0 && num2 < 10 {
		fmt.Println("Число однозначное ")
	} else if num2 >= 10 && num2 < 100 {
		fmt.Println("Число двухзначное")
	} else if num2 >= 100 && num2 < 1000 {
		fmt.Println("Число трехзначное")
	} else if num2 >= 1000 && num2 < 10000 {
		fmt.Println("Число четырехзначное")
	} else {
		fmt.Println("Число слишком большое")
	}
	// Задача
	var num5, num6 int
	num5 = 10
	num6 = 9
	if num5 > num6 {
		fmt.Println(num5)
	} else if num6 > num5 {
		fmt.Println(num6)
	} else {
		fmt.Println(num6)
	}
	// Задача 2
	//var number1, number2, number3 int
	//number1 = 8
	//number2 = 19
	//number3 = 11
	/*
		if number1 < number2 && number1 < number3 {
			if number2 < number3 {
				fmt.Println(number3, number2, number1)
			} else {
				fmt.Println(number2, number3, number1)
			}
		} else if number2 < number1 && number2 < number3 {
			if number1 > number3 {
				fmt.Println(number3, number1, number2)
			} else {
				fmt.Println(number2, number3, number1)
			}
		} else if number3 < number2 && number3 < number1 {
			if number2 < number1 {
				fmt.Println(number1, number2, number3)
			} else {
				fmt.Println(number2, number1, number3)
			}
		} else if number1 == number2 && number2 == number1 {
			fmt.Println(number1, number2, number3)
		} */
	/*if number1 >= number2 && number1 >= number3 && number2 <= number3 {
		fmt.Println(number1, number3, number2)
	} else if number1 >= number2 && number1 >= number3 && number2 >= number3 {
		fmt.Println(number1, number2, number3)
	} else if number2 >= number1 && number2 >= number3 && number3 >= number1 {
		fmt.Println(number2, number3, number1)
	} else if number2 >= number1 && number2 >= number3 && number3 <= number1 {
		fmt.Println(number2, number1, number3)
	} else if number3 >= number2 && number3 >= number1 && number2 <= number1 {
		fmt.Println(number3, number1, number2)
	} else if number3 >= number2 && number3 >= number1 && number2 >= number1 {
		fmt.Println(number3, number2, number1)*
	}*/
	a := 7
	switch a {
	case 0:
		fmt.Println("a =", a)
	case 8:
		fmt.Println("a =", a)
	case 9:
		fmt.Println("a =", a)

	}
	c := 7
	switch c + 2 {
	case 9:
		fmt.Println("c + 2 is 9")
	case 6:
		fmt.Println("c + 2 is 6")
	}
	i := -1
	switch i {
	case 1:
		fmt.Println("Positive")
		fallthrough
	case -1:
		fmt.Println("Negative")
	case 0:
		fmt.Println("Zero")
	}
	// Задача

	var n1, n2, n3 int
	n1 = 9
	n2 = 9
	n3 = 9
	if n1 >= n3 && n3 >= n2 {
		fmt.Println(n1 + n2)
	} else if n2 >= n1 && n3 >= n1 {
		fmt.Println(n2 + n3)
	} else if n3 >= n1 && n1 >= n2 {
		fmt.Println(n3 + n2)
	}
	/*
		var x1 int
		fmt.Scan(&x1)
		if x1%4 == 0 {
			if x1%100 == 0 && x1%400 != 0 {
				fmt.Println("Невысокосный")
			} else {
				fmt.Println("Высокосный")
			}
		} else {
			fmt.Println("Невысокосный")
		} */
	var x2 int
	fmt.Scan(&x2)
	if x2%4 == 0 && x2%100 == 0 && x2%400 != 0 {
		fmt.Println("Невысокосный")
	} else if x2%4 == 0 {
		fmt.Println("Высокосный")
	} else {
		fmt.Println("Невысокосный")
	}
	// Задача
	var count int
	fmt.Scan(&count)
	switch count {
	case 1:
		fmt.Printf("January has %d days\n", 31)
	case 2:
		fmt.Printf("February has %d days\n", 29)
	case 3:
		fmt.Printf("March has %d days\n", 31)
	case 4:
		fmt.Printf("April has %d days\n", 30)
	case 5:
		fmt.Printf("May has %d days\n", 31)
	case 6:
		fmt.Printf("June has %d days\n", 30)
	case 7:
		fmt.Printf("July has %d days\n", 31)
	case 8:
		fmt.Printf("August has %d days\n", 31)
	case 9:
		fmt.Printf("September has %d days\n", 30)
	case 10:
		fmt.Printf("October has %d days\n", 31)
	case 11:
		fmt.Printf("November has %d days\n", 30)
	case 12:
		fmt.Printf("December has %d days\n", 31)

	}

}
