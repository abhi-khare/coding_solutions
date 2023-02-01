package main

import "fmt"

func lowerBound(arr []int, key int) int {

	start, end := 0 , len(arr)-1

	for start <= end {
		mid := (start+end)/2

		if key == arr[mid]{
			return mid
		} else if arr[mid] > key{
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return end
}

func main() {

	var arr = []int{-2, 0, 3, 7 , 9}
	var key int = -3

	ele := lowerBound(arr, key)

	if ele == len(arr){
		fmt.Println(-1)
	} else{
		fmt.Println(arr[ele])
	}
	
}