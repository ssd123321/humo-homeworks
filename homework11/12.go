package main

func ReverseStr(s1 string) string {
	var reversed string
	r := []rune(s1)
	for i := len(r) - 1; i >= 0; i-- {
		reversed += string(r[i])
	}
	return reversed
}
