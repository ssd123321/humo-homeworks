package main

import "strings"

func escapeHTML(s string) string {
	var b strings.Builder
	for _, ch := range s {
		switch ch {
		case '<':
			b.WriteString("&lt;")
		case '>':
			b.WriteString("&gt;")
		case '"':
			b.WriteString("&quot;")
		case '\'':
			b.WriteString("&#39;")
		case '&':
			b.WriteString("&amp;")
		default:
			b.WriteRune(ch)
		}
	}
	return b.String()
}
