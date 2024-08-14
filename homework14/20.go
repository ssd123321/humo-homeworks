package main

import (
	"bufio"
	"fmt"
	"os"
)

func countCharactersPerLine(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	lineNumber := 1
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		fmt.Printf("Line %d: %d characters\n", lineNumber, len(line)-1)
		lineNumber++
	}

	return nil
}
