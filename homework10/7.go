package main

type Temperature struct {
	celsius float64
}

func (t Temperature) ToFahrenheit() float64 {
	return t.celsius*1.8 + 32
}
func (t Temperature)ToKelvin()float64{
	return t.celsius + 273.15
}
