package main

func maxSubarraySumEqualsK(nums []int, k int) int {
	left, right, sum, maxLength := 0, 0, 0, 0

	for right < len(nums) {
		sum += nums[right]

		for sum > k {
			sum -= nums[left]
			left++
		}

		if sum == k {
			if right-left+1 > maxLength {
				maxLength = right - left + 1
			}
		}

		right++
	}

	return maxLength
}
