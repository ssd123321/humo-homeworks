package main

import "strings"

func ChangeStr(s1, sub string, new string) string {
	return strings.ReplaceAll(s1, sub, new)
}
