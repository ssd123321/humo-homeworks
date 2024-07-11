package main

import (
	"fmt"
)

func main() {
	var x int
	var x1 int8
	var x2 int16
	var x3 int32
	var x4 int64
	fmt.Println(x, x1, x2, x3, x4)
	x = 1
	x1 = 2
	x2 = 3
	x3 = 4
	x4 = 5
	fmt.Println(x, x1, x2, x3, x4)
	var y float32
	var y1 float64
	fmt.Println(y, y1)
	y = 11.11
	y1 = 2.222222
	fmt.Println(y, y1)
	var z bool
	fmt.Println(z)
	z = true
	fmt.Println(z)
	var r rune
	fmt.Println(r)
	r = 'x'
	fmt.Println(r)
	var b byte
	fmt.Println(b)
	b = 100
	fmt.Println(b)
	var s string
	fmt.Println(s)
	s = "Hello"
	fmt.Println(s)
	var u uint
	var u1 uint8
	var u2 uint16
	var u3 uint32
	var u4 uint64
	fmt.Println(u, u1, u2, u3, u4)
	u = 123
	u1 = 234
	u2 = 345
	u3 = 456
	u4 = 567
	fmt.Println(u, u1, u2, u3, u4)
	var c complex64
	var c1 complex128
	fmt.Println(c, c1)
	c = 1.1233 + 1.2313i
	c1 = 1.2414242 + 2.424242374i
	fmt.Println(c, c1)

}
