package main

import "fmt"

func findPivot(arr []int) int {

	if len(arr) == 1 || arr[0] < arr[len(arr)-1]{
		return 0
	}

	start, end := 0, len(arr)-1

	for start <= end{
		mid := (start + end)/2

		if mid > 0 && arr[mid] < arr[mid-1]{
			return mid
		} else if arr[mid] < arr[len(arr)-1]{
			end = mid -1
		} else{
			start = mid +1
		}
	}

	return -1
}

func binarySearchMod( arr []int, target int, pivot int) int {

	start, end := 0, len(arr)-1

	for start <= end {
		mid := ((start+end)/2 + pivot )%len(arr)

		if arr[mid] == target{
			return mid
		} else if arr[mid] < target{
			start = (mid+1)
		} else {
			end = (mid-1)
		}
	}

	return -1
}

func binarSearchRotated(arr []int, target int) int {

	pivotIdx := findPivot(arr)

	idx := binarySearchMod(arr, target, pivotIdx)
	
	return idx

}

func main() {

	var arr = []int{4,5,6,7,0,1,2}

	var target int = 7

	idx := binarSearchRotated(arr, target)

	fmt.Println(idx)
}