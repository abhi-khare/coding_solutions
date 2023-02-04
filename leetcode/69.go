package main

import "fmt"

func SquareRoot(x int) int {
	if x == 0 {
		return 0
	} else {
		start, end := 1, x

		for start <= end {
			mid := (start + end) / 2

			square_num := mid * mid
			if square_num == x {
				return mid
			} else if square_num < x {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}

		return end
	}
}

func main() {

	var x int = 21

	sqrtNum := SquareRoot(x)

	fmt.Println(sqrtNum)
}
