package main

func GetSliceOfKeys(m map[string]int) []string {
	s := make([]string, 0)
	for key := range m {
		s = append(s, key)
	}
	return s
}
