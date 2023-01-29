package main

import "fmt"

func incrementNum(arr []int) []int {

	for iter := len(arr) - 1; iter >= 0; iter-- {
		if arr[iter] == 9 {
			arr[iter] = 0
		} else {
			arr[iter] += 1
			break
		}
	}

	if arr[0] == 0 {
		return append([]int{1}, arr...)
	}
	return arr
}

func main() {

	var num = []int{9, 9}

	num = incrementNum(num)

	fmt.Println(num)
}
