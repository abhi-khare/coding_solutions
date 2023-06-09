package main

import "fmt"

func abs(num int) int {
	if num <0{
		return -1*num
	} 
	return num
}

func findDisappearedNumbers(nums []int) []int {

	ans := []int{}
	for _, ele :=  range nums{
		val := abs(ele)
		if nums[val-1] > 0{
			nums[val-1] *= -1
		}
	}

	for idx, ele := range nums{
		if ele > 0{
			ans = append(ans, idx+1)
		}
	}
	
	return ans
}

func main() {

	nums := []int{1,1,3,4}

	ans := findDisappearedNumbers(nums)

	fmt.Println(ans)
}