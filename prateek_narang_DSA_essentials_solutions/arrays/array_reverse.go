package main

import "fmt"

func reverseArray(arr []int){

	start, end := 0, len(arr)-1

	for start < end {
		temp := arr[start]
		arr[start] = arr[end]
		arr[end] = temp 

		start += 1
		end -= 1
	}
}

func main() {

	var arr = []int{12, 3, 45, 6, 34, 56, 67}

	reverseArray(arr)

	fmt.Println(arr)
}