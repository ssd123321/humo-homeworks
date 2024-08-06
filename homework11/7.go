package main

import "strings"

func UpAndDown(s1 string) (string, string) {
	return strings.ToUpper(s1), strings.ToLower(s1)
}
