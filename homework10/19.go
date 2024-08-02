package main

type Trip struct {
	distance    float64
	costPerMile float64
}

func (t *Trip) SetCostPerMile(cost float64) {
	t.costPerMile = cost
}
func (t Trip) TotalCoast() float64 {
	return t.costPerMile * t.distance
}
