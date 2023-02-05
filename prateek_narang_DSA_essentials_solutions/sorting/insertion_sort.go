package main

import "fmt"

func insertionSort(arr []int) {

	for iter := 1; iter < len(arr); iter++ {

		ele := arr[iter]
		jter := iter - 1
		for ; jter >= 0; jter-- {

			if arr[jter] > ele {
				arr[jter+1] = arr[jter]
			} else {
				break
			}
		}
		arr[jter+1] = ele
	}
}

func main() {

	var arr = []int{-1}

	insertionSort(arr)

	fmt.Println(arr)
}
