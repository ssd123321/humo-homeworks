package main

import "fmt"

func main() {
	for i := 0; i < 9; i++ {
		fmt.Println(i)
	}
	j := 0
	for ; j < 9; j++ {
		fmt.Println(j)
	}
	var i int
	for ; i <= 9; i++ {
		fmt.Println(i)
	}
	fmt.Println(i)
	num := 546
	for num > 0 {
		a := num % 10
		fmt.Println(a)
		num = num / 10
	}
	fmt.Println(num)
	for n := 546; n > 0; n /= 10 {
		a := n % 10
		fmt.Println(a)
	}
	for x := 0; x <= 100; x++ {
		if x%3 != 0 {
			fmt.Println(x)
		}
	}
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			break
		}
		fmt.Println(i)
	}
	for i := 2; i < 10; i++ {
		for j := 2; j < 10; j++ {
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
	}
	/*var N int
	var sum float64
	fmt.Scan(&N)
	for i := 1; i <= N; i++ {
		j := 1
		sum += float64(j) / float64(i)
	}
	fmt.Println(sum)*/
	var N int
	var sum float64 = 1
	fmt.Scan(&N)
	var s float64 = 1
	for ; int(s) <= N; s++ {
		var j float64 = 0.1
		sum *= s + j
		j += 0.1
	}
	fmt.Println(sum)
}
