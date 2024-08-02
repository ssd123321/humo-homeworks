package main

type Text struct {
	content []string
}

func (t *Text) AddText(text string) {
	t.content = append(t.content, text)
}
func (t *Text) ReplaceWord(oldWord, newWord string) {
	for i := 0; i < len(t.content); i++ {
		if t.content[i] == oldWord {
			t.content[i] = newWord
		}
	}
}
func (t Text) Length() int {
	return len(t.content)
}
