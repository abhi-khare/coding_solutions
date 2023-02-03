package main

import "fmt"

func getLonelyElement(arr []int) int {

	sum := 0

	for _, ele := range arr {
		sum = (ele ^ sum)
	}
	return sum
}

func main() {

	var arr = []int{2, 2, 3, 3, 5, 5, 7}

	ele := getLonelyElement(arr)

	fmt.Println(ele)
}
