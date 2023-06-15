package main

import "fmt"

func checkPossibility(nums []int) bool {
	fixCount := 0

	if len(nums) <= 2{
		return true
	} 

	for iter:= len(nums)-1; iter>0; iter--{
		if nums[iter] < nums[iter-1]{
			fixCount +=1

			if fixCount == 2{
				return false
			}
			nums[iter-1] = nums[iter]
		}
	}

	return true
}


func main(){
	nums := []int{3,3,3,2,6}

	ans := checkPossibility(nums)
	fmt.Println(ans)
}