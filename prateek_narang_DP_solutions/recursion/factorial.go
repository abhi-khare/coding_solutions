package main

import "fmt"

func computeFactorial(num int) int {

	if num <= 1 {
		return 1
	} else {
		return num * computeFactorial(num-1)
	}
}

func main() {

	var num int

	fmt.Scan(&num)

	factorial_num := computeFactorial(num)

	println("factorial: ", factorial_num)
}