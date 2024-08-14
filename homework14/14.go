package main

import (
	"bufio"
	"os"
)

func reverseLines(fileName, outputFile string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer output.Close()

	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(output)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		reversedLine := reverseString(line)
		_, err = writer.WriteString(reversedLine)
		if err != nil {
			return err
		}
	}

	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
