package main

import "fmt"

func rotateArray(arr []int, shift int) []int {

	ans := make([]int, len(arr))

	for iter, ele := range arr {
		ans[(iter+shift)%len(arr)] = ele
	}

	return ans
}

func main() {

	var arr = []int{1, 3, 2}
	var k int = 0

	ans := rotateArray(arr, k)

	fmt.Println(ans)
}
