package main

import "sort"

func isSubset(nums1, nums2 []int) bool {
	sort.Ints(nums1)
	sort.Ints(nums2)

	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			i++
			j++
		} else if nums1[i] < nums2[j] {
			return false
		} else {
			j++
		}
	}
	return i == len(nums1)
}
