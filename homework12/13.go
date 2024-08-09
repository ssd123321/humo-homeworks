package main

func UniteMaps(m1 map[int]string, m2 map[int]string) map[int]string {
	final := make(map[int]string)
	for key := range m1 {
		for key2, value2 := range m2 {
			if key2 == key {
				m1[key] = value2
			}
		}
		final[key] = m1[key]
	}
	return final
}
