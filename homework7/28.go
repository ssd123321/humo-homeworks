package main

import "fmt"

type Course struct {
	Title      string
	Instructor Person
}

func GetCourse(c Course) {
	fmt.Printf("Название курса: %s, Инструктор: %s %s %d лет\n", c.Title, c.Instructor.FirstName, c.Instructor.LastName, c.Instructor.Age)
}
