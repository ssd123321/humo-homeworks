package main

import (
	"fmt"
	"strings"
)

func FormatAndPrint(s string, formatFunc func(string) string) {
	if s == "" {
		fmt.Println("Пустая строка")
	} else {
		formattedString := formatFunc(strings.ToUpper(s))
		fmt.Println(formattedString)
	}
}
func formatFunc(s string) string {
	return "Отформатированная строка " + s
}
