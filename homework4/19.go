package main

func CreateMultiplier(n int) func() int {
	return func() int {
		if n < 0 {
			return n * (-1 * n)
		}
		return n * n
	}
}
