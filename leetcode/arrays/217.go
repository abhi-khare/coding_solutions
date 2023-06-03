package main

import "fmt"

func containsDuplicate(nums []int) bool {

	var freqMap = make(map[int]int)

	for _, ele := range nums {

		if _, ok := freqMap[ele]; ok{
			return true
		} else{
			freqMap[ele] = 1
		}
	}

	return false

}

func main() {

	nums := []int{1}

	ans := containsDuplicate(nums)

	fmt.Println(ans)
}