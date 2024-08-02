package main

type Note struct {
	content string
	tags    []string
}
type NoteBook struct {
	Notes []Note
}

func (n *NoteBook) AddNote(content string, tags []string) {
	n.Notes = append(n.Notes, Note{content, tags})
}
func (n NoteBook) NotesWithTag(tag string) []Note {
	notes := make([]Note, 0)
	for _, note := range n.Notes {
		for _, tag2 := range note.tags {
			if tag == tag2 {
				notes = append(notes, note)
				break
			}
		}
	}
	return notes
}
