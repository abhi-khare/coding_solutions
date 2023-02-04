package main

import "fmt"

func minimumElement(arr []int) int {

	start, end := 0, len(arr)

	for start < end {
		mid := (start + end) / 2

		if arr[mid] < arr[(mid-1)%len(arr)] || arr[mid] > arr[(mid+1)%len(arr)] {
			return mid
		} else if arr[0] > arr[(mid)%len(arr)] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return 1
}

func main() {

	var arr = []int{}

	minEle := minimumElement(arr)

	fmt.Println(minEle)
}
