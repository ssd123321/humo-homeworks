package main

func RepeatStr(str string, n int) string {
	var NewStr string
	for i := 1; i <= n; i++ {
		NewStr += str
	}
	return NewStr
}
