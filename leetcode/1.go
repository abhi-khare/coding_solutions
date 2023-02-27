package main

import (
	"sort"
)

type element struct {
	val int
	idx int
}

func bSearch(arr []element, start int, end int, target int) int {

	for start <= end {
		mid := (start + end) / 2

		if arr[mid].val == target {
			return mid
		} else if arr[mid].val < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1
}

func main() {

	nums := []int{2, 11, 7, 15}
	tempNums := make([]element, len(nums))
	target := 9

	for iter, ele := range nums {
		tempNums[iter].val = ele
		tempNums[iter].idx = iter
	}

	sort.Slice(tempNums, func(i, j int) bool {
		return tempNums[i].val < tempNums[j].val
	})

	idx, jdx := 0, 0

	for iter := 0; iter < len(tempNums)-1; iter++ {

		jdx = bSearch(tempNums, iter+1, len(nums)-1, target-tempNums[iter].val)

		if jdx != -1 {
			idx = iter
			break
		}
	}

	if idx != -1 && jdx != -1 {
		print(tempNums[idx].idx, tempNums[jdx].idx)
	}
}
