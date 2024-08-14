package main

import (
	"io"
	"os"
)

func concatenateFiles(file1, file2, outputFile string) error {
	buf1 := make([]byte, 1024)
	buf2 := make([]byte, 1014)
	var str1 string
	var str2 string
	f1, err := os.OpenFile(file1, os.O_RDONLY, 0755)
	if err != nil {
		return err
	}
	for {
		_, err := f1.Read(buf1)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		str1 += string(buf1)
	}
	f2, err := os.OpenFile(file2, os.O_RDONLY, 0755)
	if err != nil {
		return err
	}
	for {
		_, err := f2.Read(buf2)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		str2 += string(buf2)
	}
	f3, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	_, err = f3.WriteString(str1 + str2)
	if err != nil{
		return err
	}
	return nil
}
