package main

import (
	"fmt"
	"sort"
)

func pairSticks(chopsticks []int, diff int) int {

	pairs := 0
	sort.Slice(chopsticks, func(i,j int)bool{
		return chopsticks[i]< chopsticks[j]
	})

	for iter:= len(chopsticks)-1 ; iter>0; {

		if chopsticks[iter] - chopsticks[iter-1] <= diff{
			iter -=2
			pairs +=1
		} else {
			iter-=1
		}
	}

	return pairs
}


func main() {

	var chopsticks = []int{1,2,3,5}
	var diff int = 1
	pairNums := pairSticks(chopsticks, diff)

	fmt.Println(pairNums)
}