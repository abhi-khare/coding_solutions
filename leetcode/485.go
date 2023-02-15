package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
    maxStreak := 0
	currentStreak :=0
	for iter:=0; iter<len(nums); iter++{
		if nums[iter]==1{
			currentStreak +=1
            if currentStreak > maxStreak{
				maxStreak = currentStreak
				
			}
		} else{
			
			currentStreak = 0
			
		}
	}

	if currentStreak > maxStreak{
		maxStreak = currentStreak
	}

	return maxStreak
}

func main(){

	var arr = []int{1,0,1,1,0,1}

	streak := findMaxConsecutiveOnes(arr)

	fmt.Println(streak)
}