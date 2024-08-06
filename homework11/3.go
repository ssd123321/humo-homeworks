package main

import "strings"
func CheckSubStr(s1 string, sub string)bool{
	return strings.Contains(s1, sub)
}