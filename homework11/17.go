package main

func DeleteDuplicates(s1 string) string {
	m := make(map[rune]bool)
	var r string
	for _, l := range s1 {
		if _, ok := m[l]; !ok {
			m[l] = true
			r += string(l)
		}
	}
	return r
}
