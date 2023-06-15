package main

import "fmt"

func findDifference(nums1 []int, nums2 []int) [][]int {

	ans1 := []int{}
	ans2 := []int{}
	hashMap1 := make(map[int]int)
	hashMap2 := make(map[int]int)
	
	for _, ele := range nums1{
		hashMap1[ele] = 0
	}

	for _, ele := range nums2{
		hashMap2[ele] = 0
	}

	for ele,_ := range hashMap1{
		if _, ok := hashMap2[ele]; !ok{
			ans1 = append(ans1, ele)
		}
	}

	
	for ele, _ := range hashMap2{
		if _, ok := hashMap1[ele]; !ok{
			ans2 = append(ans2, ele)
		}
	}

	return [][]int{ans1, ans2}
	

}

func main() {

	nums1 := []int{1,2,3,3}
	nums2 := []int{1, 1,2,2}

	ans := findDifference(nums1, nums2)

	fmt.Println(ans)

}