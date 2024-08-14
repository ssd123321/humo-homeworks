package main

import (
	"io"
	"os"
)

func reverseReadFile(fileName string) (string, error) {
	var final string
	var str string
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0755)
	if err != nil {
		return "", err
	}
	buf := make([]byte, 1024)
	for {
		_, err := file.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}
		str += string(buf)
	}
	s := []rune(str)
	for i := len(s) - 1; i >= 0; i-- {
		final += string(s[i])
	}
	return final, nil
}
