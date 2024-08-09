package main

func InvertMap(m1 map[int]string) map[string]int {
	InvertedMap := make(map[string]int)
	for key, value := range m1 {
		InvertedMap[value] = key
	}
	return InvertedMap
}
