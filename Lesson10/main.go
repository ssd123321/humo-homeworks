package main

import "fmt"

type MySlice []int

type Rectangle struct {
	width, length float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.length
}
func (r *Rectangle) Scale(factor float64) {
	r.width *= factor
	r.length *= factor

}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

type Person struct {
	name string
	age  int
}
type Employee struct {
	Person
	position string
}
func (p Person) introduce() {
	fmt.Printf("Hi I'm %s and I'm %d old\n", p.name, p.age)
}
func (e Employee) introduce() {
	fmt.Printf("Hi I'm %s, a %s and I'm %d old\n", e.name, e.position, e.age)
}
func main() {
	r := Rectangle{10, 10}
	r.Scale(2)
	fmt.Printf("%.2f\n", r.Area())
	var e = Employee{Person{"e;fc", 13}, "fewfkwe"}
	e.introduce()

}
