package main

import "sort"

func binarySearch(arr []int, start int, end int, target int) int {

	for start <= end {
		mid := (start + end) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1
}

func main() {

	nums := []int{3, 3}
	tempNums := make([]int, len(nums))
	target := 6

	for iter, ele := range nums {
		tempNums[iter] = ele
	}

	sort.Slice(tempNums, func(i, j int) bool {
		return tempNums[i] < tempNums[j]
	})

	idx, jdx := 0, 0

	for iter := 0; iter < len(tempNums)-1; iter++ {

		jdx = binarySearch(tempNums, iter+1, len(tempNums)-1, target-tempNums[iter])

		if jdx != -1 {
			idx = iter
			break
		}
	}

	print(idx, jdx)

	ans1, ans2 := 0, 0

	for iter, ele := range nums {
		if ele == tempNums[idx] {
			ans1 = iter
		}
		if ele == tempNums[jdx] {
			ans2 = iter
		}
	}
	print(ans1, ans2)
}
