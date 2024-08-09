package main
func CheckDuplicates(m map[string]int)bool{
	for key, value := range m{
		for key2, value2 := range m{
			if value2 == value && key != key2{
				return true
			}
		}
	}
	return true
}