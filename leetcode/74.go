package main

import (
	"fmt"
)

func BinarySearch2D(arr [][]int, target int) bool {

	if target < arr[0][0] || target > arr[len(arr)-1][len(arr[0])-1] {
		return false
	} else {
		start, end := 0, len(arr)*len(arr[0])-1

		for start <= end {
			mid := (start + end) / 2

			yidx := mid / len(arr[0])
			xidx := mid - yidx*len(arr[0])

			if arr[yidx][xidx] == target {
				return true
			} else if arr[yidx][xidx] > target {
				end = mid - 1
			} else {
				start = mid + 1
			}

		}

		return false
	}
}

func main() {

	var arr = [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	var target int = 8

	idx := BinarySearch2D(arr, target)

	fmt.Println(idx)
}
