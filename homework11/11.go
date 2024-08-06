package main

import "strings"

func ChangeSub(s1, sub string, newsub string) string {
	return strings.Replace(s1, sub, newsub, 1)
}
