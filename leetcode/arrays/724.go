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
    
	leftSum, rightSum := getPrefixSum(nums)

	for iter:=0; iter<len(nums); iter++{
		
	}
}


func main(){

	nums := []int{1,7}

	idx := pivotIndex(nums)
	a,b := getPrefixSum(nums)
	fmt.Println(idx,a,b)
}