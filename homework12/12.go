package main
func CheckKeys(m map[int]int, given int)[]int{
	x := make([]int, 0)
	for _, value := range m{
		if value > given{
			x = append(x, value)
		}
	}
	return x
}