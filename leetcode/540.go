package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	start, end := 0, len(nums)-1

	for start <= end {
		mid := (start + end) / 2

		leftArraySize, rightArraySize := mid, len(nums)-1-mid

		if nums[mid-1] != nums[mid] && nums[mid+1] == nums[mid] {
			return mid
		} else if nums[mid-1] == nums[mid] && leftArraySize%2 == 0 {
			end = mid - 1
		} else if nums[mid-1] == nums[mid] && leftArraySize%2 != 0 {
			start = mid + 1
		} else if nums[mid-1] == nums[mid] && rightArraySize%2 == 0 {
			end = mid - 1
		} else if nums[mid-1] == nums[mid] && rightArraySize%2 != 0 {
			start = mid + 1
		}
	}

	return -1
}

func main() {

	var arr = []int{1, 1, 2, 3, 3, 4, 4, 8, 8}

	singleEle := singleNonDuplicate(arr)

	fmt.Println(singleEle)

}
