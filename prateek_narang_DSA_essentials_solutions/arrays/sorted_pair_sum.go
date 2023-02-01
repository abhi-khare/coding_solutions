package main

import (
	"fmt"
	"math"
)

func closestPairs(arr []int, val int) [2]int {

	start, end := 0, len(arr)-1

	ans := [2]int{start, end}
	var diff float64 = math.Abs(float64(arr[start]-arr[end]))

	for start <= end {
		sum := arr[start] + arr[end]

		if math.Abs(float64(sum-val)) < diff{
			diff = math.Abs(float64(sum-val))
			ans[0] = start
			ans[1] = end
		}

		if sum > val{
			end = end-1
		} else if start < val{
			start = start + 1
		} else {
			return ans
		}
	}

	return ans
}

func main() {

	var arr = []int{10, 22, 28, 29, 30, 40}
	var val int = 68

	idx := closestPairs(arr, val)

	fmt.Println(arr[idx[0]], " ", arr[idx[1]])
	
}