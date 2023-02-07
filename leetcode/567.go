package main

import "fmt"


func checkInclusion( s1 string, s2 string) bool {

	mapS1 := make(map[string]int)
	// converting s1 into a hashmap
	for _, ele := range s1{
		_, ok := mapS1[string(ele)]
		if !ok{
			mapS1[string(ele)] = 1
		} else{
			mapS1[string(ele)] += 1
		}
	}

	iter :=0 
	for iter < len(s2) {
		char = string(s2[iter])
		mapS1 := make(map[string]int)
		val, ok := mapS1[char]
		
		if !ok 
	}

}

func main() {

	s1, s2 := "", ""

	status := checkInclusion(s1, s2)

	fmt.Println(status)

}