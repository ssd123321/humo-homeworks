package main

import "strconv"

func CountNumberInString(s string) int {
	var n int
	for _, v := range s {
		number, err := strconv.Atoi(string(v))
		if err != nil {
			continue
		}
		n += number
	}
	return n
}
