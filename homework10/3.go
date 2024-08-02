package main

import "math"

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.radius
}
