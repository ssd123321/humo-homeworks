package main

import (
	"fmt"
	"os"
)

func generateRepeatedLinesFile(fileName, line string, count int) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	for i := 0; i < count; i++ {
		_, err = fmt.Fprintln(file, line)
		if err != nil {
			return err
		}
	}

	return nil
}
