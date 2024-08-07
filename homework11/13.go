package main

func CountLetters(s1 string) map[string]int {
	m := make(map[string]int)
	for _, v := range s1 {
		m[string(v)] += 1
	}
	return m
}
