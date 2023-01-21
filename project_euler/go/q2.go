package main

import "fmt"

func main() {
	var a int32 = 1 // initial seed value
	var b int32 = 2
	var sum int32 = 0 // variable to store the sum of even values fibonacci numbers

	for b < 4000000 {
		sum += b
		m, n := a+2*b, 2*a+3*b
		a, b = m, n
	}

	fmt.Printf("sum of even values fibonacci numbers < 4M : %v", sum)
}
