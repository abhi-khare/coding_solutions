package main

import "fmt"

// simple double loop method to find the triplet
func main() {

	var sum int64 = 1000

	for iter := int64(1); iter < 999; iter++ {

		for jter := iter + 1; jter < 1000; jter++ {

			c := (sum - iter - jter)
			if iter*iter+jter*jter == c*c {
				fmt.Println("product of ABC is ", iter, jter, c)
			}
		}
	}
}
