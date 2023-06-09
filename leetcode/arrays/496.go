package main

import "fmt"

func max(a int, b int) int{
	if a >b{
		return a
	}

	return b
}
func getPreProcessedArray(arr []int) []int{
	preProcessArray := make([]int, len(arr))
	preProcessArray[len(arr)-1] = -1
	var nextGreatest int = -1

	for iter:= len(arr)-2; iter>=0; iter--{

		if arr[iter+1] > arr[iter]{
			preProcessArray[iter] = arr[iter+1]
		} else if nextGreatest > arr[iter]{
			preProcessArray[iter] = preProcessArray[]
		}
		greaterEle := max(nextGreatest, arr[iter+1])
		if greaterEle > arr[iter]{
			preProcessArray[iter] = greaterEle
			nextGreatest = greaterEle
		} else{
			preProcessArray[iter] = -1
		}
	}

	return preProcessArray
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	hashMap := make(map[int]int)

	preProcessArray := getPreProcessedArray(nums2)

	for idx, ele := range nums2{
		hashMap[ele] = preProcessArray[idx]
	}

	for idx, val := range nums1{
		nums1[idx] = hashMap[val]
	}

	return nums1

}

func main() {

	//nums1 := []int{4, 1, 2}
	nums2 := []int{5,3,7,8, 7}

	ans := getPreProcessedArray( nums2)

	fmt.Println(ans)

}