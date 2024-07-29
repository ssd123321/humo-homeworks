package main

import "fmt"

func main() {
	// 1
	x := make([]int, 3)
	x[0] = 10
	x[1] = 20
	x[2] = 3
	fmt.Println(maxElement(x))
	// 2
	fmt.Println(minElement(x))
	// 3
	fmt.Println(countPositive(x))
	// 4
	fmt.Println(sumArray(x))
	// 5
	fmt.Println(averageArray(x))
	// 6
	fmt.Println(removeNumber(x, 10))
	// 7
	fmt.Println(multiplyArray(x, 12))
	// 8
	fmt.Println(findIndices(x, 23))
	// 9
	fmt.Println(copyArray(x))
	// 10
	fmt.Println(combineArrays(x, []int{1, 2, 4}))
	// 11
	fmt.Println(swapMaxMin(x))
	// 12
	fmt.Println(isPalindrome(x))
	// 13
	fmt.Println(secondMax(x))
	// 14
	fmt.Println(reverseArray(x))
	// 15
	fmt.Println(removeDuplicates(x))
	// 16
	fmt.Println(moveZerosToEnd(x))
	// 17
	fmt.Println(intersection(x, []int{1, 4, 193}))
	// 18
	fmt.Println(isSubset(x, []int{1, 5, 3}))
	// 19
	fmt.Println(mergeSortedArrays(x, []int{538, 2382}))
	// 20
	fmt.Println(lengthOfLongestSubarray(x))
	// 21
	fmt.Println(subarraySum(x, 34))
	// 22
	fmt.Println(twoSum(x, 21))
	// 23
	fmt.Println(firstMissingPositive(x))
	// 24
	fmt.Println(maxSubarraySumWithTwoElements(x))
	// 25
	fmt.Println(maxSubarraySumEqualsK(x, 32))
	// 26
	fmt.Println(maximumProduct(x))
	// 27
	fmt.Println(maxSubarraySum(x))
	// 28
	fmt.Println(moveNegatives(x))
	// 29
	fmt.Println(maxLengthZeroSubarray(x))
	// 30
	fmt.Println(findGCD(x))
}
