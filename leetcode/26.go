package main

import "fmt"

func removeDuplicates(arr []int) int {
	 
	idx := 0

	for iter:=1; iter<len(arr); iter++{
		if arr[idx] != arr[iter]{
			idx +=1
			arr[idx] = arr[iter] 
		}
	}

	return idx + 1
}

func main() {

	var arr = []int{-23,0,1}

	k := removeDuplicates(arr)

	fmt.Print(k, arr)
}