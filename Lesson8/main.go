package main

import "fmt"

func main() {
	/*
		var array [5]int
		array[0] = 1
		array[1] = 2
		array2 := [...]string{"1", "2", "3"}
		fmt.Println(array, array2)
		var arr [5]int
		for i := 0; i < 5; i++ {
			arr[i] = i
		}
		fmt.Println(arr)
		fmt.Println(len(arr), cap(arr))
		for i, v := range arr {
			fmt.Println(i, v)
		}
		slice := []int{1, 2, 3, 4, 5}
		fmt.Println(slice)
		fmt.Println(cap(slice))
		slice = arr[1:4]
		fmt.Println(slice)
		slice2 := make([]int, 10, 20)
		slice2 = append(slice2, 10)
		fmt.Println(len(slice2), cap(slice2))
		s1 := []int{1, 2, 3}
		s2 := []int{4, 5}
		s3 := append(s2, s1...)
		fmt.Println(s3)
		fmt.Println(len(s3), cap(s3))
		src := []int{1, 2, 3}
		dst := make([]int, len(src))
		number := copy(dst, src)
		fmt.Println(dst, number)
	*/
	// Задачи
	var arr1 [7]int
	var value int = 1
	for i := 0; i < len(arr1); i++ {
		arr1[i] = value
		value++
	}
	fmt.Println(arr1)
	// 2
	arr2 := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(arr2)
	// 3
	var arr3 = [4]bool{true, false, true, false}
	fmt.Println(arr3)
	// 4
	var arr4 [10]int
	fmt.Println(arr4)
	// 5
	arr5 := [4]bool{true, true, false, false}
	fmt.Println(arr5[1], arr5[3])
	// 6
	var arr6 = [3]bool{true, false, true}
	arr6[0] = false
	fmt.Println(arr6)
	//
	var arr7 = [4]string{"Hello", "How are you", "What are you doing", "What's happened"}
	for i, value := range arr7 {
		fmt.Println(i, value)
	}
	//
	a := [5]int{1, 2, 3, 4, 5}
	slice := a[1:3:4]
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))
	//

}
