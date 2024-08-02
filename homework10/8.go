package main

type Student2 struct {
	name string
	age  int
}
type Classroom struct {
	Student []Student2
}

func (c *Classroom) AddStudent(name string, age int) {
	c.Student = append(c.Student, Student2{name: name, age: age})
}
func (c Classroom) AverageAge() float64 {
	var sum float64
	for _, student := range c.Student {
		sum += float64(student.age)
	}
	return sum / float64(len(c.Student))
}
