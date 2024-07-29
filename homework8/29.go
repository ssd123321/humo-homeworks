package main

func maxLengthZeroSubarray(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	maxLength := 0
	sum := 0
	start, end := 0, 0

	for end < n {
		sum += nums[end]

		for sum == 0 {
			if end-start+1 > maxLength {
				maxLength = end - start + 1
			}
			sum -= nums[start]
			start++
		}

		end++
	}

	return maxLength
}
