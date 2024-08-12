package main

import (
	"fmt"
)

type TaxCalculator interface {
	CalculateIncomeTax(income float64) float64
	CalculateVAT(amount float64) float64
}

type SimpleTaxCalculator struct {
	incomeTaxRate float64
	vatRate       float64
}

func (c *SimpleTaxCalculator) CalculateIncomeTax(income float64) float64 {
	return income * c.incomeTaxRate
}

func (c *SimpleTaxCalculator) CalculateVAT(amount float64) float64 {
	return amount * c.vatRate
}

func calculateTaxes(c TaxCalculator, income, amount float64) {
	incomeTax := c.CalculateIncomeTax(income)
	vat := c.CalculateVAT(amount)
	fmt.Printf("Income Tax: %.2f, VAT: %.2f\n", incomeTax, vat)
}

