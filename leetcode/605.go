package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {

	count := 0

	if len(flowerbed) == 1 {
		if flowerbed[0] == 0{
		count +=1}

	} else if len(flowerbed)==2{
		if flowerbed[0] ==0 && flowerbed[1] == 0{
		count+=1}
	} else{
		
		for idx, ele := range flowerbed{

			if idx == 0{
				if ele == 0 && flowerbed[idx+1] == 0{
					flowerbed[idx]=1
					count+=1
				}
			} else if idx == len(flowerbed)-1 {
				if ele ==0 && flowerbed[idx-1]==0{
				count +=1
				flowerbed[idx]=1}
			} else{
				if ele == 0 && flowerbed[idx+1] == 0 && flowerbed[idx-1] == 0{
					count+=1
					flowerbed[idx]=1
				}
			}
		}

	}



	if count >= n{
		return true
	} else{
		return false
	}
}

func main() {

	input := []int{0,1,0}
	n := 0

	ans := canPlaceFlowers(input, n)

	fmt.Print(ans)
}