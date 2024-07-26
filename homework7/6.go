package main

type UnaryOperation func(int) int

func SquareNumber(x int) int {
	return x * x
}
