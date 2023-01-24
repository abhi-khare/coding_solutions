package main

import "fmt"

func main() {

	// Just like other general purpose PLs, Go allows users to create fixed sized arrays of in-built and user created data types.
	// We can create arrays of type int16/32/64, float32/34, string etc. using following two methods:
	// var <variableName> = [array_size]<data_type>{[optional initialisation]}
	// <variableName> := [array_size]<data_type>{[optional initialisation]}
	// If not initialised by the users, arrays get initialised by default values Ex: 0 for int/floats and "" for string

	var intArr = [3]int{}
	var floatArr = [5]float64{}

	intArr_2 := [3]int{}
	strArray := [10]string{}

	arrFloat := [4]float32{1.21, 1.32, 1.56, 6.78}

	// Don'ts:
	// arrFloat := [4]float32{1.21, 1.32, 1.56, 6.78, 3.24}  This will give error as more elements were initialised than the size of array.
	// arr := []float32{121, 1.23} This is not a traditional fixed sized array, but slice which is like a dynamic array, can be resized as per need.

	// We can use runes to create character arrays. runes are like char, but behind the scene they are just int32.
	// A golang string is not a sequence of rune, they are sequence of bytes.

	charArr := [3]rune{'a', 'b', 'c'}

	fmt.Println(intArr, intArr_2, floatArr, strArray, arrFloat, charArr)

}
