package main

func maxSubarraySumWithTwoElements(nums []int) int {
	maxSum := nums[0]
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				sum += nums[j]
				if sum > maxSum {
					maxSum = sum
				}
			}
		}
	}
	return maxSum
}
