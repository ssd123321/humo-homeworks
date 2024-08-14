package main

import (
	"bufio"
	"os"
)

func readFirstLine(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	b := bufio.NewScanner(file)
	b.Scan()
	return b.Text(), nil

}
