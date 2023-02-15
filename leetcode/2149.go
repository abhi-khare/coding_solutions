package main 

import "fmt"

func rearrangeArray(nums []int) []int {
	ans := make([]int, len(nums))

	pos,neg := 0,1

	for iter:=0; iter<len(nums);iter++{

		if nums[iter] > 0{
			ans[pos] = nums[iter]
			pos +=2
			
		} else {
			ans[neg] = nums[iter]
			neg+=2
			
		}
	}
	return ans
}

func main(){

	var nums = []int{2,3,-1,-3}

	updatedNums := rearrangeArray(nums)

	fmt.Println(updatedNums)
}