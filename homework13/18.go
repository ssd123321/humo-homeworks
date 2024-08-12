package main

import (
    "fmt"
)

type Payroll interface {
    AddEmployee(name string, salary float64)
    GetSalary(name string) float64
    TotalPayroll() float64
}

type CompanyPayroll struct {
    employees map[string]float64
}
func (p *CompanyPayroll) AddEmployee(name string, salary float64) {
    p.employees[name] = salary
}

func (p *CompanyPayroll) GetSalary(name string) float64 {
    return p.employees[name]
}

func (p *CompanyPayroll) TotalPayroll() float64 {
    var total float64
    for _, salary := range p.employees {
        total += salary
    }
    return total
}

func managePayroll(p Payroll) {
    p.AddEmployee("John Doe", 5000.0)
    p.AddEmployee("Jane Smith", 6000.0)
    p.AddEmployee("Bob Johnson", 4500.0)

    fmt.Printf("John Doe's Salary: %.2f\n", p.GetSalary("John Doe"))
    fmt.Printf("Total Payroll: %.2f\n", p.TotalPayroll())
}


