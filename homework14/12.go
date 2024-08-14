package main

import (
	"bufio"
	"os"
	"strings"
)

func replaceWordInFile(fileName, oldWord, newWord string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	tmpFile, err := os.Create(fileName + ".tmp")
	if err != nil {
		return err
	}
	defer tmpFile.Close()

	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(tmpFile)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		newLine := strings.ReplaceAll(line, oldWord, newWord)
		_, err = writer.WriteString(newLine)
		if err != nil {
			return err
		}
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	err = tmpFile.Close()
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
	err = os.Rename(fileName+".tmp", fileName)
	if err != nil {
		return err
	}

	return nil
}
