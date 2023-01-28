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
	}else {
		if y%2 == 0{
			half_power :=  power(x, y/2) 
			return half_power*half_power
		}else{
			half_power :=  power(x, y/2) 
			return x*half_power*half_power
			
		}
		
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