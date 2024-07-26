package main

import "fmt"

type Count = int

func ReverseCount(value int) {
	for ; value >= 0; value-- {
		fmt.Println(value)
	}
}
