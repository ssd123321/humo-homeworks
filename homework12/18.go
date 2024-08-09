package main

func CopyMap(m map[int]string) map[int]string {
	m2 := make(map[int]string)
	for key, value := range m {
		m2[key] = value
	}
	return m2
}
