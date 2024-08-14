package main

import (
	"io"
	"os"
)

func countCharacters(fileName string) (int, error) {
	x, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	buf := make([]byte, 1000)
	var str string
	for {
		n, err := x.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil && err != io.EOF {
			return 0, err
		}
		str += string(buf[:n])
	}
	return len([]rune(str)), nil
}
