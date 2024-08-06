package main

import "strings"

func CheckEquality(s1, s2 string) bool {
	return strings.EqualFold(s1, s2)
}
