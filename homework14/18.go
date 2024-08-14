package main
import (
	"bufio"
	"os"
	"strings"
   )
   
   func countLinesWithWord(fileName, word string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
	 return 0, err
	}
	defer file.Close()
   
	reader := bufio.NewReader(file)
   
	count := 0
   
	for {
	 line, err := reader.ReadString('\n')
	 if err != nil {
	  break
	 }
   
	 if strings.Contains(line, word) {
	  count++
	 }
	}
   
	return count, nil
   }
   