/* func aVeryBigSum(ar []int64, pointOfSum *int64) {
	// Write your code here
	var arLength int
	arLength = len(ar)
	for i := 0; i < arLength; i++ {
		*pointOfSum += ar[i]
	}
} */

package main

import "fmt"

func aVeryBigSum(ar []int64) int64 {
	// Write your code here
	var arLength int
	arLength = len(ar)
	var pointOfSumInScope int64
	for i := 0; i < arLength; i++ {
		pointOfSumInScope += ar[i]
	}
	return pointOfSumInScope
}

func main() {
	var ar []int64
	ar = []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	var result int64
	result = aVeryBigSum(ar)
	fmt.Println(result)
}
