package main

import "fmt"

func getMovesCount(num_disk int) int {

	if num_disk == 1{
		return 1
	} else {
		c1 := getMovesCount(num_disk-1)
		return 1 + 2*c1
	}
}

func main() {

	var num_disk int

	fmt.Scan(&num_disk)

	movesCount := getMovesCount(num_disk)

	fmt.Print(movesCount)
}