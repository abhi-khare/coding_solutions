package main

import "fmt"

func subarraySum(nums []int, target int) int {
	
	ans := 0

	// computing prefix sum
	prefixSum := make([]int, len(nums)+1)
	for iter:=1; iter<=len(nums); iter++{
		prefixSum[iter] = prefixSum[iter-1] + nums[iter-1]
	}

	fmt.Println(prefixSum)

	for iter:=0; iter<=len(nums)-1; iter++{
		for jter:=iter; jter<=len(nums)-1; jter++{

			// compute sum
			sum := prefixSum[jter+1] - prefixSum[iter]

			if sum == target{
				ans +=1
			}
		}
	}

	return ans
}

func main(){

	var arr = []int{1,2,3}
	var target int = 3

	count := subarraySum(arr, target)

	fmt.Println(count)
}