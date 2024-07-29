package main

func countPositive(arr []int) int {
	count := 0
	for _, value := range arr {
		if value > 0 {
			count++
		}
	}
	return count
}
