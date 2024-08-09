package main

func CheckKey(key string, m map[string]int) bool {
	_, ok := m[key]
	return ok
}
