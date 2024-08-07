package main

func generateSubstrings(input string) []string {
	var substrings []string
	for i := 0; i < len(input); i++ {
		for j := i + 1; j <= len(input); j++ {
			substrings = append(substrings, input[i:j])
		}
	}
	return substrings
}
