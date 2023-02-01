package main

import "fmt"

func printSubArray(arr []int){

	for iter:= 0; iter< len(arr)-1; iter++{
		for jter:=iter+1; jter<len(arr); jter++{

			for kter:=iter; kter<=jter; kter++{
				fmt.Print(arr[kter], " ")
			}
			fmt.Println("")
		}
	}
}


func main(){

	var arr = []int{12, 3, 4, 11, 17, 8, 9, 10}

	printSubArray(arr)
}