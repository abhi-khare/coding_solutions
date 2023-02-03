package main

import "fmt"

func binarySearch(arr []int, target int) int {

	start, end := 0, len(arr)-1

	for start <= end {
		mid := (start + end) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return start

}

func main() {

	var arr = []int{1, 4, 8, 11}

	var target int = 12

	idx := binarySearch(arr, target)

	fmt.Println(idx)
}
