package main

import "fmt"

type nodeGraph struct {
	vertexId int
	ptr      *nodeGraph
}

func main() {

	var vertex_cnt int
	fmt.Scan(&vertex_cnt)

	var adjList = make([]*nodeGraph, vertex_cnt)

	fmt.Print(len(adjList))

	var edgeNum int
	fmt.Println("Enter the number of Edges and then enter all the edges [0 indexing]")
	fmt.Scan(&edgeNum)

	for iter := 0; iter < edgeNum; iter++ {
		var m, n int

		fmt.Scan(&m, &n)

		// creating node
		var node nodeGraph
		node.vertexId = n
		node.ptr = nil

		if adjList[m] == nil {
			adjList[m] = &node
		} else {
			node.ptr = adjList[m]
			adjList[m] = &node
		}
	}

	// iterating through graph

	for iter := 0; iter < vertex_cnt; iter++ {
		iterator := adjList[iter]
		for iterator != nil {
			println(iter, iterator.vertexId)
			iterator = iterator.ptr
		}
	}

}
