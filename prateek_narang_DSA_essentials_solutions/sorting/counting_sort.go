package main

import "fmt"

func countingSort(arr []int) []int {

	max_ele := 0

	for _, ele := range arr {
		if ele > max_ele {
			max_ele = ele
		}
	}

	freqTable := make([]int, max_ele+2)

	for _, ele := range arr {
		freqTable[ele] += 1
	}

	idx := 0
	for val, freq := range freqTable {

		if freq != 0 {
			for jter := 1; jter <= freq; jter++ {
				arr[idx] = val
				idx += 1
			}
		}
	}

	return arr

}

func main() {

	var arr = []int{2, 2, 2, 3, 1, 1, 2}

	arr = countingSort(arr)

	fmt.Println(arr)
}
