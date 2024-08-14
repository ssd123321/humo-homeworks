package main

import "log"

func main() {
	err := readAndWriteToFile("fs.txt")
	if err != nil {
		log.Fatal(err)
	}
}
