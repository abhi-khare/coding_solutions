package main

import (
	"fmt"
)

func abs(num int) int {
	if num < 0 {
		return -1 * num
	}
	return num
}

func SquareAndSort(arr []int) []int {

	ans := make([]int, len(arr))

	start, end, idx := 0, len(arr)-1, len(arr)-1

	for start <= end {
		if abs(arr[start]) > abs(arr[end]) {
			ans[idx] = arr[start] * arr[start]
			idx -= 1
			start += 1
		} else {
			ans[idx] = arr[end] * arr[end]
			idx -= 1
			end -= 1
		}
	}

	return ans
}

func main() {

	var arr = []int{2}

	ans := SquareAndSort(arr)

	fmt.Println(ans)

}
