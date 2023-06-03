package main

import "fmt"

func getConcatenation(nums []int) []int {

    ans_size := len(nums)
    ans := make([]int, 2*ans_size)

    for iter:=0; iter<ans_size; iter++{
        ans[iter] = nums[iter]
        ans[iter+ ans_size] = nums[iter]
    }

    return ans
}

func main() {

		nums := []int{1,2,3}

		ans := getConcatenation(nums)

		fmt.Println(ans)
}
