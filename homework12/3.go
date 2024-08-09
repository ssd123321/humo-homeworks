package main
func UpgradeValueByKey(m map[string]int, key string, value int)bool{
	if _, ok := m[key];!ok{
		return false
	}
	m[key] = value
	return true
}