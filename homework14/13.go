package main

import (
	"bufio"
	"os"
	"strings"
)

func wordFrequency(fileName string) (map[string]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	wordFreq := make(map[string]int)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		words := strings.Split(line, " ")
		for _, word := range words {
			word = strings.TrimSpace(word)
			if word != "" {
				wordFreq[word]++
			}
		}
	}

	return wordFreq, nil
}
