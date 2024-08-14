package main

import "os"

func writeStringToFile(fileName string, content string) error {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	file.WriteString(content)
	return nil
}
