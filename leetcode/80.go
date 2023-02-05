package main

import "fmt"

func removeDuplicatesTwice(arr []int) int {

	idx := 0
	for iter := 0; iter < len(arr); {
		jter := iter
		for ; jter < len(arr); jter++ {

			if arr[jter] != arr[iter] {
				break
			}

			if jter-iter <= 1 {
				arr[idx] = arr[jter]
				idx += 1
			}

		}
		iter = jter
	}

	return idx
}

func main() {

	var nums = []int{0, 1, 2}

	count := removeDuplicatesTwice(nums)

	fmt.Println(count, nums)
}
