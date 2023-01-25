package main

import "fmt"

func main() {

	// Similar to single dimension arrays, Go allows users to create multi-dimensional arrays.
	// variable declaration format : var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type

	var intArray2D [2][3]int
	var intArr2D_2 = [3][3]float32{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}
	intArr2D_3 := [3][3]int32{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}

	var strArray = [2][2]string{{"I", "am"}, {"super", "human"}}

	var charArr2D = [2][3]rune{{'a', 'b'}, {'h', 'i'}}

	fmt.Println(intArray2D, intArr2D_2, intArr2D_3)
	fmt.Println(strArray)
	fmt.Printf("%c", charArr2D)

}
