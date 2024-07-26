package main

type BMI float64

func WhichBMI(b BMI) string {
	if b > 25 {
		return "Obesity"
	} else if b >= 23 && b <= 24.9 {
		return "Overweight"
	} else if b >= 18.5 && b <= 22.9 {
		return "Normal weight"
	}
	return "Underweight"
}
