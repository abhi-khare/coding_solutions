package main

import "fmt"

func main() {

	// We can use slice to create stack. Stack mainly contains two operations
	// push and pop in LIFO manner. We can use append and : operator to mimic push and pop.
	// While popping an element using : operator we need to make sure that it has
	// atleast 1 element to be popped.

	var intStack = []int{}

	var size_ int
	fmt.Println("Enter the number of elements ")
	fmt.Scan(&size_)

	for iter := 0; iter < size_; iter++ {
		var ele int
		fmt.Scan(&ele)
		intStack = append(intStack, ele)
	}

	fmt.Println("Final Stack after all push operations: ", intStack)

	intStack = intStack[:len(intStack)-1]
	fmt.Println("Final Stack after a pop operations: ", intStack)
	intStack = intStack[:len(intStack)-1]
	fmt.Println("Final Stack after a pop operations: ", intStack)
	intStack = intStack[:len(intStack)-1]
	fmt.Println("Final Stack after a pop operations: ", intStack)
	intStack = intStack[:len(intStack)-1]
	fmt.Println("Final Stack after a pop operations: ", intStack)

}
