package main

import "fmt"


func linearSearch(arr []int, key int) int {

	for iter:=0; iter < len(arr); iter++{
		if arr[iter] == key{
			return iter
		}
	}

	return -1
}

func main() {

	var key, size_ int

	fmt.Println("Enter the size of the array!")
	fmt.Scan(&size_)

	var arr = make([]int, size_)

	fmt.Println("Enter the elements of the array")
	for iter:=0; iter<size_; iter++{
		fmt.Scan(&arr[iter])
	}

	fmt.Println("Enter the key")
	fmt.Scan(&key)


	idx := linearSearch(arr, key)

	if idx == -1{
		fmt.Print("Key not found in the array")
	} else{
		fmt.Print("Element found at index: ", idx)
	}

}