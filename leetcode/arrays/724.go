package main

import "fmt"

func getPrefixSum(arr []int) ([]int, []int){

	leftSum, rightSum := make([]int, len(arr)), make([]int , len(arr))

	// computing left prefix sum
	sum1, sum2 := 0, 0
	for idx, ele := range arr{
		sum1 += ele
		sum2 += arr[len(arr)-idx-1]
		leftSum[idx] = sum1
		rightSum[len(arr)-1-idx] = sum2
	}
	
	return leftSum, rightSum

}

func pivotIndex(nums []int) int {

    
	leftPrefixSum, rightPrefixSum := getPrefixSum(nums)

	pivot := 0
	for pivot < len(nums){
		leftSum := leftPrefixSum[pivot] - nums[pivot]
		rightSum := rightPrefixSum[pivot] - nums[pivot]

		if leftSum == rightSum{
			return pivot
		}

		pivot +=1
	}

	return -1
}



func main(){

	nums := []int{0,2}

	pivot := pivotIndex(nums)
	fmt.Println(pivot)
}