package main

import (
	"fmt"
	"math"
)

// general solution is (m+n)! / (m! * n!)
// which can be simplified to 2*(n/2)[sum 2n-1 , 2n-3, 2n-5 .... n+1]/ (n/2)!
func main() {

	num := 20

	solution := int64(1)

	for iter := num + 1; iter <= 2*num-1; iter += 2 {
		solution = solution * int64(iter)
	}

	solution = solution * int64(math.Pow(2, float64(num/2)))

	for iter := 1; iter <= num/2; iter++ {
		solution = solution / int64(iter)
	}

	fmt.Println("possible number of paths are ", solution)
}
