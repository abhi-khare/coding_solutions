package main 

import "fmt"

func moveZeroes(nums []int)  {
    idx:= 0
	for iter:=0; iter<len(nums);iter++{
		if nums[iter] != 0{
			// swapping elements at index : idx and iter
			temp := nums[iter]
			nums[iter] = nums[idx]
			nums[idx] = temp
			idx+=1
		}
	}
}

func main(){

	var arr = []int{0,0,0,3}

	moveZeroes(arr)

	fmt.Println(arr)
}