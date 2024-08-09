package main

func GetKey(given int, m map[string]int) string {
	for key, value := range m {
		if value == given {
			return key
		}
	}
	return ""
}
