package main

func subarraySum(nums []int, target int) [][]int {
	n := len(nums)
	result := [][]int{}

	for i := 0; i < n; i++ {
		sum := 0
		current := []int{}
		for j := i; j < n; j++ {
			sum += nums[j]
			current = append(current, nums[j])
			if sum == target {
				result = append(result, current)
			} else if sum > target {
				break
			}
		}
	}

	return result
}
