package main

import "fmt"

func printMAxSum(arr []int) int {
	var maxSum int = -1000000
	var sumTillNow int = 0

	for iter := 0; iter < len(arr); iter++ {
		sumTillNow += arr[iter]

		if sumTillNow < 0 {
			sumTillNow = 0
		}

		if sumTillNow > maxSum {
			maxSum = sumTillNow
		}
	}

	return maxSum
}

func main() {

	var arr = []int{-3, -2, -2, -1, -2, -4}

	ans := printMAxSum(arr)

	fmt.Println(ans)
}