package main

func combineArrays(arr1 []int, arr2 []int) []int {
	result := []int{}
	for _, value := range arr1 {
		result = append(result, value)
	}
	for _, value := range arr2 {
		result = append(result, value)
	}
	return result
}
