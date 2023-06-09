package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {

	count := 0

	for iter:=0; iter<len(flowerbed); iter++{
		if flowerbed[iter]==1{

			if iter!=0{
			flowerbed[iter-1] = -1
			}

			if iter != len(flowerbed)-1{
			flowerbed[iter+1] = -1
			}
		}
	}

	for idx, ele := range flowerbed{
		if ele == 0 && flowerbed[idx-1] !=1 && flowerbed[idx+1]!=1{
			flowerbed[idx]=1
			count ++
		}
	}

	if count >= n{
		return true
	} else{
		return false
	}

}

func main() {

	flowerBed := []int{ 1,0,0,0,0,1}
	n := 2

	ans := canPlaceFlowers(flowerBed, n)

	fmt.Println(ans)
}