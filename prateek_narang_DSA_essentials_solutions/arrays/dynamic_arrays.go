package main

import "fmt"

func main(){

	// In GO, dynamic arrays are simply called slice. 
	// To create slice, the format is same as array, but without the size.
	var intSlice = []int{1,2,3}
	var intSlice2 = []int{}

	// string slice initialised statically
  	// similarly, we can make slice of any datatype
	var strSlice = []string{"I", "am", "going", "home"}
	var floatSlice = []float32{}

	// creating character slices. char slice or arrays are not strings
  	// runes are just the alias for int32, that we can
  	// format to characters
	var charSlice = []rune{'a','b','h','i'} 

	fmt.Println("Statically initialised int slice: ", intSlice)
	fmt.Println("int slice: ", intSlice2)
	fmt.Println("Statically initialised string slice: ", strSlice)
	fmt.Println("float slice: ", floatSlice)
	fmt.Printf("statically  initialised rune slice: %c", charSlice)
}