package main

import (
	"bufio"
	"os"
)

func compareFiles(file1, file2 string) (bool, error) {
	f1, err := os.Open(file1)
	if err != nil {
		return false, err
	}
	defer f1.Close()

	f2, err := os.Open(file2)
	if err != nil {
		return false, err
	}
	defer f2.Close()

	r1 := bufio.NewReader(f1)
	r2 := bufio.NewReader(f2)

	for {
		b1, err1 := r1.ReadBytes('\n')
		b2, err2 := r2.ReadBytes('\n')

		if err1 != nil || err2 != nil {
			return err1 == err2, nil
		}

		if string(b1) != string(b2) {
			return false, nil
		}
	}
}
