package main

import "fmt"

func main() {

	// We can use slice to create queue. Stack mainly contains two operations
	// push and pop performed in FIFO manner. We can use append and : operator
	// to mimic push and pop. While popping an element using : operator we need to
	// make sure that it has atleast 1 element to be popped.

	var intQueue = []int{}

	var size_ int
	fmt.Println("Enter the number of elements ")
	fmt.Scan(&size_)

	for iter := 0; iter < size_; iter++ {
		var ele int
		fmt.Scan(&ele)
		intQueue = append(intQueue, ele)
	}

	fmt.Println("Final queue after all push operations: ", intQueue)

	intQueue = intQueue[1:]
	fmt.Println("Final Stack after a pop operations: ", intQueue)
	intQueue = intQueue[1:]
	fmt.Println("Final Stack after a pop operations: ", intQueue)
	intQueue = intQueue[1:]
	fmt.Println("Final Stack after a pop operations: ", intQueue)
	intQueue = intQueue[1:]
	fmt.Println("Final Stack after a pop operations: ", intQueue)

}
