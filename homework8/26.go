package main

import "math"

func maximumProduct(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	min1, min2 := math.MaxInt32, math.MaxInt32
	max1, max2, max3 := math.MinInt32, math.MinInt32, math.MinInt32

	for _, num := range nums {
		if num < min1 {
			min2 = min1
			min1 = num
		} else if num < min2 {
			min2 = num
		}

		if num > max1 {
			max3 = max2
			max2 = max1
			max1 = num
		} else if num > max2 {
			max3 = max2
			max2 = num
		} else if num > max3 {
			max3 = num
		}
	}

	return max(max1*max2*max3, min1*min2*max1)
}
