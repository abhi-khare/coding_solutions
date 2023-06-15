package main

import "fmt"


func getPreProcessedArray(arr []int) map[int]int{
	stack := []int{1000000}
	hashMap := make(map[int]int)

	for _ , ele := range arr{

		topEle := stack[len(stack)-1]
		if ele < topEle{
			stack = append(stack, ele)
		} else{
			for topEle < ele{
				hashMap[topEle] = ele
				stack = stack[:len(stack)-1]
				topEle = stack[len(stack)-1]
			}
			stack = append(stack, ele)
		}
	}

	for _, ele := range stack{
		hashMap[ele] = -1
	}

	return hashMap

}

func nextGreaterElement(nums1 []int, nums2 []int) []int {


	hashMap := getPreProcessedArray(nums2)


	for idx, val := range nums1{
		nums1[idx] = hashMap[val]
	}

	return nums1

}

func main() {

	nums1 := []int{3}
	nums2 := []int{3}

	ans := nextGreaterElement(nums1, nums2)

	fmt.Println(ans)

}