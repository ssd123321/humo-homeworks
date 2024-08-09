package main

func UpgradeValue(m map[int]string, f func(string) string) map[int]string {
	for key := range m {
		m[key] = f(m[key])
	}
	return m
}
