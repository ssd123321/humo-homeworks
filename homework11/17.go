package main

func DeleteDuplicates(s1 string) string {
	var NewStr string
	for i, v := range s1 {
		for j, v2 := range s1 {
			if string(v2) == string(v) && i != j {
				continue
			} else {
				NewStr
			}
		}
	}
}
