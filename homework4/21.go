package main

func CreateValidator(password string) func(str string) bool {
	return func(str string) bool {
		if password == str {
			return true
		}
		return false
	}
}
