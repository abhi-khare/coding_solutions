package main

import "fmt"

func main() {

	var sum int32 = 0 // storing the sum of multiple of 3 and 5 in 'sum' variable

	// accumulating the multiples of 3
	for i := 1; i < 1000; i++ {

		if i%3 == 0 || i%5 == 0 {
			sum = sum + int32(i)
		}
	}

	fmt.Printf("Sum of multiples of 3 and 5 : %v", sum)
}
