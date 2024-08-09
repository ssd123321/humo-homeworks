package main

func CheckIdentity(s []int) int {
	m := make(map[int]int)
	for _, value := range s {
		m[value] = 0
	}
	return len(m)
}
