package main

import "strings"

func CheckPrefixAndSuffix(s1, prefix, suffix string) (bool, bool) {
	return strings.HasPrefix(s1, prefix), strings.HasSuffix(s1, suffix)
}
