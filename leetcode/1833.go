package main

import "sort"
import "fmt"

func main() {

	costs := []int{1,6,3,1,2,5}
	money := 20
	ice_Creams := 0

	sort.Slice(costs, func(i, j int) bool {
		return costs[i] < costs[j]
	})

	fmt.Println(costs)

	for _,cost := range costs {
		if cost <= money {
			ice_Creams += 1
			money -= cost
		} else {
			break
		}
	}

	print(ice_Creams)
}