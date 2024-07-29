package main

func moveZerosToEnd(arr []int) []int {
	nonZero := []int{}
	for _, value := range arr {
		if value != 0 {
			nonZero = append(nonZero, value)
		}
	}
	for i := len(nonZero); i < len(arr); i++ {
		nonZero = append(nonZero, 0)
	}
	return nonZero
}
