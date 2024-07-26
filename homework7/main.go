package main

import (
	"fmt"
)

func main() {
	// 1
	fmt.Println(CheckTemperature(32.3))
	// 2
	fmt.Println(ReturnAge(20))
	// 3
	fmt.Println(IsSpeedOk(34))
	// 4
	fmt.Println(CheckScore(54))
	// 5
	fmt.Println(WhichBMI(24))
	// 6
	fmt.Println(SquareNumber(5))
	// 7
	fmt.Println(DoubleNumber(25))
	// 8
	fmt.Println(CheckNumber(33))
	// 9
	fmt.Println(IsPositive(11))
	// 10
	var func1, func2, func3 Operation
	func1 = Sum
	func2 = Minus
	func3 = Multiply
	var massive [3]Operation
	massive[0] = func1
	massive[1] = func2
	massive[2] = func3
	answers := Combined(34, 4, massive)
	for _, answer := range answers {
		fmt.Println(answer)
	}
	// 11
	ReverseCount(90)
	// 12
	fmt.Println(CheckBatteryLevel(90))
	// 13
	fmt.Println(CheckWeight(75))
	// 14
	fmt.Println(CheckGrade(80))
	// 15
	fmt.Println(CheckHealth(23, 110))
	// 16
	p := Product{"Chair", 120}
	GetProduct(p)
	// 17
	person := Person{"Said", "Mirzoev", 20}
	PrintPerson(person)
	// 18
	p1 := Product{"Car", 20000}
	p2 := Product{"Bicycle", 5000}
	fmt.Println(GetMoreExpensiveProduct(p1, p2))
	// 19
	item1 := Item{"Pen", 32}
	item2 := Item{"Pencil", 40}
	item3 := Item{"Car", 50}
	fmt.Println(GetQuantities(item1, item2, item3))
	// 20
	book1 := Book{"Pirati", "Dostoevsky"}
	book2 := Book{"Moryaki", "Pushkin"}
	book3 := Book{"Cheloveki", "Lermontov"}
	l := Library{"Super library", []Book{book1, book2, book3}}
	CheckLibrary(l)
	// 21
	ChangePriceOfProduct(&p, 150)
	fmt.Println(p)
	// 22
	IncreaseAge(&person)
	fmt.Println(person)
	// 23
	ChangeProduct(&p, "Wheelchair", 350)
	fmt.Println(p)
	// 24
	IncreaseItemQuantity(&item1, 40)
	fmt.Println(item1)
	// 25
	Task1 := Task{"Kill", true}
	Task2 := Task{"Reach 1000 level", false}
	Task3 := Task{"Reach Boss", false}
	y := make([]Task, 3)
	y[0] = Task1
	y[1] = Task2
	y[2] = Task3
	ChangeStatus(&y, "Reach Boss")
	fmt.Println(y)
	// 26
	person2 := Person2{"Said", Address{"Rudaki", "Dushanbe"}}
	GetPersonsAddress(person2)
	// 27
	c := Company{"Apple", Address{"Some street", "New York"}}
	GetCompanyAddress(c)
	// 28
	course := Course{"C++", Person{"Bek", "Bekov", 23}}
	GetCourse(course)
	// 29
	e := Event{"Party", Address{"Rudaki", "Dushanbe"}}
	GetEvent(e)
	// 30
	project := Project{"Supplying", Person2{"Someone", Address{"Rudaki", "Dushanbe"}}}
	GetProject(project)
	// 31
	TwoNodes()
	// 32
	ThreeNodes()
	// 33
	s5 := &Node{36, nil}
	s4 := &Node{35, s5}
	s3 := &Node{34, s4}
	s2 := &Node{33, s3}
	s1 := &Node{32, s2}
	GetAllValues(s1)
	// 34
	AddNode(s1)
	GetAllValues(s1)
	// 35
	DeleteNode(s1, 33)
	GetAllValues(s1)
	// 36
	var bicycle struct {
		Name  string
		Price float64
	}
	bicycle = struct {
		Name  string
		Price float64
	}{"BMW", 90000}
	PrintBicycle(bicycle)
	// 37
	Person := struct {
		FirstName string
		LastName  string
		Age       int
	}{
		FirstName: "John",
		LastName:  "Johnson",
		Age:       47,
	}
	PrintPersons(Person)
	// 38
	var Car = struct {
		Make  string
		Model string
		Year  int
	}{
		Make:  "BMW",
		Model: "X5",
		Year:  2025,
	}
	PrintCar(Car)
	// 39
	book := struct {
		Title  string
		Author string
	}{
		Title:  "War and peace",
		Author: "Lermontov",
	}
	PrintBook(book)
	// 40
	Event := struct {
		Title    string
		Date     string
		Location struct {
			Street string
			City   string
		}
	}{
		Title: "Holiday",
		Date:  "19 Мая 7 часов",
		Location: struct {
			Street string
			City   string
		}{"Rudaki", "Dushanbe"},
	}
	PrintEvent(Event)
}
func PrintBicycle(car struct {
	Name  string
	Price float64
}) {
	fmt.Printf("Название: %s, Цена: %.2f\n", car.Name, car.Price)
}
func PrintPersons(Person struct {
	FirstName string
	LastName  string
	Age       int
}) {
	fmt.Printf("Имя: %s, Фамилия: %s, Возраст: %d\n", Person.FirstName, Person.LastName, Person.Age)
}
func PrintCar(Car struct {
	Make  string
	Model string
	Year  int
}) {
	fmt.Printf("Фирма: %s, Модель: %s, Год: %d\n", Car.Make, Car.Model, Car.Year)
}
func PrintBook(b struct {
	Title  string
	Author string
}) {
	fmt.Printf("Название книги: %s, Автор: %s\n", b.Title, b.Author)
}
func PrintEvent(Event struct {
	Title    string
	Date     string
	Location struct {
		Street string
		City   string
	}
}) {
	fmt.Printf("Название: %s, Дата: %s, Город: %s, Улица %s\n", Event.Title, Event.Date, Event.Location.City, Event.Location.Street)
}
