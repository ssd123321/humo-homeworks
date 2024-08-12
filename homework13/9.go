package main

import (
	"fmt"
	"strings"
)

type PalindromeChecker interface {
	IsPalindrome() bool
	IsWordPalindrome() bool
}

type TextPalindromeChecker struct {
	text string
}

func (c *TextPalindromeChecker) IsPalindrome() bool {
	s := strings.ToLower(c.text)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func (c *TextPalindromeChecker) IsWordPalindrome() bool {
	words := strings.Split(c.text, " ")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		if words[i] != words[j] {
			return false
		}
	}
	return true
}

func checkPalindrome(c PalindromeChecker) {
	fmt.Println("Is Palindrome:", c.IsPalindrome())
	fmt.Println("Is Word Palindrome:", c.IsWordPalindrome())
}
