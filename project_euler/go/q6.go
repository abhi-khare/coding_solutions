package main

import "fmt"

// sum of first n number is n*(n+1)/2
// sum of square of n numbers is n(n+1)(2n+1)/6

func main() {

	var num int64 = 100

	sum_n_squared := num * num * (num + 1) * (num + 1) / 4

	squared_num_n := num * (2*num + 1) * (num + 1) / 6

	fmt.Println("difference : ", sum_n_squared-squared_num_n)
}
