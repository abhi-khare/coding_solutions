package main

import "fmt"

func roman2Int(roman string, mapping map[string]int) int {
	num := 0
	iter:=0
	for ; iter<len(roman)-1; iter++{
		char := string(roman[iter])
		nextChar := string(roman[iter+1])
		if char == "I" && (nextChar == "V" || nextChar == "X"){
			num += mapping[nextChar] - mapping[char]
			iter += 1
		} else if char == "X" && (nextChar == "L" || nextChar == "C"){
			num += mapping[nextChar] - mapping[char]
			iter += 1
		} else if char == "C" && (nextChar == "D" || nextChar == "M"){
			num += mapping[nextChar] - mapping[char]
			iter += 1
		} else{
			num += mapping[char]
		}

	}
	if iter != len(roman){
		num += mapping[string(roman[len(roman)-1])]
	}
	
	return num
}

func main() {

	roman := "CI"

	mapping := make(map[string]int)
	mapping["I"] = 1
	mapping["V"] = 5
	mapping["X"] = 10
	mapping["L"] = 50
	mapping["C"] = 100
	mapping["D"] = 500
	mapping["M"] = 1000
	

	num := roman2Int(roman, mapping)

	fmt.Println(num)
}