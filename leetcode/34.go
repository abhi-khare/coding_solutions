package main

import "fmt"


func getStrictLowerBound(arr []int, key int) int{

	if arr[0] == key{
		return 0
	}

	start, end := 0, len(arr)-1

	for start <= end {
		mid := (start+end)/2

		if arr[mid] == key && arr[mid-1] < key{
			return mid
		} else if key <= arr[mid] {
			end = mid -1
		} else {
			start = mid + 1
		}
	}

	return -1
}


func getStrictUpperBound(arr []int, key int) int{
	if arr[len(arr)-1] == key{
		return len(arr)-1
	}

	start, end := 0, len(arr)-1

	for start <= end {
		mid := (start+end)/2

		if arr[mid] == key && arr[mid+1] > key{
			return mid
		} else if key < arr[mid] {
			end = mid -1
		} else {
			start = mid + 1
		}
	}

	return -1
}

func getBoundIndices(arr []int, key int) []int{

	if len(arr)> 0 && key >= arr[0] && key <= arr[len(arr)-1]{
		lowIdx := getStrictLowerBound(arr, key)

		if lowIdx == -1{
			return []int{-1, -1}
		}
		highIdx := getStrictUpperBound(arr, key)

		return []int{lowIdx, highIdx}
	} else{
		return []int{-1,-1}
	}
	
}

func main() {

	var arr = []int{5}
	var key int = 5

	bound := getBoundIndices(arr, key)

	fmt.Println(bound)

}