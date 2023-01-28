package main

import "fmt"

func getPossiblePairs(num int) int {

	if num <=1{
		return 1
	} else {
		return getPossiblePairs(num-1) + (num-1)*getPossiblePairs(num-2)
	}
}
func main() {

	var num int

	fmt.Scan(&num)

	pairCount := getPossiblePairs(num)

	fmt.Print(pairCount)
}