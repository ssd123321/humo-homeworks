package main

import (
	"fmt"
	"strings"
)

type StringCleaner interface {
	TrimSpaces() string
	RemoveSpaces() string
}

type TextCleaner struct {
	text string
}

func (c *TextCleaner) TrimSpaces() string {
	return strings.TrimSpace(c.text)
}

func (c *TextCleaner) RemoveSpaces() string {
	return strings.ReplaceAll(c.text, " ", "")
}

func cleanString(c StringCleaner) {
	fmt.Println("Trimmed Spaces:", c.TrimSpaces())
	fmt.Println("Removed Spaces:", c.RemoveSpaces())
}
