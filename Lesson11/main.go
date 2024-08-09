package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	x := "Hello, World"
	fmt.Println(x[0])
	y := string(x[0])
	fmt.Println(y)
	z := x[7:12] // Автоматически превращает в string
	fmt.Println(z)
	fmt.Println(strings.Contains(x, " "))
	fmt.Println(strings.Index(x, ", W"))
	// Байты и руны
	str := "Hello, Мир"
	fmt.Println([]byte(str))
	fmt.Println(string([]byte(str)))
	for _, v := range str {
		fmt.Println(v)
		fmt.Println(string(v))
	}
	t := "пeokeskesopcсусущс"
	// x1, x2 := utf8.DecodeRuneInString(t)
	x3, x4 := utf8.DecodeRune([]byte(t))
	fmt.Println(x3, x4)
	// strings
	// strconv
	// unicode
	// utf8
	var str1 string = "Hello, World!"
	str2 := "Hello, Go!"
	str3 := `This is a
	multiple string`
	fmt.Println(str1, str2, str3)
	s1 := "Hello, "
	s2 := "World!"
	s3 := s1 + s2
	fmt.Println(s3)
	s := "Hello, world!"
	c := s[1]
	fmt.Println(c)
	sub := s[7:12]
	fmt.Println(sub)
	s1byte := []byte{101, 103}
	s2byte := []byte{101, 103}
	s1 = string(s1byte)
	s2 = string(s2byte)
	fmt.Println(s1, s2)
	fmt.Println(s1 == s2)
	// Поиск подстроки в строке
	s = "Hello, World!"
	fmt.Println(strings.Contains(s, "World"))
	fmt.Println(strings.Index(s, "World"))
	// Руны и Байты в Go
	s = "Hello, 世界"
	bytes := []byte(s)
	fmt.Println(bytes)
	runes := []rune(s)
	fmt.Println(runes)
	// ASCII и UTF-8
	s = "Hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
}
