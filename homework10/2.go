package main

type Student struct {
	Grades []int
}

func (s *Student) AddGrade(grade int) {
	s.Grades = append(s.Grades, grade)
}
func (s Student) AverageGrade() float64 {
	sum := 0
	for _, value := range s.Grades {
		sum += value
	}
	return float64(sum) / float64(len(s.Grades))
}
