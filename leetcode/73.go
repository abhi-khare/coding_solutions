package main

import "fmt"

func setZeroes(matrix [][]int) {

	flag1, flag2 := 0, 0

	for iter := 0; iter < len(matrix); iter++ {
		for jter := 0; jter < len(matrix[0]); jter++ {
			if matrix[0][jter] == 0 {
				flag1 = 1
			}

			if matrix[iter][0] == 0 {
				flag2 = 1
			}
			if matrix[iter][jter] == 0 {
				matrix[0][jter] = 0
				matrix[iter][0] = 0
			}

		}
	}
	fmt.Println(matrix)

	for iter := 1; iter < len(matrix); iter++ {
		if matrix[iter][0] == 0 {
			for jter := 1; jter < len(matrix[0]); jter++ {
				matrix[iter][jter] = 0
			}
		}
	}
	fmt.Println(matrix)

	for jter := 1; jter < len(matrix[0]); jter++ {
		if matrix[0][jter] == 0 {
			for iter := 1; iter < len(matrix); iter++ {
				matrix[iter][jter] = 0
			}
		}
	}

	fmt.Println(matrix)

	if flag1 == 1 {
		for jter := 0; jter < len(matrix[0]); jter++ {
			matrix[0][jter] = 0
		}
	}

	if flag2 == 1 {
		for iter := 0; iter < len(matrix); iter++ {
			matrix[iter][0] = 0
		}
	}

}

func main() {

	var arr = [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(arr)

	fmt.Println(arr)

}
