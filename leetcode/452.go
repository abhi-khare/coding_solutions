package main

import (
	"fmt"
	"sort"
)



func main() {

	points := [][]int{{3, 5},{3, 5},{6, 7}}

	// first we will sort the array by the end of the ballon coordinate

	sort.Slice(points, func(i, j int) bool{
		return points[i][1] < points[j][1]
			
	})

	arrow_required := 1

	start_idx := points[0][1]

	for iter:=0; iter< len(points); iter++{
		if points[iter][0] > start_idx{
			arrow_required +=1
			start_idx = points[iter][1]
		}
	}

	fmt.Println(arrow_required)
}