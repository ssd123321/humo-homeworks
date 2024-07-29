package main

func removeNumber(arr []int, num int) []int {
	result := []int{}
	for _, value := range arr {
		if value != num {
			result = append(result, value)
		}
	}
	return result
}
