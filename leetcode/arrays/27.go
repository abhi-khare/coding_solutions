package main

import "fmt"

func removeElement(nums []int, val int) int {
	idx := -1

	for iter:=0; iter<len(nums); iter++{
		if nums[iter] != val{
			nums[idx+1] = nums[iter]
			idx +=1
		}
	}
	return idx +1
}

func main() {

	nums := []int{3,2,2,3}
	val := 3

	ans := removeElement(nums, val)

	fmt.Println(ans)
}
