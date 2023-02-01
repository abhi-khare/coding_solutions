package main 

import "fmt"

func printPairs(arr []int) {

	for iter:=0; iter< len(arr)-1; iter++{

		for jter:= iter+1; jter<len(arr); jter++{
			fmt.Println(arr[iter], arr[jter])
		}
	}
}


func main(){

	var arr = []int{12,11,2,3,16,13}

	printPairs(arr)
}