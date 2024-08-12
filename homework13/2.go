package main

import (
	"fmt"
	"strings"
)

type Formatter interface {
	ToUpper() string
	ToLower() string
}

type MyFormatter struct {
	text string
}

func (f *MyFormatter) ToUpper() string {
	return strings.ToUpper(f.text)
}

func (f *MyFormatter) ToLower() string {
	return strings.ToLower(f.text)
}

func formatString(f Formatter) {
	fmt.Println("Uppercase:", f.ToUpper())
	fmt.Println("Lowercase:", f.ToLower())
}
