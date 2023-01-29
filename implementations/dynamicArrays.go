package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Just like vector in C++, Go allows users to create variable sized arrays of in-built and user created data types.
	// These are called slice. We can create slice of type int16/32/64, float32/34, string etc. using following three methods:
	// var <variableName> = []<data_type>{[optional initialisation]}
	// <variableName> := []<data_type>{[optional initialisation]}
	// var <variableName>  []<data_type>
	// If not initialised by the users, slice remains null

	var intSlice = []int{}
	var floatSlice = []float64{}
	var Slc []int32

	intSlc_2 := []int{}
	strSlice := []string{}

	sliceFloat := []float32{1.21, 1.32, 1.56, 6.78}

	// appending element to the slice
	intSlice = append(intSlice, 2)

	// create slice using make.
	var size_ int
	fmt.Scan(&size_)
	dynSlice := make([]int, size_)

	dynSlice = append(dynSlice, 2)

	rt := reflect.TypeOf(dynSlice)
	switch rt.Kind() {
	case reflect.Slice:
		fmt.Println("is a slice with element type", rt.Elem())
	case reflect.Array:
		fmt.Println("is an array with element type", rt.Elem())
	default:
		fmt.Println("is something else entirely")
	}

	// Don'ts:
	// arrFloat := [4]float32{1.21, 1.32, 1.56, 6.78, 3.24}  This will give error as more elements were initialised than the size of array.
	// arr := []float32{121, 1.23} This is not a traditional fixed sized array, but slice which is like a dynamic array, can be resized as per need.

	// We can use runes to create character arrays. runes are like char, but behind the scene they are just int32.
	// A golang string is not a sequence of rune, they are sequence of bytes.

	charSlc := []rune{'a', 'b', 'c'}

	fmt.Println(intSlice, intSlc_2, floatSlice, strSlice, sliceFloat, charSlc, Slc, dynSlice)

}
