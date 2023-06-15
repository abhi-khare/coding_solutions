package main

import "fmt"

type NumArray struct {
	arr []int
}

func Constructor(nums []int) NumArray {
	var prefixSum NumArray
	sum := 0
	for _, ele := range nums {
		sum += ele
		prefixSum.arr = append(prefixSum.arr, sum )
	}

	return prefixSum
}

func (this *NumArray) SumRange(left int, right int) int {

	if left == 0{
		return this.arr[right]
	} else{
		return this.arr[right] - this.arr[left-1]
	}

	
    
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

func main(){

	temp := Constructor([]int{-2,2})

	param_1 := temp.SumRange(1,1)

	fmt.Println(param_1)

}
