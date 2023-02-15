package main

import "fmt"

func setZeroes(matrix [][]int) {
	rowMap, colMap := make([]int, len(matrix)), make([]int, len(matrix[0]))

	for iter := 0; iter < len(matrix); iter++ {
		for jter := 0; jter < len(matrix[0]); jter++ {

			if matrix[iter][jter] == 0 {
				rowMap[iter] = 1
				colMap[jter] = 1
			}
		}
	}

	for iter := 0; iter < len(matrix); iter++ {
		for jter := 0; jter < len(matrix[0]); jter++ {
			if rowMap[iter] == 1 || colMap[jter] == 1 {
				matrix[iter][jter] = 0
			}
		}
	}

}

func main() {

	var arr = [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(arr)

	fmt.Println(arr)

}
