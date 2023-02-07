package main

import "fmt"

func minimumElement(arr []int) int {

	if arr[0] < arr[len(arr)-1]{
		return arr[0]
	} else {
		start, end := 0, len(arr)-1

		for start < end {
			mid := (start + end) / 2
			if arr[mid] < arr[mid-1] {
				return arr[mid]
			}  else if arr[mid] < arr[0] && arr[mid] < arr[len(arr)-1] {
				
			} else if arr[mid] < arr[0] && arr[mid] < arr[len(arr)-1] {
			
			} 
		}

		return arr[0]
	}
}

func main() {

	var arr = []int{3,4,5,1,2}

	minEle := minimumElement(arr)

	fmt.Println(minEle)
}
