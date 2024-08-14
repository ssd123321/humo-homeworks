package main

import (
	"io"
	"os"
	"strings"
)

func countUniqueWords(fileName string) (int, error) {
	var str string
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	buf := make([]byte, 0)
	for {
		_, err = file.Read(buf)
		if err == io.EOF {
			break
		}else if err != nil{
			return 0, err
		}
		str += string(buf)
	}
	m := make(map[string]int)
	words := strings.Fields(str)
	for _, word := range words {
		m[word]++
	}
	return len(m), nil
}
