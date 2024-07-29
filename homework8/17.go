package main

import "sort"

func intersection(nums1, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	var result []int
	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			result = append(result, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return result
}
