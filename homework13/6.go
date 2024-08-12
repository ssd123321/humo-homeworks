package main

type SubstringExtractor interface {
	Substr(start, length int) string
	Suffix(length int) string
}

type TextExtractor struct {
	text string
}

func (e *TextExtractor) Substr(start, length int) string {
	if start >= len(e.text) {
		return ""
	}
	end := start + length
	if end > len(e.text) {
		end = len(e.text)
	}
	return e.text[start:end]
}

func (e *TextExtractor) Suffix(length int) string {
	if length > len(e.text) {
		return e.text
	}
	return e.text[len(e.text)-length:]
}
