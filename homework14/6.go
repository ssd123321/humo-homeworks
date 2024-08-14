package main

import (
	"io"
	"os"
)

func copyFile(src, dst string) error {
	file1, err := os.OpenFile(src, os.O_RDONLY, 0755)
	if err != nil {
		return err
	}
	file2, err := os.OpenFile(dst, os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	_, err = io.Copy(file2, file1)
	if err != nil {
		return err
	}
	return nil

}
