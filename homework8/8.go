package main

func findIndices(arr []int, num int) []int {
	indices := []int{}
	for i, value := range arr {
		if value == num {
			indices = append(indices, i)
		}
	}
	return indices
}
