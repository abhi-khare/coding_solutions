package main

import "fmt"


func power(x int, y int) int{

	if (x == 1) {
		return 1
	} else if (x == 0 && y!=0){
		return 0
	} else if  (x == 0 && y==0){
		return -1
	}

	if y == 0{
		return 1
	} else {
		return x * power(x, y-1)
	}
}

func main() {

	var x, y int

	fmt.Scan(&x, &y)

	// x^y 

	output := power(x,y)

	if output == -1{
		print("Undefined")
	} else{
		print(output)
	}

	
}