package main

func swapMaxMin(arr []int) []int {
	maxIndex := 0
	minIndex := 0
	for i, value := range arr {
		if value > arr[maxIndex] {
			maxIndex = i
		}
		if value < arr[minIndex] {
			minIndex = i
		}
	}
	arr[maxIndex], arr[minIndex] = arr[minIndex], arr[maxIndex]
	return arr
}
