package main 

import "fmt"

func missingNumber(nums []int) int {
    
	n := len(nums)
	expectedSum := (n+1)*n/2
	actualSum :=0

	for _, ele:= range nums{
		actualSum += ele
	}

	return expectedSum-actualSum
}

func main(){

	var arr = []int{1}

	missingEle := missingNumber(arr)

	fmt.Println(missingEle)
}