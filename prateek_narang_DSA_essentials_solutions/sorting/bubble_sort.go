package main

import "fmt"

func bubbleSort(arr []int, reverse bool) {
	cnt := 0
	for iter := 0; iter < len(arr)-1; iter++ {

		sortRequired := false
		for jter := 0; jter < len(arr)-1; jter++ {

			if reverse == true {
				if arr[jter] > arr[jter+1] {
					sortRequired = true
					temp := arr[jter]
					arr[jter] = arr[jter+1]
					arr[jter+1] = temp
					cnt += 1
				}
			} else {
				if arr[jter] < arr[jter+1] {
					sortRequired = true
					temp := arr[jter]
					arr[jter] = arr[jter+1]
					arr[jter+1] = temp
					cnt += 1
				}
			}

		}
		if !sortRequired {
			break
		}
	}

	fmt.Println("Total Steps required: ", cnt)
}

func main() {

	var arr = []int{0, 1, 2, 3, 4}

	bubbleSort(arr, false)

	fmt.Println(arr)
}
