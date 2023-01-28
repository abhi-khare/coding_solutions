package main

import "fmt"

func getPossibleWays(len int) int {
	if len < 4 {
		return 1
	} else{
		c1 := getPossibleWays(len-1)
		c2 := getPossibleWays(len-4)

		return c1+c2
	}
}

func main() {

	var len int
	fmt.Scan(&len)

	count := getPossibleWays(len)

	fmt.Print(count)
}