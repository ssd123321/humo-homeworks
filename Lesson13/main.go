package main

import (
	"fmt"
)

type UserType interface {
	GetUserInfo() string
	GetUserName() string
	GetUserAge() int
}
type User struct {
	Name string
	Age  int
}
type Auto interface {
	Move()
	Stop()
	Wait()
}
type Animal interface {
	Speak() string
}
type Dog struct {
	Sound string
}

func (d Dog) Speak() string {
	return d.Sound
}

type Cow struct {
	Sound string
}

func (c Cow) Speak() string {
	return c.Sound
}
func (u *User) GetUserInfo() string {
	return fmt.Sprintf("Name %s, Age %d", u.Name, u.Age)
}
func (u *User) GetUserName() string {
	return u.Name
}
func (u *User) GetUserAge() int {
	return u.Age
}

type User2 struct {
	Name    string
	Age     int
	Address string
}

func (u User2) GetUserInfo() string {
	return fmt.Sprintf("Name %s, Age %d, Address %s", u.Name, u.Age, u.Address)
}
func (u User2) GetUserName() string {
	return u.Name
}
func (u User2) GetUserAge() int {
	return u.Age + 4
}

type Shape interface {
	Area() float64
}
type Rectangle struct {
	x1, x2 float64
}

func (r Rectangle) Area() float64 {
	return r.x1 * r.x2
}

type Square struct {
	x1, x2 float64
}

func (s Square) Area() float64 {
	return s.x1 * s.x2
}
func PrintShapeInfo(s Shape) {
	switch v := s.(type) {
	case Rectangle:
		fmt.Printf("%v value is rectangle\n", v)
	case Square:
		fmt.Printf("%v value is rectangle\n", v)
	}
}
func main() {
	/*
		x := User{
			Name: "cwcmwc",
			Age:  47,
		}
		x2 := User2{
			Name:    "cmmcc",
			Age:     3939,
			Address: "ckckmckmc",
		}
		GetUser(&x)
		GetUser(x2)
		c := Cow{Sound: "Mooo"}
		d := Dog{Sound: "Meow"}
		GetAnimalSound(c)
		GetAnimalSound(d)
	*/

	var m interface{}
	m = 43
	m = "Hello"
	// Утверждение типа (type assertion) в GO
	value, ok := m.(string)
	fmt.Println(value, ok)
	x := Square{}
	x2 := Rectangle{}
	PrintShapeInfo(x)
	PrintShapeInfo(x2)

}
func GetUser(x UserType) {
	fmt.Println(x.GetUserAge(), x.GetUserName(), x.GetUserInfo())
}
func GetAnimalSound(a Animal) {
	fmt.Println(a.Speak())
}
