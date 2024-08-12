package main

import (
	"fmt"
	"strings"
)

type StringConcatenator interface {
	Concat() string
	Join(separator string) string
}

type StringJoiner struct {
	strings []string
}

func (j *StringJoiner) Concat() string {
	return strings.Join(j.strings, "")
}

func (j *StringJoiner) Join(separator string) string {
	return strings.Join(j.strings, separator)
}

func concatenateStrings(c StringConcatenator) {
	fmt.Println("Concatenated:", c.Concat())
	fmt.Println("Joined with '-':", c.Join("-"))
}
