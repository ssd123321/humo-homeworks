package main

import (
	"bufio"
	"os"
)

func readAndWriteToFile(fileName string) error {
	file, err := os.OpenFile(fileName, os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	_, err = file.WriteString(scanner.Text())
	if err != nil {
		return err
	}
	return nil
}
