package main

func removeDuplicates(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	result := make([]int, 0, len(arr))
	result = append(result, arr[0])

	for i := 1; i < len(arr); i++ {
		if arr[i] != result[len(result)-1] {
			result = append(result, arr[i])
		}
	}

	return result
}
