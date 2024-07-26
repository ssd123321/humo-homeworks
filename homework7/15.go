package main

type BloodPressure = float64

func CheckHealth(bmi BMI, bp BloodPressure) string {
	if bmi > 25 || bp > 130 {
		return "Unhealthy"
	} else if (bmi >= 23 && bmi <= 24.9) || bp > 120 {
		return "At risk"
	} else {
		return "Healthy"
	}
}
