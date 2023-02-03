package main

import (
	"fmt"
	"sort"
)

func getMajorityElement(arr []int) int {

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	return arr[len(arr)/2]
}

func mooreVoting(arr []int) int {

	majorEle := arr[0]
	freq := 1

	for _, ele := range arr {

		if ele == majorEle {
			freq += 1
		} else if ele != majorEle {

			freq -= 1

			if freq == 0 {
				majorEle = ele
				freq = 1
			}
		}
	}

	return majorEle
}

func main() {
	var arr = []int{3, 2, 3}

	majorityEle := mooreVoting(arr)

	fmt.Println(majorityEle)
}
