package main

import "fmt"

func computeFibonacci(num int) int {

	if num == 0 || num == 1{
		return num
	} else {
		return computeFibonacci(num-1) + computeFibonacci(num-2)
	}
}

func main() {
	var num int

	fmt.Scan(&num)

	for iter:=0; iter < num; iter++{
		fibonacci_num := computeFibonacci(iter)

		fmt.Println(fibonacci_num)
	}
	
}