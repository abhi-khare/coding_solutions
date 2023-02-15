package main 

import "fmt"

func maxSubArraySum(nums []int) int {

	// checks for all negative edge case
	largest := nums[0]
	flag := 0
	for _, ele := range nums{
		if ele >0{
			flag = 1
		}
		if ele > largest{
			largest = ele
		}
	}

	if flag==0{
		return largest
	}

	maxSum := 0
	currentSum := 0

	for iter:=0; iter<len(nums); iter++{
		currentSum += nums[iter] 

		if currentSum > maxSum{
			maxSum = currentSum
		}

		if currentSum < 0{
			currentSum = 0
		}
	}
	return maxSum
}

func main(){

	var arr = []int{1}
	maxSum := maxSubArraySum(arr)

	fmt.Println(maxSum)

}