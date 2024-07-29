package main

func moveNegatives(nums []int) []int {
	n := len(nums)
	left, right := 0, n-1

	for left <= right {
		if nums[left] < 0 {
			left++
		} else {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		}
	}

	return nums
}
