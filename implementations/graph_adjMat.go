package main

import "fmt"

func main() {

	var vertex_cnt int
	fmt.Scan(&vertex_cnt)

	// creating adjacency matrix
	var adjMat = make([][]int, vertex_cnt)
	for i := range adjMat {
		adjMat[i] = make([]int, vertex_cnt)
	}

	var edge int
	fmt.Println("Enter the number of Edges and then enter all the edges [0 indexing]")
	fmt.Scan(&edge)

	for iter := 0; iter < edge; iter++ {
		var m, n int
		fmt.Scan(&m, &n)
		adjMat[m][n] = 1
	}
	fmt.Println(adjMat)
}
