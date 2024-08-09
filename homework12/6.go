package main

import (
	"sort"
)

func SortKeys(map[int]string) []int {
	s := make([]int, 0)
	m := make(map[int]string)
	for key := range m {
		s = append(s, key)
	}
	sort.Ints(s)
	return s
}
