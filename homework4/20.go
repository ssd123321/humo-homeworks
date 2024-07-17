package main

import "fmt"

func CreateGreeter(str string) func(name string) string {
	return func(name string) string {
		if str == "" {
			return fmt.Sprintf("Hello, %s!", name)
		}
		return fmt.Sprintf("%s, %s!", str, name)
	}
}
