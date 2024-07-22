package main

import "fmt"

func main() {
	var x int
	y := &x
	*y = 42
	fmt.Println(*y)
	fmt.Println(y)
	changeValue(y, 55)
	fmt.Println(x)
	AddValue(y, 5)
	fmt.Println(x)
	multiplyValue(y, 2)
	fmt.Println(x)
	var s float64 = 45.23
	t := &s
	n := &t
	fmt.Println(n)
	fmt.Println(*n)
	fmt.Println(**n)
}
func changeValue(x *int, value int) {
	*x = value
}
func AddValue(x *int, value int) {
	*x += value
}
func multiplyValue(x *int, value int) {
	*x *= value
}
