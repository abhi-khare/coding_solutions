package main

import "fmt"

func largestElement(arr []int) int {

	var maxEle int = 0

	for _, ele := range arr{
		if ele > maxEle{
			maxEle = ele
		}
	}

	return maxEle
}

func main() {

	var arr = []int{-3, 4, 12, 3, 9, 10}

	ele := largestElement(arr)

	fmt.Println(ele)
}