package main

func copyArray(arr []int) []int {
	result := []int{}
	for _, value := range arr {
		result = append(result, value)
	}
	return result
}
