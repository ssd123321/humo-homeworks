package main

import "fmt"

func CheckAge(age int) {
	if age >= 18 {
		fmt.Println("Совершеннолетний")
	} else {
		fmt.Println("Несовершеннолетний")
	}
}
