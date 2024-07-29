package main

func multiplyArray(arr []int, factor int) []int {
	result := []int{}
	for _, value := range arr {
		result = append(result, value*factor)
	}
	return result
}
