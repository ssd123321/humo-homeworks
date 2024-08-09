package main
func DeleteValueByKey(key string, m map[string]int)bool{
	if _, ok := m[key]; !ok{
		return false
	}
	delete(m, key)
	return true
}