package main

import "fmt"

func main(){

	// statically created and initialised array
	var arr = [3]int{1,2,3}

	// statically created array
	var arr2 = [4]int{}

	// string arrays initialised statically
  	// similarly, we can make array of any datatype
	var strArr = [4]string{"I", "am", "going", "home"}
	var floatarr = [4]float32{}

	// creating character arrays
  	// runes are just the alias for int32, that we can
  	// format to characters
	var charArr = [4]rune{'a','b','h','i'} 

	fmt.Println("Statically created and initialised int array: ", arr)
	fmt.Println("Statically created int array: ", arr2)
	fmt.Println("Statically created and initialised string array: ", strArr)
	fmt.Println("Statically created float array: ", floatarr)
	fmt.Printf("statically created and initialised rune array: %c", charArr)
}