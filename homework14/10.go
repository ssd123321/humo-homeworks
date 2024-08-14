package main

import "os"

func fileExists(fileName string) bool {
	_, err := os.Open(fileName)
	if err != nil{
		return false
	}
	return true
}
