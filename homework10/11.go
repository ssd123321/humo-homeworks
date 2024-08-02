package main

type Employee struct {
	name   string
	salary float64
}
type Payroll struct {
	employees []Employee
}

func (p *Payroll) AddEmployee(name string, salary float64) {
	p.employees = append(p.employees, Employee{name, salary})
}
func (p Payroll) AverageSalary() float64 {
	var salary float64
	for _, employee := range p.employees {
		salary += employee.salary
	}
	return salary / float64(len(p.employees))
}
