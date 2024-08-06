package main

func RemoveLetters(s1 string, r rune) string {
	var NewStr string
	for _, letter := range s1 {
		if letter != r {
			NewStr += string(letter)
		}
	}
	return NewStr
}
