package main

import "fmt"

func selectionSort(arr []int, reverse bool) {

	for iter := 0; iter < len(arr)-1; iter++ {

		idx := iter

		for jter := iter + 1; jter < len(arr); jter++ {

			if reverse == true {
				if arr[jter] < arr[iter] {
					idx = jter
				}
			} else {
				if arr[jter] > arr[iter] {
					idx = jter
				}
			}

		}

		temp := arr[iter]
		arr[iter] = arr[idx]
		arr[idx] = temp
	}
}

func main() {

	var arr = []int{1, 2}

	selectionSort(arr, true)

	fmt.Println(arr)
}
