package main
import (
	"bufio"
	"os"
	"strings"
)

func removeEmptyLines(fileName, outputFile string) error {
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
		if strings.TrimSpace(line) != "" {
			_, err = writer.WriteString(line)
			if err != nil {
				return err
			}
		}
	}

	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
}