package main

func mergeSortedArrays(arr1 []int, arr2 []int) []int {
	result := []int{}
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}
	for ; i < len(arr1); i++ {
		result = append(result, arr1[i])
	}
	for ; j < len(arr2); j++ {
		result = append(result, arr2[j])
	}
	return result
}
