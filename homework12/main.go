package main

import "fmt"

func main() {
	// 1
	m := map[string]int{
		"Hello": len([]rune("Hello")),
		"How":   len([]rune("How")),
		"Are":   len([]rune("Are")),
	}
	for key := range m {
		fmt.Println(key, m[key])
	}
	// 2
	

}
