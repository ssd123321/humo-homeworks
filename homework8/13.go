package main

func secondMax(arr []int) int {
	max1 := arr[0]
	max2 := arr[0]
	for _, value := range arr {
		if value > max1 {
			max2 = max1
			max1 = value
		} else if value > max2 && value != max1 {
			max2 = value
		}
	}
	return max2
}
