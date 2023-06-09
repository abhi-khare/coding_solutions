package main

import "fmt"

func arraySign(nums []int) int {
    numNeg, numPos := 0, 0

	for _, ele := range nums{
		if ele <0{
			numNeg +=1
		} else if ele > 0{
			numPos +=1
		} else{
			return 0
		}
	}

	if numNeg%2 ==0{
		return 1
	} else{
		return -1
	}
}


func main(){

	nums := []int{1}

	ans := arraySign(nums)

	fmt.Println(ans)
}