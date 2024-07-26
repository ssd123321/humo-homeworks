package main

type Check func(int) bool

func CheckNumber(x int) bool {
	return x%2 == 0
}
