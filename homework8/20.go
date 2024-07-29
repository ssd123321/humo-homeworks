package main

func lengthOfLongestSubarray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	maxLength := 1
	left, right := 0, 0

	for right < n {
		for i := left; i < right; i++ {
			if nums[i] == nums[right] {
				left = i + 1
				break
			}
		}
		maxLength = max(maxLength, right-left+1)

		right++
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
