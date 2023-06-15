package main

import "fmt"

func getProductArrays(nums []int) ([]int, []int){
	prefixProduct, suffixProduct := make([]int, len(nums)), make([]int, len(nums))
	
	leftProd, rightProd := 1,1

	for iter:=0; iter<len(nums); iter++{
		leftProd *= nums[iter]
		rightProd *= nums[len(nums)-iter-1]

		prefixProduct[iter] = leftProd
		suffixProduct[len(nums)-iter-1] = rightProd
	}

	return prefixProduct, suffixProduct

}

func productExceptSelf(nums []int) []int {
	
	prefixProduct, suffixProduct := getProductArrays(nums)

	for idx, _ := range nums{
		if idx ==0{
			nums[idx] = suffixProduct[idx+1]
		} else if idx == len(nums)-1{
			nums[idx] = prefixProduct[idx-1]
		} else{
			nums[idx] = prefixProduct[idx-1] * suffixProduct[idx+1]
		}
	}
	return nums
}

func main() {

	nums := []int{-1}

	ans := productExceptSelf(nums)


	fmt.Println(ans)
}