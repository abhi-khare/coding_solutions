package main

import (
	"fmt"
	"sort"
)

type points struct {
	xCor int
	yCor int
}

func main() {

	var size_ int = 5

	var coordinates = make([]points, size_)

	for iter := 0; iter < size_; iter++ {
		var pt points
		fmt.Scan(&pt.xCor, &pt.yCor)

		coordinates[iter] = pt
	}

	sort.Slice(coordinates, func(i, j int) bool {
		return (coordinates[i].xCor < coordinates[j].xCor) || ((coordinates[i].xCor == coordinates[j].xCor) && (coordinates[i].yCor < coordinates[j].yCor))
	})

	fmt.Println(coordinates)
}
