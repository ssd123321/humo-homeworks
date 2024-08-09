package main
func GetAllValues(m map[string]int)[]int{
	x := make([]int, 0)
	for _, value := range m{
		x = append(x, value)
	}
	return x
}