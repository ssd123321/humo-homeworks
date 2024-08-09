package main

import "fmt"

func SumAllValues(m map[int]string) string {
	var str string
	for _, value := range m {
		str += fmt.Sprintf("%s,", value)
	}
	return string([]rune(str[:len(str)-1]))
}
