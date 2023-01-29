package main

import (
	"fmt"
)


func binarySearch(arr []int, key int) int {

	start, end := 0, len(arr)-1

	for start <= end {
		mid := (start+end)/2

		if arr[mid] == key{
			return mid
		} else if arr[mid] > key{
			end = mid - 1
		} else {
			start = mid  + 1
		}
	}

	return -1
}

func main() {

	var key, size_ int

	fmt.Println("Enter the size of the array!")
	fmt.Scan(&size_)

	var arr = make([]int, size_)

	fmt.Println("Enter the elements of the array in sorted sequence")
	for iter:=0; iter<size_; iter++{
		fmt.Scan(&arr[iter])
	}


	fmt.Println("Enter the key")
	fmt.Scan(&key)


	idx := binarySearch(arr, key)

	if idx == -1{
		fmt.Print("Key not found in the array")
	} else{
		fmt.Print("Element found at index: ", idx)
	}

}