package main

import "fmt"

func printIncreasingOrder(num int){
	if num == 1{
		fmt.Print(1, " ")
	} else{
		printIncreasingOrder(num-1)
		fmt.Print(num, " ")
	}
}

func printDecreasingOrder(num int){
	if num==1{
		fmt.Print(1, " ")
	} else{
		fmt.Print(num, " ")
		printDecreasingOrder(num-1)
	}
}

func main() {

	var num int

	fmt.Scan(&num)

	printIncreasingOrder(num)

	printDecreasingOrder(num)
}