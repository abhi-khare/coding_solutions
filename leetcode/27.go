package main

import "fmt"

func removeElement(arr []int, val int) int {
	 
	idx := 0

	for _, ele := range arr{
		if ele != val{
			arr[idx] = ele
			idx += 1
		}
	}

	return idx 
}

func main() {

	var arr = []int{-3}

	var val int = 3

	k := removeElement(arr, val)

	fmt.Print(k, arr)
}