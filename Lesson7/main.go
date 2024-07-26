package main

import "fmt"

type Age int
type Adder func(num1, num2 int) int
type Myint = int
type Person2 struct {
	name    string
	age     int
	address Address
}
type Person struct {
	name string
	age  int
}
type Address struct {
	city  string
	state string
}
type Node struct {
	Value int
	Next  *Node
}
type Price = float64
type Count = int
type Temperature float64
type User struct {
	Name string
	Age  int
}
type Product struct {
	Name  string
	Price Price
}
type Book struct {
	Title  string
	Author string
	Pages  int
}
type Rectangle struct {
	Width  float64
	Height float64
}
type Coordinate struct {
	X int
	Y int
}

func main() {
	var x Age
	fmt.Println(x)
	var y Adder
	y = func(num1, num2 int) int {
		return num1 + num2
	}

	fmt.Println(y(4, 5))

	var t Myint = 45
	fmt.Println(t)
	var p Person = Person{"Said", 20}
	fmt.Println(p)
	c := new(Person)
	c.age = 42
	c.name = "Said"
	fmt.Println(*c)
	n := &Person{"Said", 23}
	fmt.Println(n)
	fmt.Printf("%+v\n", n)
	fmt.Println(n.name, n.age)
	var P Person2 = Person2{"Said", 20, Address{"Dushanbe", "Tajikistan"}}
	fmt.Println(P)
	var a Address = Address{
		"Dushanbe",
		"Tajikistan",
	}
	fmt.Println(a)
	var p1 Person2 = Person2{"Said", 20, a}
	fmt.Println(p1)
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 3}
	n1.Next = n2
	n2.Next = n3
	current := n1
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}

	fmt.Println(NumberOne(20))
	fmt.Println(NumberTwo(34))
	fmt.Println(NumberThree(123))
	var o1 Operation
	var o2 Operation
	var o3 Operation
	o1 = Plus
	fmt.Println(o1(23, 32))
	o2 = Minus
	fmt.Println(o2(23, 11))
	o3 = Multiply
	fmt.Println(o3(2, 3))
	var c1 Comparator
	var c2 Comparator
	var c3 Comparator
	c1 = IsEqual
	fmt.Println(c1(23, 23))
	c2 = More
	fmt.Println(c2(12, 22))
	c3 = Less
	fmt.Println(c3(22, 11))
	var u1 UnaryOperation
	u1 = Double
	fmt.Println(u1(3))
	var u2 UnaryOperation
	u2 = Triple
	fmt.Println(u2(3))
	decr(54)
	fmt.Println(CheckTemperature(0))
	fmt.Println(ReturnwithDiscount(100))
	PrintUser(User{"Said", 20})
	fmt.Println(ReturnPrice(Product{"Bed", 10000}))
	PrintBook(Book{"Frozen Heart", "Mister Bin", 340})
	r := &Rectangle{20.45, 20.23}
	Sum(r)
	s1 := &Coordinate{32, 22}
	s2 := &Coordinate{32, 22}
	fmt.Println(Compare(s1, s2))
}
func Compare(x *Coordinate, y *Coordinate) string {
	if *x == *y {
		return "Равны"
	}
	return "Неравны"
}
func Sum(r *Rectangle) {
	fmt.Printf("Площадь равна %.2f\n", r.Height*r.Width)
}
func PrintBook(b Book) {
	fmt.Printf("%+v\n", b)
}
func ReturnPrice(p Product) Price {
	return p.Price
}
func PrintUser(u User) {
	fmt.Printf("%+v\n", u)
}
func ReturnwithDiscount(p Price) Price {
	return p - (p * 10 / 100)
}
func CheckTemperature(t Temperature) string {
	if t > 0 {
		return "Больше нуля"
	} else if t < 0 {
		return "Меньше нуля"
	}
	return "Равно нулю"
}
func decr(count Count) {
	for ; count >= 0; count-- {
		fmt.Println(count)
	}
}
func Plus(n1, n2 int) int {
	return n1 + n2
}
func Minus(n1, n2 int) int {
	return n1 - n2
}
func Multiply(n1, n2 int) int {
	return n1 * n2
}
func Double(n1 int) int {
	return n1 * 2
}
func Triple(n1 int) int {
	return n1 * 3
}

type UnaryOperation func(int) int
type Number int
type Scope int
type Operation func(n1, n2 int) int
type Comparator func(n1, n2 int) bool

func IsEqual(n1, n2 int) bool {
	return n1 == n2
}
func More(n1, n2 int) bool {
	return n1 > n2
}
func Less(n1, n2 int) bool {
	return n1 < n2
}
func NumberOne(a Age) string {
	if a >= 18 {
		return "Совершеннолетний"
	}
	return "Несовершеннолетний"
}
func NumberTwo(num Number) string {
	if num%2 != 0 {
		return "Нечетное число"
	}
	return "Четное число"
}
func NumberThree(score Scope) string {
	if score >= 0 && score <= 100 {
		return "В диапазоне"
	}
	return "Вне диапазона"
}
