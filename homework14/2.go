package main

import (
	"bufio"
	"os"
)

func countLines(fileName string) (int, error) {
	x, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	s := bufio.NewScanner(x)
	s.Split(bufio.ScanRunes)
	var i int
	for s.Scan() {
		i++
	}
	return i, nil
}
