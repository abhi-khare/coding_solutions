package main

import "fmt"

func checkSorted(arr []int) bool {
	numPivots := 0
	for iter := 0; iter <= len(arr)-1; iter++ {
		if arr[(iter+1)%len(arr)] >= arr[iter] {
			continue
		} else {
			numPivots += 1
			if numPivots == 2 {
				return false
			}
		}
	}
	return true
}

func main() {

	var arr = []int{4, 1, 2, 3, 5}

	status := checkSorted(arr)

	fmt.Println(status)
}
